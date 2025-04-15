package types

import (
	"bytes"
	"fmt"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/node/appinterface"
	"github.com/gatechain/gatechainsdk/gatechain/staking/exported"
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
	"strings"
	"time"

	"github.com/gatechain/crypto"
	// tmtypes "github.com/tendermint/tendermint/types"
	yaml "gopkg.in/yaml.v2"
)

// nolint
const (
	// TODO: Why can't we just have one string description which can be JSON by convention
	MaxMonikerLength  = 70
	MaxIdentityLength = 3000
	MaxWebsiteLength  = 140
	MaxDetailsLength  = 280
)

// Implements Validator interface
var _ exported.ValidatorI = Validator{}

// Validator defines the total amount of bond shares and their exchange rate to
// coins. Slashing results in a decrease in the exchange rate, allowing correct
// calculation of future undelegations without iterating over delegators.
// When coins are delegated to this validator, the validator is credited with a
// delegation whose number of bond shares is based on the amount of coins delegated
// divided by the current exchange rate. Voting power can be calculated as total
// bonded shares multiplied by exchange rate.
type Validator struct {
	OperatorAddress         types2.ValAddress `json:"operator_address" yaml:"operator_address"`       // address of the validator's operator; bech encoded in JSON
	ConsPubKey              crypto.PubKey     `json:"consensus_pubkey" yaml:"consensus_pubkey"`       // the consensus public key of the validator; bech encoded in JSON
	Jailed                  bool              `json:"jailed" yaml:"jailed"`                           // has the validator been jailed from bonded status?
	Status                  types2.BondStatus `json:"status" yaml:"status"`                           // validator status (bonded/unbonding/unbonded)
	Tokens                  types2.Int        `json:"tokens" yaml:"tokens"`                           // delegated tokens
	PowerRate               types2.Dec        `json:"power_rate" yaml:"power_rate"`                   // poewe rate
	DelegatorShares         types2.Dec        `json:"delegator_shares" yaml:"delegator_shares"`       // total shares issued to a validator's delegators
	Description             Description       `json:"description" yaml:"description"`                 // description terms for the validator
	UnbondingHeight         int64             `json:"undelegating_height" yaml:"undelegating_height"` // if unbonding, height at which this validator has begun unbonding
	UnbondingCompletionTime time.Time         `json:"undelegating_time" yaml:"undelegating_time"`     // if unbonding, min time for the validator to complete unbonding
	Commission              Commission        `json:"commission" yaml:"commission"`                   // commission parameters
}

// custom marshal yaml function due to consensus pubkey
func (v Validator) MarshalYAML() (interface{}, error) {
	bs, err := yaml.Marshal(struct {
		OperatorAddress         types2.ValAddress
		ConsPubKey              string
		Jailed                  bool
		Status                  types2.BondStatus
		Tokens                  types2.Int
		DelegatorShares         types2.Dec
		PowerRate               types2.Dec
		Description             Description
		UnbondingHeight         int64
		UnbondingCompletionTime time.Time
		Commission              Commission
	}{
		OperatorAddress:         v.OperatorAddress,
		ConsPubKey:              types2.MustBech32ifyConsPub(v.ConsPubKey),
		Jailed:                  v.Jailed,
		Status:                  v.Status,
		Tokens:                  v.Tokens,
		DelegatorShares:         v.DelegatorShares,
		PowerRate:               v.PowerRate,
		Description:             v.Description,
		UnbondingHeight:         v.UnbondingHeight,
		UnbondingCompletionTime: v.UnbondingCompletionTime,
		Commission:              v.Commission,
	})
	if err != nil {
		return nil, err
	}

	return string(bs), nil
}

// Validators is a collection of Validator
type Validators []Validator

func (v Validators) String() (out string) {
	for _, val := range v {
		out += val.String() + "\n"
	}
	return strings.TrimSpace(out)
}

// genesis accounts contain an address
func (v Validators) Contains(acc types2.ValAddress) bool {
	for _, val := range v {
		if val.OperatorAddress.Equals(acc) {
			return true
		}
	}
	return false
}

// ToSDKValidators -  convenience function convert []Validators to []framework.Validators
func (v Validators) ToSDKValidators() (validators []exported.ValidatorI) {
	for _, val := range v {
		validators = append(validators, val)
	}
	return validators
}

// NewValidator - initialize a new validator
func NewValidator(operator types2.ValAddress, pubKey crypto.PubKey, description Description) Validator {
	return Validator{
		OperatorAddress:         operator,
		ConsPubKey:              pubKey,
		Jailed:                  false,
		Status:                  types2.Unbonded,
		Tokens:                  types2.ZeroInt(),
		DelegatorShares:         types2.ZeroDec(),
		PowerRate:               types2.ZeroDec(),
		Description:             description,
		UnbondingHeight:         int64(0),
		UnbondingCompletionTime: time.Unix(0, 0).UTC(),
		Commission:              NewCommission(types2.ZeroDec(), types2.ZeroDec(), types2.ZeroDec()),
	}
}

// return the redelegation
func MustMarshalValidator(cdc *codec.Codec, validator Validator) []byte {
	return cdc.MustMarshalBinaryLengthPrefixed(validator)
}

// unmarshal a redelegation from a store value
func MustUnmarshalValidator(cdc *codec.Codec, value []byte) Validator {
	validator, err := UnmarshalValidator(cdc, value)
	if err != nil {
		panic(err)
	}
	return validator
}

// unmarshal a redelegation from a store value
func UnmarshalValidator(cdc *codec.Codec, value []byte) (validator Validator, err error) {
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &validator)
	return validator, err
}

// String returns a human readable string representation of a validator.
func (v Validator) String() string {
	bechConsPubKey, err := types2.Bech32ifyConsPub(v.ConsPubKey)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(`Validator
  Operator Address:           %s
  Validator Consensus Pubkey: %s
  Jailed:                     %v
  Status:                     %s
  Tokens:                     %s
  Delegator Shares:           %s
  Power Rate:                 %s
  Description:                %s
  Unbonding Height:           %d
  Unbonding Completion Time:  %v
  Commission:                 %s`, v.OperatorAddress, bechConsPubKey,
		v.Jailed, v.Status, v.Tokens,
		v.DelegatorShares, v.PowerRate, v.Description,
		v.UnbondingHeight, v.UnbondingCompletionTime, v.Commission)
}

// this is a helper struct used for JSON de- and encoding only
type bechValidator struct {
	OperatorAddress         types2.ValAddress `json:"operator_address" yaml:"operator_address"`       // the bech32 address of the validator's operator
	ConsPubKey              string            `json:"consensus_pubkey" yaml:"consensus_pubkey"`       // the bech32 consensus public key of the validator
	Jailed                  bool              `json:"jailed" yaml:"jailed"`                           // has the validator been jailed from bonded status?
	Status                  types2.BondStatus `json:"status" yaml:"status"`                           // validator status (bonded/unbonding/unbonded)
	Tokens                  types2.Int        `json:"tokens" yaml:"tokens"`                           // delegated tokens
	PowerRate               types2.Dec        `json:"power_rate" yaml:"power_rate"`                   // poewe rate
	DelegatorShares         types2.Dec        `json:"delegator_shares" yaml:"delegator_shares"`       // total shares issued to a validator's delegators
	Description             Description       `json:"description" yaml:"description"`                 // description terms for the validator
	UnbondingHeight         int64             `json:"undelegating_height" yaml:"undelegating_height"` // if unbonding, height at which this validator has begun unbonding
	UnbondingCompletionTime time.Time         `json:"undelegating_time" yaml:"undelegating_time"`     // if unbonding, min time for the validator to complete unbonding
	Commission              Commission        `json:"commission" yaml:"commission"`                   // commission parameters
}

// MarshalJSON marshals the validator to JSON using Bech32
func (v Validator) MarshalJSON() ([]byte, error) {
	bechConsPubKey, err := types2.Bech32ifyConsPub(v.ConsPubKey)
	if err != nil {
		return nil, err
	}

	return codec.Cdc.MarshalJSON(bechValidator{
		OperatorAddress:         v.OperatorAddress,
		ConsPubKey:              bechConsPubKey,
		Jailed:                  v.Jailed,
		Status:                  v.Status,
		Tokens:                  v.Tokens,
		DelegatorShares:         v.DelegatorShares,
		PowerRate:               v.PowerRate,
		Description:             v.Description,
		UnbondingHeight:         v.UnbondingHeight,
		UnbondingCompletionTime: v.UnbondingCompletionTime,
		Commission:              v.Commission,
	})
}

// UnmarshalJSON unmarshals the validator from JSON using Bech32
func (v *Validator) UnmarshalJSON(data []byte) error {
	bv := &bechValidator{}
	if err := codec.Cdc.UnmarshalJSON(data, bv); err != nil {
		return err
	}
	consPubKey, err := types2.GetConsPubKeyBech32(bv.ConsPubKey)
	if err != nil {
		return err
	}
	*v = Validator{
		OperatorAddress:         bv.OperatorAddress,
		ConsPubKey:              consPubKey,
		Jailed:                  bv.Jailed,
		Tokens:                  bv.Tokens,
		Status:                  bv.Status,
		DelegatorShares:         bv.DelegatorShares,
		PowerRate:               bv.PowerRate,
		Description:             bv.Description,
		UnbondingHeight:         bv.UnbondingHeight,
		UnbondingCompletionTime: bv.UnbondingCompletionTime,
		Commission:              bv.Commission,
	}
	return nil
}

// only the vitals
func (v Validator) TestEquivalent(v2 Validator) bool {
	return v.ConsPubKey.Equals(v2.ConsPubKey) &&
		bytes.Equal(v.OperatorAddress, v2.OperatorAddress) &&
		v.Status.Equal(v2.Status) &&
		v.Tokens.Equal(v2.Tokens) &&
		v.DelegatorShares.Equal(v2.DelegatorShares) &&
		v.Description == v2.Description &&
		v.Commission.Equal(v2.Commission)
}

// return the TM validator address
func (v Validator) ConsAddress() types2.ConsAddress {
	return types2.ConsAddress(v.ConsPubKey.Address())
}

// IsBonded checks if the validator status equals Bonded
func (v Validator) IsBonded() bool {
	return v.GetStatus().Equal(types2.Bonded)
}

// IsUnbonded checks if the validator status equals Unbonded
func (v Validator) IsUnbonded() bool {
	return v.GetStatus().Equal(types2.Unbonded)
}

// IsUnbonding checks if the validator status equals Unbonding
func (v Validator) IsUnbonding() bool {
	return v.GetStatus().Equal(types2.Unbonding)
}

// constant used in flags to indicate that description field should not be updated
const DoNotModifyDesc = "[do-not-modify]"

// Description - description fields for a validator
type Description struct {
	Moniker  string `json:"moniker" yaml:"moniker"`   // name
	Identity string `json:"identity" yaml:"identity"` // optional identity signature (ex. UPort or Keybase)
	Website  string `json:"website" yaml:"website"`   // optional website link
	Details  string `json:"details" yaml:"details"`   // optional details
}

// NewDescription returns a new Description with the provided values.
func NewDescription(moniker, identity, website, details string) Description {
	return Description{
		Moniker:  moniker,
		Identity: identity,
		Website:  website,
		Details:  details,
	}
}

// UpdateDescription updates the fields of a given description. An error is
// returned if the resulting description contains an invalid length.
func (d Description) UpdateDescription(d2 Description) (Description, types2.Error) {
	if d2.Moniker == DoNotModifyDesc {
		d2.Moniker = d.Moniker
	}
	if d2.Identity == DoNotModifyDesc {
		d2.Identity = d.Identity
	}
	if d2.Website == DoNotModifyDesc {
		d2.Website = d.Website
	}
	if d2.Details == DoNotModifyDesc {
		d2.Details = d.Details
	}

	return Description{
		Moniker:  d2.Moniker,
		Identity: d2.Identity,
		Website:  d2.Website,
		Details:  d2.Details,
	}.EnsureLength()
}

// EnsureLength ensures the length of a validator's description.
func (d Description) EnsureLength() (Description, types2.Error) {
	if len(d.Moniker) > MaxMonikerLength {
		return d, ErrDescriptionLength(DefaultCodespace, "moniker", len(d.Moniker), MaxMonikerLength)
	}
	if len(d.Identity) > MaxIdentityLength {
		return d, ErrDescriptionLength(DefaultCodespace, "identity", len(d.Identity), MaxIdentityLength)
	}
	if len(d.Website) > MaxWebsiteLength {
		return d, ErrDescriptionLength(DefaultCodespace, "website", len(d.Website), MaxWebsiteLength)
	}
	if len(d.Details) > MaxDetailsLength {
		return d, ErrDescriptionLength(DefaultCodespace, "details", len(d.Details), MaxDetailsLength)
	}

	return d, nil
}

// ABCIValidatorUpdate returns an abci.ValidatorUpdate from a staking validator type
// with the full validator power
func (v Validator) ABCIValidatorUpdate() appinterface.AccountDelta {
	return appinterface.AccountDelta{
		Address: v.ConsAddress().Bytes(),
		Power:   uint64(v.ConsensusPower()),
	}
}

// ABCIValidatorUpdateZero returns an abci.ValidatorUpdate from a staking validator type
// with zero power used for validator updates.
func (v Validator) ABCIValidatorUpdateZero() appinterface.AccountDelta {
	return appinterface.AccountDelta{
		Address: v.ConsAddress().Bytes(),
		Power:   0,
	}
}

func (v Validator) ValidatorUpdatePower(amount types2.Int) appinterface.AccountDelta {
	addr := v.ConsAddress().Bytes()
	return appinterface.AccountDelta{
		Address: types2.GetAddressTO48(addr),
		Power:   v.AlgPower(amount),
	}
}

// SetInitialCommission attempts to set a validator's initial commission. An
// error is returned if the commission is invalid.
func (v Validator) SetInitialCommission(commission Commission) (Validator, types2.Error) {
	if err := commission.Validate(); err != nil {
		return v, err
	}

	v.Commission = commission
	return v, nil
}

func (v Validator) SetPowerRate(rate types2.Dec) (Validator, types2.Error) {
	v.PowerRate = rate
	return v, nil
}

// In some situations, the exchange rate becomes invalid, e.g. if
// Validator loses all tokens due to slashing. In this case,
// make all future delegations invalid.
func (v Validator) InvalidExRate() bool {
	return v.Tokens.IsZero() && v.DelegatorShares.IsPositive()
}

// calculate the token worth of provided shares
func (v Validator) TokensFromShares(shares types2.Dec) types2.Dec {
	return (shares.MulInt(v.Tokens)).Quo(v.DelegatorShares)
}

// calculate the token worth of provided shares, truncated
func (v Validator) TokensFromSharesTruncated(shares types2.Dec) types2.Dec {
	return (shares.MulInt(v.Tokens)).QuoTruncate(v.DelegatorShares)
}

// TokensFromSharesRoundUp returns the token worth of provided shares, rounded
// up.
func (v Validator) TokensFromSharesRoundUp(shares types2.Dec) types2.Dec {
	return (shares.MulInt(v.Tokens)).QuoRoundUp(v.DelegatorShares)
}

// SharesFromTokens returns the shares of a delegation given a bond amount. It
// returns an error if the validator has no tokens.
func (v Validator) SharesFromTokens(amt types2.Int) (types2.Dec, types2.Error) {
	if v.Tokens.IsZero() {
		return types2.ZeroDec(), ErrInsufficientShares(DefaultCodespace)
	}

	return v.GetDelegatorShares().MulInt(amt).QuoInt(v.GetTokens()), nil
}

// SharesFromTokensTruncated returns the truncated shares of a delegation given
// a bond amount. It returns an error if the validator has no tokens.
func (v Validator) SharesFromTokensTruncated(amt types2.Int) (types2.Dec, types2.Error) {
	if v.Tokens.IsZero() {
		return types2.ZeroDec(), ErrInsufficientShares(DefaultCodespace)
	}

	return v.GetDelegatorShares().MulInt(amt).QuoTruncate(v.GetTokens().ToDec()), nil
}

// get the bonded tokens which the validator holds
func (v Validator) BondedTokens() types2.Int {
	if v.IsBonded() {
		return v.Tokens
	}
	return types2.ZeroInt()
}

// get the consensus-engine power
// a reduction of 10^6 from validator tokens is applied
func (v Validator) ConsensusPower() int64 {
	if v.IsBonded() {
		return v.PotentialConsensusPower()
	}
	return 0
}

// potential consensus-engine power
func (v Validator) PotentialConsensusPower() int64 {
	return types2.TokensToConsensusPower(v.Tokens)
}

func (v Validator) AlgPower(amount types2.Int) uint64 {
	total := v.Tokens.Add(amount)
	power := v.PowerRate.MulInt(total).TruncateInt()
	gmPower := types2.TokensToConsensusPower(power)
	return uint64(gmPower)
}

// UpdateStatus updates the location of the shares within a validator
// to reflect the new status
func (v Validator) UpdateStatus(newStatus types2.BondStatus) Validator {
	v.Status = newStatus
	return v
}

// AddTokensFromDel adds tokens to a validator
func (v Validator) AddTokensFromDel(amount types2.Int) (Validator, types2.Dec) {

	// calculate the shares to issue
	var issuedShares types2.Dec
	if v.DelegatorShares.IsZero() {
		// the first delegation to a validator sets the exchange rate to one
		issuedShares = amount.ToDec()
	} else {
		shares, err := v.SharesFromTokens(amount)
		if err != nil {
			panic(err)
		}

		issuedShares = shares
	}

	v.Tokens = v.Tokens.Add(amount)
	v.DelegatorShares = v.DelegatorShares.Add(issuedShares)

	return v, issuedShares
}

// RemoveTokens removes tokens from a validator
func (v Validator) RemoveTokens(tokens types2.Int) Validator {
	if tokens.IsNegative() {
		panic(fmt.Sprintf("should not happen: trying to remove negative tokens %v", tokens))
	}
	if v.Tokens.LT(tokens) {
		panic(fmt.Sprintf("should not happen: only have %v tokens, trying to remove %v", v.Tokens, tokens))
	}
	v.Tokens = v.Tokens.Sub(tokens)
	return v
}

// RemoveDelShares removes delegator shares from a validator.
// NOTE: because token fractions are left in the valiadator,
//
//	the exchange rate of future shares of this validator can increase.
func (v Validator) RemoveDelShares(delShares types2.Dec) (Validator, types2.Int) {

	remainingShares := v.DelegatorShares.Sub(delShares)
	var issuedTokens types2.Int
	if remainingShares.IsZero() {

		// last delegation share gets any trimmings
		issuedTokens = v.Tokens
		v.Tokens = types2.ZeroInt()
	} else {

		// leave excess tokens in the validator
		// however fully use all the delegator shares
		issuedTokens = v.TokensFromShares(delShares).TruncateInt()
		v.Tokens = v.Tokens.Sub(issuedTokens)
		if v.Tokens.IsNegative() {
			panic("attempting to remove more tokens than available in validator")
		}
	}

	v.DelegatorShares = remainingShares
	return v, issuedTokens
}

// nolint - for ValidatorI
func (v Validator) IsJailed() bool                 { return v.Jailed }
func (v Validator) GetMoniker() string             { return v.Description.Moniker }
func (v Validator) GetStatus() types2.BondStatus   { return v.Status }
func (v Validator) GetOperator() types2.ValAddress { return v.OperatorAddress }
func (v Validator) GetConsPubKey() crypto.PubKey   { return v.ConsPubKey }
func (v Validator) GetConsAddr() types2.ConsAddress {
	return types2.ConsAddress(v.ConsPubKey.Address())
}
func (v Validator) GetTokens() types2.Int          { return v.Tokens }
func (v Validator) GetBondedTokens() types2.Int    { return v.BondedTokens() }
func (v Validator) GetPowerRate() types2.Dec       { return v.PowerRate }
func (v Validator) GetConsensusPower() int64       { return v.ConsensusPower() }
func (v Validator) GetCommission() types2.Dec      { return v.Commission.Rate }
func (v Validator) GetDelegatorShares() types2.Dec { return v.DelegatorShares }

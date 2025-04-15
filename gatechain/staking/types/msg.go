package types

import (
	"bytes"
	"encoding/json"
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"

	"github.com/gatechain/crypto"
)

// ensure Msg interface compliance at compile time
var (
	_ types2.Msg = &MsgCreateValidator{}
	_ types2.Msg = &MsgEditValidator{}
	_ types2.Msg = &MsgEditValidatorMaxRate{}
	_ types2.Msg = &MsgDelegate{}
	_ types2.Msg = &MsgUndelegate{}
	_ types2.Msg = &MsgBeginRedelegate{}
	_ types2.Msg = &MsgUndelegateByRetrievalAccount{}
)

//______________________________________________________________________

// MsgCreateValidator - struct for bonding transactions
type MsgCreateValidator struct {
	Description      Description       `json:"description" yaml:"description"`
	Commission       CommissionRates   `json:"commission" yaml:"commission"`
	DelegatorAddress types2.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types2.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	PubKey           crypto.PubKey     `json:"pubkey" yaml:"pubkey"`
	Extra            []byte            `json:"extra" yaml:"extra"`
}

type msgCreateValidatorJSON struct {
	Description      Description       `json:"description" yaml:"description"`
	Commission       CommissionRates   `json:"commission" yaml:"commission"`
	DelegatorAddress types2.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types2.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	PubKey           string            `json:"pubkey" yaml:"pubkey"`
	Extra            []byte            `json:"extra" yaml:"extra"`
}

// Default way to create validator. Delegator address and validator address are the same
func NewMsgCreateValidator(
	valAddr types2.ValAddress, pubKey crypto.PubKey,
	description Description, commission CommissionRates, extra []byte,
) MsgCreateValidator {

	return MsgCreateValidator{
		Description:      description,
		DelegatorAddress: types2.AccAddress(valAddr),
		ValidatorAddress: valAddr,
		PubKey:           pubKey,
		Commission:       commission,
		Extra:            extra,
	}
}

// nolint
func (msg MsgCreateValidator) Route() string { return RouterKey }
func (msg MsgCreateValidator) Type() string  { return "create_validator" }

// Return address(es) that must sign over msg.GetSignBytes()
func (msg MsgCreateValidator) GetSigners() []types2.AccAddress {
	// delegator is first signer so delegator pays fees
	addrs := []types2.AccAddress{msg.DelegatorAddress}

	if !bytes.Equal(msg.DelegatorAddress.Bytes(), msg.ValidatorAddress.Bytes()) {
		// if validator addr is not same as delegator addr, validator must sign
		// msg as well
		addrs = append(addrs, types2.AccAddress(msg.ValidatorAddress))
	}
	return addrs
}

// MarshalJSON implements the json.Marshaler interface to provide custom JSON
// serialization of the MsgCreateValidator type.
func (msg MsgCreateValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgCreateValidatorJSON{
		Description:      msg.Description,
		Commission:       msg.Commission,
		DelegatorAddress: msg.DelegatorAddress,
		ValidatorAddress: msg.ValidatorAddress,
		PubKey:           types2.MustBech32ifyConsPub(msg.PubKey),
		Extra:            msg.Extra,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface to provide custom
// JSON deserialization of the MsgCreateValidator type.
func (msg *MsgCreateValidator) UnmarshalJSON(bz []byte) error {
	var msgCreateValJSON msgCreateValidatorJSON
	if err := json.Unmarshal(bz, &msgCreateValJSON); err != nil {
		return err
	}

	msg.Description = msgCreateValJSON.Description
	msg.Commission = msgCreateValJSON.Commission
	msg.DelegatorAddress = msgCreateValJSON.DelegatorAddress
	msg.ValidatorAddress = msgCreateValJSON.ValidatorAddress
	msg.Extra = msgCreateValJSON.Extra
	var err error
	msg.PubKey, err = types2.GetConsPubKeyBech32(msgCreateValJSON.PubKey)
	if err != nil {
		return err
	}
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgCreateValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgCreateValidator) ValidateBasic() types2.Error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.DelegatorAddress.Empty() {
		return ErrNilDelegatorAddr(DefaultCodespace)
	}
	if msg.ValidatorAddress.Empty() {
		return ErrNilValidatorAddr(DefaultCodespace)
	}
	if msg.Extra == nil {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "extra must be included")
	}
	if !types2.AccAddress(msg.ValidatorAddress).Equals(msg.DelegatorAddress) {
		return ErrBadValidatorAddr(DefaultCodespace)
	}
	if msg.Description == (Description{}) {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "description must be included")
	}
	if msg.Commission == (CommissionRates{}) {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "commission must be included")
	}
	if err := msg.Commission.Validate(); err != nil {
		return err
	}
	if msg.PubKey == nil || len(msg.PubKey.Bytes()) == 0 {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "pubkey is empty")
	}
	if !msg.DelegatorAddress.Equals(types2.AccAddress(msg.PubKey.Address().Bytes())) {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "public key is error which pubKey.address is not equal to delegator address")
	}
	return nil
}

// MsgEditValidator - struct for editing a validator
type MsgEditValidator struct {
	Description
	ValidatorAddress types2.ValAddress `json:"address" yaml:"address"`

	// We pass a reference to the new commission rate and min self delegation as it's not mandatory to
	// update. If not updated, the deserialized rate will be zero with no way to
	// distinguish if an update was intended.
	//
	// REF: #2373
	CommissionRate *types2.Dec `json:"commission_rate" yaml:"commission_rate"`
}

func NewMsgEditValidator(valAddr types2.ValAddress, description Description, newRate *types2.Dec) MsgEditValidator {
	return MsgEditValidator{
		Description:      description,
		CommissionRate:   newRate,
		ValidatorAddress: valAddr,
	}
}

// nolint
func (msg MsgEditValidator) Route() string { return RouterKey }
func (msg MsgEditValidator) Type() string  { return "edit_validator" }
func (msg MsgEditValidator) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{types2.AccAddress(msg.ValidatorAddress)}
}

// get the bytes for the message signer to sign on
func (msg MsgEditValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgEditValidator) ValidateBasic() types2.Error {
	if msg.ValidatorAddress.Empty() {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "nil validator address")
	}

	if msg.Description == (Description{}) {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "transaction must include some information to modify")
	}

	if msg.CommissionRate != nil {
		if msg.CommissionRate.GT(types2.OneDec()) || msg.CommissionRate.LT(types2.ZeroDec()) {
			return types2.NewError(DefaultCodespace, CodeInvalidInput, "commission rate must be between 0 and 1, inclusive")
		}
	}

	return nil
}

type MsgEditValidatorMaxRate struct {
	ValidatorAddress types2.ValAddress `json:"address" yaml:"address"`
	MaxRate          *types2.Dec       `json:"max_rate" yaml:"max_rate"` // maximum commission rate which validator can ever charge, as a fraction
	MaxChangeRate    *types2.Dec       `json:"max_change_rate" yaml:"max_change_rate"`
}

func NewMsgEditValidatorMaxRate(valAddr types2.ValAddress, newMaxRate *types2.Dec, newMaxChangeRate *types2.Dec) MsgEditValidatorMaxRate {
	return MsgEditValidatorMaxRate{
		ValidatorAddress: valAddr,
		MaxRate:          newMaxRate,
		MaxChangeRate:    newMaxChangeRate,
	}
}

func (msg MsgEditValidatorMaxRate) Route() string { return RouterKey }
func (msg MsgEditValidatorMaxRate) Type() string  { return "edit_validator" }
func (msg MsgEditValidatorMaxRate) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{types2.AccAddress(msg.ValidatorAddress)}
}

// get the bytes for the message signer to sign on
func (msg MsgEditValidatorMaxRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgEditValidatorMaxRate) ValidateBasic() types2.Error {
	if msg.ValidatorAddress.Empty() {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "nil validator address")
	}

	if msg.MaxRate == nil {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "nil commission max rate")
	}
	if msg.MaxChangeRate == nil {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "nil commission max change rate")
	}

	if msg.MaxRate != nil {
		if msg.MaxRate.GT(types2.OneDec()) || msg.MaxRate.LT(types2.ZeroDec()) {
			return types2.NewError(DefaultCodespace, CodeInvalidInput, "commission max rate must be between 0 and 1, inclusive")
		}
	}

	if msg.MaxChangeRate != nil {
		if msg.MaxChangeRate.GT(types2.OneDec()) || msg.MaxChangeRate.LT(types2.ZeroDec()) {
			return types2.NewError(DefaultCodespace, CodeInvalidInput, "commission max change rate must be between 0 and 1, inclusive")
		}
	}
	return nil
}

// MsgValidatorSwitch - struct for a validator switch state to online offline
type MsgValidatorSwitchState struct {
	ValidatorAddress types2.ValAddress `json:"address" yaml:"address"`
	Extra            []byte            `json:"extra" yaml:"extra"`
}

func NewMsgValidatorSwitchState(valAddr types2.ValAddress, data []byte) MsgValidatorSwitchState {
	return MsgValidatorSwitchState{
		ValidatorAddress: valAddr,
		Extra:            data,
	}
}

// nolint
func (msg MsgValidatorSwitchState) Route() string { return RouterKey }
func (msg MsgValidatorSwitchState) Type() string  { return "switch_validator" }
func (msg MsgValidatorSwitchState) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{types2.AccAddress(msg.ValidatorAddress)}
}

// get the bytes for the message signer to sign on
func (msg MsgValidatorSwitchState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgValidatorSwitchState) ValidateBasic() types2.Error {
	if msg.ValidatorAddress.Empty() {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "nil validator address")
	}

	if msg.Extra == nil {
		return types2.NewError(DefaultCodespace, CodeInvalidInput, "extra must be included")
	}

	return nil
}

// MsgDelegate - struct for bonding transactions
type MsgDelegate struct {
	DelegatorAddress types2.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types2.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	Amount           types2.Coin       `json:"amount" yaml:"amount"`
}

func NewMsgDelegate(delAddr types2.AccAddress, valAddr types2.ValAddress, amount types2.Coin) MsgDelegate {
	return MsgDelegate{
		DelegatorAddress: delAddr,
		ValidatorAddress: valAddr,
		Amount:           amount,
	}
}

// nolint
func (msg MsgDelegate) Route() string { return RouterKey }
func (msg MsgDelegate) Type() string  { return "delegation" }
func (msg MsgDelegate) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{msg.DelegatorAddress}
}

// get the bytes for the message signer to sign on
func (msg MsgDelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgDelegate) ValidateBasic() types2.Error {
	if msg.DelegatorAddress.Empty() {
		return ErrNilDelegatorAddr(DefaultCodespace)
	}
	if msg.ValidatorAddress.Empty() {
		return ErrNilValidatorAddr(DefaultCodespace)
	}
	if msg.Amount.Amount.LTE(types2.ZeroInt()) {
		return ErrBadDelegationAmount(DefaultCodespace)
	}
	return nil
}

//______________________________________________________________________

type MsgUndelegateByRetrievalAccount struct {
	SecurityAddress  types2.AccAddress   `json:"security_address" yaml:"security_address"`
	DelegatorAddress []types2.AccAddress `json:"delegator_address" yaml:"delegator_address"`
}

func NewMsgUndelegateByRetrievalAccount(secAddr types2.AccAddress, delAddr []types2.AccAddress) MsgUndelegateByRetrievalAccount {

	return MsgUndelegateByRetrievalAccount{
		SecurityAddress:  secAddr,
		DelegatorAddress: delAddr,
	}
}

// nolint
func (msg MsgUndelegateByRetrievalAccount) Route() string { return RouterKey }
func (msg MsgUndelegateByRetrievalAccount) Type() string  { return "undelegateBySecurityAddress" }
func (msg MsgUndelegateByRetrievalAccount) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{msg.SecurityAddress}
}

func (msg MsgUndelegateByRetrievalAccount) ValidateBasic() types2.Error {
	if msg.SecurityAddress.Empty() {
		return ErrNilSecurityAddr(DefaultCodespace)
	}
	for _, delegatorAddress := range msg.DelegatorAddress {
		if delegatorAddress.Empty() {
			return ErrNilDelegatorAddr(DefaultCodespace)
		}
	}
	return nil
}

func (msg MsgUndelegateByRetrievalAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// MsgDelegate - struct for bonding transactions
type MsgBeginRedelegate struct {
	DelegatorAddress    types2.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorSrcAddress types2.ValAddress `json:"con-account_src_address" yaml:"con-account_src_address"`
	ValidatorDstAddress types2.ValAddress `json:"con-account_dst_address" yaml:"con-account_dst_address"`
	Amount              types2.Coin       `json:"amount" yaml:"amount"`
}

func NewMsgBeginRedelegate(delAddr types2.AccAddress, valSrcAddr,
	valDstAddr types2.ValAddress, amount types2.Coin) MsgBeginRedelegate {

	return MsgBeginRedelegate{
		DelegatorAddress:    delAddr,
		ValidatorSrcAddress: valSrcAddr,
		ValidatorDstAddress: valDstAddr,
		Amount:              amount,
	}
}

// nolint
func (msg MsgBeginRedelegate) Route() string { return RouterKey }
func (msg MsgBeginRedelegate) Type() string  { return "redelegation" }
func (msg MsgBeginRedelegate) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{msg.DelegatorAddress}
}

// get the bytes for the message signer to sign on
func (msg MsgBeginRedelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgBeginRedelegate) ValidateBasic() types2.Error {
	if msg.DelegatorAddress.Empty() {
		return ErrNilDelegatorAddr(DefaultCodespace)
	}
	if msg.ValidatorSrcAddress.Empty() {
		return ErrNilValidatorAddr(DefaultCodespace)
	}
	if msg.ValidatorDstAddress.Empty() {
		return ErrNilValidatorAddr(DefaultCodespace)
	}
	if msg.Amount.Amount.LTE(types2.ZeroInt()) {
		return ErrBadSharesAmount(DefaultCodespace)
	}
	return nil
}

// MsgUndelegate - struct for unbonding transactions
type MsgUndelegate struct {
	DelegatorAddress types2.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types2.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	Amount           types2.Coin       `json:"amount" yaml:"amount"`
}

func NewMsgUndelegate(delAddr types2.AccAddress, valAddr types2.ValAddress, amount types2.Coin) MsgUndelegate {
	return MsgUndelegate{
		DelegatorAddress: delAddr,
		ValidatorAddress: valAddr,
		Amount:           amount,
	}
}

// nolint
func (msg MsgUndelegate) Route() string { return RouterKey }
func (msg MsgUndelegate) Type() string  { return "undelegation" }
func (msg MsgUndelegate) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{msg.DelegatorAddress}
}

// get the bytes for the message signer to sign on
func (msg MsgUndelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgUndelegate) ValidateBasic() types2.Error {
	if msg.DelegatorAddress.Empty() {
		return ErrNilDelegatorAddr(DefaultCodespace)
	}
	if msg.ValidatorAddress.Empty() {
		return ErrNilValidatorAddr(DefaultCodespace)
	}
	if msg.Amount.Amount.LTE(types2.ZeroInt()) {
		return ErrBadSharesAmount(DefaultCodespace)
	}
	return nil
}

package types

import (
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"time"

	"github.com/gatechain/crypto"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// Implements Validator interface

// Validator defines the total amount of bond shares and their exchange rate to
// coins. Slashing results in a decrease in the exchange rate, allowing correct
// calculation of future undelegations without iterating over delegators.
// When coins are delegated to this validator, the validator is credited with a
// delegation whose number of bond shares is based on the amount of coins delegated
// divided by the current exchange rate. Voting power can be calculated as total
// bonded shares multiplied by exchange rate.
type Validator struct {
	OperatorAddress         types.ValAddress `json:"operator_address" yaml:"operator_address"`       // address of the validator's operator; bech encoded in JSON
	ConsPubKey              crypto.PubKey    `json:"consensus_pubkey" yaml:"consensus_pubkey"`       // the consensus public key of the validator; bech encoded in JSON
	Jailed                  bool             `json:"jailed" yaml:"jailed"`                           // has the validator been jailed from bonded status?
	Status                  types.BondStatus `json:"status" yaml:"status"`                           // validator status (bonded/unbonding/unbonded)
	Tokens                  types.Int        `json:"tokens" yaml:"tokens"`                           // delegated tokens
	PowerRate               types.Dec        `json:"power_rate" yaml:"power_rate"`                   // poewe rate
	DelegatorShares         types.Dec        `json:"delegator_shares" yaml:"delegator_shares"`       // total shares issued to a validator's delegators
	Description             Description      `json:"description" yaml:"description"`                 // description terms for the validator
	UnbondingHeight         int64            `json:"undelegating_height" yaml:"undelegating_height"` // if unbonding, height at which this validator has begun unbonding
	UnbondingCompletionTime time.Time        `json:"undelegating_time" yaml:"undelegating_time"`     // if unbonding, min time for the validator to complete unbonding
	Commission              Commission       `json:"commission" yaml:"commission"`                   // commission parameters
}

// Validators is a collection of Validator
type Validators []Validator

// this is a helper struct used for JSON de- and encoding only
type bechValidator struct {
	OperatorAddress         types.ValAddress `json:"operator_address" yaml:"operator_address"`       // the bech32 address of the validator's operator
	ConsPubKey              string           `json:"consensus_pubkey" yaml:"consensus_pubkey"`       // the bech32 consensus public key of the validator
	Jailed                  bool             `json:"jailed" yaml:"jailed"`                           // has the validator been jailed from bonded status?
	Status                  types.BondStatus `json:"status" yaml:"status"`                           // validator status (bonded/unbonding/unbonded)
	Tokens                  types.Int        `json:"tokens" yaml:"tokens"`                           // delegated tokens
	PowerRate               types.Dec        `json:"power_rate" yaml:"power_rate"`                   // poewe rate
	DelegatorShares         types.Dec        `json:"delegator_shares" yaml:"delegator_shares"`       // total shares issued to a validator's delegators
	Description             Description      `json:"description" yaml:"description"`                 // description terms for the validator
	UnbondingHeight         int64            `json:"undelegating_height" yaml:"undelegating_height"` // if unbonding, height at which this validator has begun unbonding
	UnbondingCompletionTime time.Time        `json:"undelegating_time" yaml:"undelegating_time"`     // if unbonding, min time for the validator to complete unbonding
	Commission              Commission       `json:"commission" yaml:"commission"`                   // commission parameters
}

// // UnmarshalJSON unmarshals the validator from JSON using Bech32
func (v *Validator) UnmarshalJSON(data []byte) error {
	bv := &bechValidator{}
	if err := codec.Cdc.UnmarshalJSON(data, bv); err != nil {
		return err
	}
	consPubKey, err := types.GetConsPubKeyBech32(bv.ConsPubKey)
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

// Description - description fields for a validator
type Description struct {
	Moniker  string `json:"moniker" yaml:"moniker"`   // name
	Identity string `json:"identity" yaml:"identity"` // optional identity signature (ex. UPort or Keybase)
	Website  string `json:"website" yaml:"website"`   // optional website link
	Details  string `json:"details" yaml:"details"`   // optional details
}

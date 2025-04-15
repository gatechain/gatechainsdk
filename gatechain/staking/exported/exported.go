package exported

import (
	"github.com/gatechain/crypto"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// DelegationI delegation bond for a delegated proof of stake system
type DelegationI interface {
	GetDelegatorAddr() types.AccAddress // delegator framework.AccAddress for the bond
	GetValidatorAddr() types.ValAddress // validator operator address
	GetShares() types.Dec               // amount of validator's shares held in this delegation
}

// ValidatorI expected validator functions
type ValidatorI interface {
	IsJailed() bool                                                   // whether the validator is jailed
	GetMoniker() string                                               // moniker of the validator
	GetStatus() types.BondStatus                                      // status of the validator
	IsBonded() bool                                                   // check if has a bonded status
	IsUnbonded() bool                                                 // check if has status unbonded
	IsUnbonding() bool                                                // check if has status unbonding
	GetOperator() types.ValAddress                                    // operator address to receive/return validators coins
	GetConsPubKey() crypto.PubKey                                     // validation consensus pubkey
	GetConsAddr() types.ConsAddress                                   // validation consensus address
	GetTokens() types.Int                                             // validation tokens
	GetBondedTokens() types.Int                                       // validator bonded tokens
	GetConsensusPower() int64                                         // validation power in tendermint
	GetCommission() types.Dec                                         // validator commission rate
	GetDelegatorShares() types.Dec                                    // total outstanding delegator shares
	TokensFromShares(types.Dec) types.Dec                             // token worth of provided delegator shares
	TokensFromSharesTruncated(types.Dec) types.Dec                    // token worth of provided delegator shares, truncated
	TokensFromSharesRoundUp(types.Dec) types.Dec                      // token worth of provided delegator shares, rounded up
	SharesFromTokens(amt types.Int) (types.Dec, types.Error)          // shares worth of delegator's bond
	SharesFromTokensTruncated(amt types.Int) (types.Dec, types.Error) // truncated shares worth of delegator's bond
}

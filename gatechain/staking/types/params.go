package types

import (
	"bytes"
	"fmt"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	params "github.com/gatechain/gatechainsdk/gatechain/params/subspace"
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
	"time"
)

// Staking params default values
const (
	// DefaultUnbondingTime reflects three weeks in seconds as the default
	// unbonding time.
	// TODO: Justify our choice of default here.
	DefaultUnbondingTime time.Duration = time.Hour * 24 * 7 * 3

	// Default maximum number of bonded validators
	DefaultMaxValidators uint16 = 100

	// Default maximum entries in a UBD/RED pair
	DefaultMaxEntries uint16 = 7

	DefaultPowRate uint16 = 1

	DefaultMaxPowRate uint16 = 2

	DefaultRewardUnitGT uint16 = 18
)

// nolint - Keys for parameter access
var (
	KeyUnbondingTime = []byte("UnbondingTime")
	KeyMaxValidators = []byte("MaxValidators")
	KeyMaxEntries    = []byte("KeyMaxEntries")
	KeyBondDenom     = []byte("BondDenom")
	KeyPowRate       = []byte("PowRate")
	KeyMaxPowRate    = []byte("MaxPowRate")
	KeyRewardUnitGT  = []byte("RewardUnitGT")
)

var _ params.ParamSet = (*Params)(nil)

var DefaultAddMaxPowRate = types2.NewDec(1).Quo(types2.NewDec(100))

// Params defines the high level settings for staking
type Params struct {
	UnbondingTime time.Duration `json:"undelegating_time" yaml:"undelegating_time"` // time duration of unbonding
	MaxValidators uint16        `json:"max_con-accounts" yaml:"max_con-accounts"`   // maximum number of validators (max uint16 = 65535)
	MaxEntries    uint16        `json:"max_entries" yaml:"max_entries"`             // max entries for either unbonding delegation or redelegation (per pair/trio)
	// note: we need to be a bit careful about potential overflow here, since this is user-determined
	BondDenom    string `json:"bond_denom" yaml:"bond_denom"`         // bondable coin denomination
	PowRate      uint16 `json:"pow_rate" yaml:"pow_rate"`             // pow rate for validator
	MaxPowRate   uint16 `json:"max_pow_rate" yaml:"max_pow_rate"`     // pow rate for validator
	RewardUnitGT uint16 `json:"reward_uint_gt" yaml:"reward_uint_gt"` // gt count for rewarduint
}

// NewParams creates a new Params instance
func NewParams(unbondingTime time.Duration, maxValidators, maxEntries uint16,
	bondDenom string, powRate uint16, maxPowRate uint16, rewardUnitGT uint16) Params {

	return Params{
		UnbondingTime: unbondingTime,
		MaxValidators: maxValidators,
		MaxEntries:    maxEntries,
		BondDenom:     bondDenom,
		PowRate:       powRate,
		MaxPowRate:    maxPowRate,
		RewardUnitGT:  rewardUnitGT,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyUnbondingTime, &p.UnbondingTime},
		{KeyMaxValidators, &p.MaxValidators},
		{KeyMaxEntries, &p.MaxEntries},
		{KeyBondDenom, &p.BondDenom},
		{KeyPowRate, &p.PowRate},
		{KeyMaxPowRate, &p.MaxPowRate},
		{KeyRewardUnitGT, &p.RewardUnitGT},
	}
}

// Equal returns a boolean determining if two Param types are identical.
// TODO: This is slower than comparing struct fields directly
func (p Params) Equal(p2 Params) bool {
	bz1 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&p)
	bz2 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&p2)
	return bytes.Equal(bz1, bz2)
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultUnbondingTime, DefaultMaxValidators,
		DefaultMaxEntries, types2.DefaultBondDenom, DefaultPowRate,
		DefaultMaxPowRate, DefaultRewardUnitGT)
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	return fmt.Sprintf(`Params:
 Undelegating Time:    %s
 Max Con-accounts:    %d
 Max Entries:       %d
 Bonded Coin Denom: %s
 Pow Rate:          %d
 Max Pow Rate:      %d`, p.UnbondingTime,
		p.MaxValidators, p.MaxEntries, p.BondDenom, p.PowRate, p.MaxPowRate)
}

// unmarshal the current staking params value from store key or panic
func MustUnmarshalParams(cdc *codec.Codec, value []byte) Params {
	params, err := UnmarshalParams(cdc, value)
	if err != nil {
		panic(err)
	}
	return params
}

// unmarshal the current staking params value from store key
func UnmarshalParams(cdc *codec.Codec, value []byte) (params Params, err error) {
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &params)
	if err != nil {
		return
	}
	return
}

// validate a set of params
func (p Params) Validate() error {
	if p.BondDenom == "" {
		return fmt.Errorf("staking parameter BondDenom can't be an empty string")
	}
	if p.MaxValidators == 0 {
		return fmt.Errorf("staking parameter MaxValidators must be a positive integer")
	}
	if p.MaxPowRate < p.PowRate {
		return fmt.Errorf("staking parameter MaxPowRate must gt PowRate")
	}
	//if sdk.NewDec(int64(p.MaxPowRate)).LT(p.RewordRate) {
	//	return fmt.Errorf("staking parameter MaxPowRate must gt RewordRate")
	//}
	return nil
}

package types

import (
	"fmt"
	"time"

	params "github.com/gatechain/gatechainsdk/gatechain/params/subspace"
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

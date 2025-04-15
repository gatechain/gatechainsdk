package types

import (
	"fmt"
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
	"strings"
)

// historical rewards for a validator
// height is implicit within the store key
// cumulative reward ratio is the sum from the zeroeth period
// until this period of rewards / tokens, per the spec
// The reference count indicates the number of objects
// which might need to reference this historical entry
// at any point.
// ReferenceCount =
//
//	  number of outstanding delegations which ended the associated period (and might need to read that record)
//	+ number of slashes which ended the associated period (and might need to read that record)
//	+ one per validator for the zeroeth period, set on initialization
type ValidatorHistoricalRewards struct {
	CumulativeRewardRatio types2.DecCoins `json:"cumulative_reward_ratio" yaml:"cumulative_reward_ratio"`
	ReferenceCount        uint16          `json:"reference_count" yaml:"reference_count"`
}

// create a new ValidatorHistoricalRewards
func NewValidatorHistoricalRewards(cumulativeRewardRatio types2.DecCoins, referenceCount uint16) ValidatorHistoricalRewards {
	return ValidatorHistoricalRewards{
		CumulativeRewardRatio: cumulativeRewardRatio,
		ReferenceCount:        referenceCount,
	}
}

// current rewards and current period for a validator
// kept as a running counter and incremented each block
// as long as the validator's tokens remain constant
type ValidatorCurrentRewards struct {
	Rewards types2.DecCoins `json:"rewards" yaml:"rewards"` // current rewards
	Period  uint64          `json:"period" yaml:"period"`   // current period
}

// create a new ValidatorCurrentRewards
func NewValidatorCurrentRewards(rewards types2.DecCoins, period uint64) ValidatorCurrentRewards {
	return ValidatorCurrentRewards{
		Rewards: rewards,
		Period:  period,
	}
}

// accumulated commission for a validator
// kept as a running counter, can be withdrawn at any time
type ValidatorAccumulatedCommission = types2.DecCoins

// return the initial accumulated commission (zero)
func InitialValidatorAccumulatedCommission() ValidatorAccumulatedCommission {
	return ValidatorAccumulatedCommission{}
}

// validator slash event
// height is implicit within the store key
// needed to calculate appropriate amounts of staking token
// for delegations which withdraw after a slash has occurred
type ValidatorSlashEvent struct {
	ValidatorPeriod uint64     `json:"validator_period" yaml:"validator_period"` // period when the slash occurred
	Fraction        types2.Dec `json:"fraction" yaml:"fraction"`                 // slash fraction
}

// create a new ValidatorSlashEvent
func NewValidatorSlashEvent(validatorPeriod uint64, fraction types2.Dec) ValidatorSlashEvent {
	return ValidatorSlashEvent{
		ValidatorPeriod: validatorPeriod,
		Fraction:        fraction,
	}
}

func (vs ValidatorSlashEvent) String() string {
	return fmt.Sprintf(`Period:   %d
Fraction: %s`, vs.ValidatorPeriod, vs.Fraction)
}

// ValidatorSlashEvents is a collection of ValidatorSlashEvent
type ValidatorSlashEvents []ValidatorSlashEvent

func (vs ValidatorSlashEvents) String() string {
	out := "Validator Slash Events:\n"
	for i, sl := range vs {
		out += fmt.Sprintf(`  Slash %d:
    Period:   %d
    Fraction: %s
`, i, sl.ValidatorPeriod, sl.Fraction)
	}
	return strings.TrimSpace(out)
}

// outstanding (un-withdrawn) rewards for a validator
// inexpensive to track, allows simple sanity checks
type ValidatorOutstandingRewards = types2.DecCoins

type ValidatorUpdate struct {
	Power   int64
	Address types2.ValAddress
}

func NewValidatorUpdate(power int64, address types2.ValAddress) ValidatorUpdate {
	return ValidatorUpdate{
		Power:   power,
		Address: address,
	}
}

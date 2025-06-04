package types

import (
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// accumulated commission for a validator
// kept as a running counter, can be withdrawn at any time
type ValidatorAccumulatedCommission = types.DecCoins

// outstanding (un-withdrawn) rewards for a validator
// inexpensive to track, allows simple sanity checks
type ValidatorOutstandingRewards = types.DecCoins

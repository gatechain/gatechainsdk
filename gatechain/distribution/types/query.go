package types

import (
	"fmt"
	"strings"

	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// QueryDelegatorTotalRewardsResponse defines the properties of
// QueryDelegatorTotalRewards query's response.
type QueryDelegatorTotalRewardsResponse struct {
	Rewards []DelegationDelegatorReward `json:"rewards" yaml:"rewards"`
	Total   types.DecCoins              `json:"total" yaml:"total"`
}

func (res QueryDelegatorTotalRewardsResponse) String() string {
	out := "Delegator Total Rewards:\n"
	out += "  Rewards:"
	for _, reward := range res.Rewards {
		out += fmt.Sprintf(`
	Con-account: %s
	Reward: %s`, reward.ValidatorAddress, reward.Reward)
	}
	out += fmt.Sprintf("\n  Total: %s\n", res.Total)
	return strings.TrimSpace(out)
}

// DelegationDelegatorReward defines the properties
// of a delegator's delegation reward.
type DelegationDelegatorReward struct {
	ValidatorAddress types.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	Reward           types.DecCoins   `json:"reward" yaml:"reward"`
}

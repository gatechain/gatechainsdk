package types

import (
	"fmt"
	"strings"

	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
)

// QueryDelegatorTotalRewardsResponse defines the properties of
// QueryDelegatorTotalRewards query's response.
type QueryDelegatorTotalRewardsResponse struct {
	Rewards []DelegationDelegatorReward `json:"rewards" yaml:"rewards"`
	Total   types2.DecCoins             `json:"total" yaml:"total"`
}

// NewQueryDelegatorTotalRewardsResponse constructs a QueryDelegatorTotalRewardsResponse
func NewQueryDelegatorTotalRewardsResponse(rewards []DelegationDelegatorReward,
	total types2.DecCoins) QueryDelegatorTotalRewardsResponse {
	return QueryDelegatorTotalRewardsResponse{Rewards: rewards, Total: total}
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
	ValidatorAddress types2.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	Reward           types2.DecCoins   `json:"reward" yaml:"reward"`
}

// NewDelegationDelegatorReward constructs a DelegationDelegatorReward.
func NewDelegationDelegatorReward(valAddr types2.ValAddress,
	reward types2.DecCoins) DelegationDelegatorReward {
	return DelegationDelegatorReward{ValidatorAddress: valAddr, Reward: reward}
}

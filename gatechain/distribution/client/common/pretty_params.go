package common

import (
	"encoding/json"
	"fmt"
)

// Convenience struct for CLI output
type PrettyParams struct {
	CommunityTax          json.RawMessage `json:"community_tax"`
	WithdrawAddrEnabled   json.RawMessage `json:"withdraw_addr_enabled"`
	FirstCommitteeReward  json.RawMessage `json:"first_committee_reward"`
	SecondCommitteeReward json.RawMessage `json:"second_committee_reward"`
	ThirdCommitteeReward  json.RawMessage `json:"third_committee_reward"`
}

// Construct a new PrettyParams
func NewPrettyParams(communityTax json.RawMessage, withdrawAddrEnabled json.RawMessage, firstCommitteeReward json.RawMessage,
	secondCommitteeReward json.RawMessage, thirdCommitteeReward json.RawMessage) PrettyParams {
	return PrettyParams{
		CommunityTax:          communityTax,
		WithdrawAddrEnabled:   withdrawAddrEnabled,
		FirstCommitteeReward:  firstCommitteeReward,
		SecondCommitteeReward: secondCommitteeReward,
		ThirdCommitteeReward:  thirdCommitteeReward,
	}
}

func (pp PrettyParams) String() string {
	return fmt.Sprintf(`Distribution Params:
  Community Tax:          %s
  Withdraw Addr Enabled:  %s
  First CommitteeReward:  %s
  Second CommitteeReward:  %s
  Third CommitteeReward:  %s`, pp.CommunityTax, pp.WithdrawAddrEnabled, pp.FirstCommitteeReward,
		pp.SecondCommitteeReward, pp.ThirdCommitteeReward)

}

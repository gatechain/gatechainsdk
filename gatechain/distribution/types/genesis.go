package types

import (
	"fmt"
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
)

// the address for where distributions rewards are withdrawn to by default
// this struct is only used at genesis to feed in default withdraw addresses
type DelegatorWithdrawInfo struct {
	DelegatorAddress types2.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	WithdrawAddress  types2.AccAddress `json:"withdraw_address" yaml:"withdraw_address"`
}

// used for import/export via genesis json
type ValidatorOutstandingRewardsRecord struct {
	ValidatorAddress   types2.ValAddress `json:"validator_address" yaml:"validator_address"`
	OutstandingRewards types2.DecCoins   `json:"outstanding_rewards" yaml:"outstanding_rewards"`
}

// used for import / export via genesis json
type ValidatorAccumulatedCommissionRecord struct {
	ValidatorAddress types2.ValAddress              `json:"validator_address" yaml:"validator_address"`
	Accumulated      ValidatorAccumulatedCommission `json:"accumulated" yaml:"accumulated"`
}

// used for import / export via genesis json
type ValidatorHistoricalRewardsRecord struct {
	ValidatorAddress types2.ValAddress          `json:"validator_address" yaml:"validator_address"`
	Period           uint64                     `json:"period" yaml:"period"`
	Rewards          ValidatorHistoricalRewards `json:"rewards" yaml:"rewards"`
}

// used for import / export via genesis json
type ValidatorCurrentRewardsRecord struct {
	ValidatorAddress types2.ValAddress       `json:"validator_address" yaml:"validator_address"`
	Rewards          ValidatorCurrentRewards `json:"rewards" yaml:"rewards"`
}

// used for import / export via genesis json
type DelegatorStartingInfoRecord struct {
	DelegatorAddress types2.AccAddress     `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types2.ValAddress     `json:"validator_address" yaml:"validator_address"`
	StartingInfo     DelegatorStartingInfo `json:"starting_info" yaml:"starting_info"`
}

// used for import / export via genesis json
type ValidatorSlashEventRecord struct {
	ValidatorAddress types2.ValAddress   `json:"validator_address" yaml:"validator_address"`
	Height           uint64              `json:"height" yaml:"height"`
	Period           uint64              `json:"period" yaml:"period"`
	Event            ValidatorSlashEvent `json:"validator_slash_event" yaml:"validator_slash_event"`
}

// GenesisState - all distribution state that must be provided at genesis
type GenesisState struct {
	FeePool                         FeePool                                `json:"fee_pool" yaml:"fee_pool"`
	CommunityTax                    types2.Dec                             `json:"community_tax" yaml:"community_tax"`
	FirstCommitteeReward            types2.Dec                             `json:"first_committee_reward" yaml:"first_committee_reward"`
	SecondCommitteeReward           types2.Dec                             `json:"second_committee_reward" yaml:"second_committee_reward"`
	ThirdCommitteeReward            types2.Dec                             `json:"third_committee_reward" yaml:"third_committee_reward"`
	WithdrawAddrEnabled             bool                                   `json:"withdraw_addr_enabled" yaml:"withdraw_addr_enabled"`
	DelegatorWithdrawInfos          []DelegatorWithdrawInfo                `json:"delegator_withdraw_infos" yaml:"delegator_withdraw_infos"`
	PreviousProposer                types2.ValAddress                      `json:"previous_proposer" yaml:"previous_proposer"`
	OutstandingRewards              []ValidatorOutstandingRewardsRecord    `json:"outstanding_rewards" yaml:"outstanding_rewards"`
	ValidatorAccumulatedCommissions []ValidatorAccumulatedCommissionRecord `json:"validator_accumulated_commissions" yaml:"validator_accumulated_commissions"`
	ValidatorHistoricalRewards      []ValidatorHistoricalRewardsRecord     `json:"validator_historical_rewards" yaml:"validator_historical_rewards"`
	ValidatorCurrentRewards         []ValidatorCurrentRewardsRecord        `json:"validator_current_rewards" yaml:"validator_current_rewards"`
	DelegatorStartingInfos          []DelegatorStartingInfoRecord          `json:"delegator_starting_infos" yaml:"delegator_starting_infos"`
	ValidatorSlashEvents            []ValidatorSlashEventRecord            `json:"validator_slash_events" yaml:"validator_slash_events"`
}

func NewGenesisState(feePool FeePool, communityTax types2.Dec,
	withdrawAddrEnabled bool, dwis []DelegatorWithdrawInfo, pp types2.ValAddress, r []ValidatorOutstandingRewardsRecord,
	acc []ValidatorAccumulatedCommissionRecord, historical []ValidatorHistoricalRewardsRecord,
	cur []ValidatorCurrentRewardsRecord, dels []DelegatorStartingInfoRecord,
	slashes []ValidatorSlashEventRecord, firstCommitteeReward, secondCommitteeReward, thirdCommitteeReward types2.Dec) GenesisState {

	return GenesisState{
		FeePool:                         feePool,
		CommunityTax:                    communityTax,
		WithdrawAddrEnabled:             withdrawAddrEnabled,
		DelegatorWithdrawInfos:          dwis,
		PreviousProposer:                pp,
		OutstandingRewards:              r,
		ValidatorAccumulatedCommissions: acc,
		ValidatorHistoricalRewards:      historical,
		ValidatorCurrentRewards:         cur,
		DelegatorStartingInfos:          dels,
		ValidatorSlashEvents:            slashes,
		FirstCommitteeReward:            firstCommitteeReward,
		SecondCommitteeReward:           secondCommitteeReward,
		ThirdCommitteeReward:            thirdCommitteeReward,
	}
}

// get raw genesis raw message for testing
func DefaultGenesisState() GenesisState {
	return GenesisState{
		FeePool:                         InitialFeePool(),
		CommunityTax:                    types2.NewDecWithPrec(0, 2), // 2%
		WithdrawAddrEnabled:             true,
		DelegatorWithdrawInfos:          []DelegatorWithdrawInfo{},
		PreviousProposer:                nil,
		OutstandingRewards:              []ValidatorOutstandingRewardsRecord{},
		ValidatorAccumulatedCommissions: []ValidatorAccumulatedCommissionRecord{},
		ValidatorHistoricalRewards:      []ValidatorHistoricalRewardsRecord{},
		ValidatorCurrentRewards:         []ValidatorCurrentRewardsRecord{},
		DelegatorStartingInfos:          []DelegatorStartingInfoRecord{},
		ValidatorSlashEvents:            []ValidatorSlashEventRecord{},
		FirstCommitteeReward:            types2.NewDecWithPrec(40, 2),
		SecondCommitteeReward:           types2.NewDecWithPrec(35, 2),
		ThirdCommitteeReward:            types2.NewDecWithPrec(25, 2),
	}
}

// ValidateGenesis validates the genesis state of distribution genesis input
func ValidateGenesis(data GenesisState) error {
	if data.CommunityTax.IsNegative() || data.CommunityTax.GT(types2.OneDec()) {
		return fmt.Errorf("mint parameter CommunityTax should non-negative and "+
			"less than one, is %s", data.CommunityTax.String())
	}
	if data.FirstCommitteeReward.IsNegative() {
		return fmt.Errorf("mint parameter FirstCommitteeReward should be positive, is %s",
			data.FirstCommitteeReward.String())
	}
	if data.SecondCommitteeReward.IsNegative() {
		return fmt.Errorf("mint parameter SecondCommitteeReward should be positive, is %s",
			data.SecondCommitteeReward.String())
	}
	if data.ThirdCommitteeReward.IsNegative() {
		return fmt.Errorf("mint parameter ThirdCommitteeReward should be positive, is %s",
			data.ThirdCommitteeReward.String())
	}

	return data.FeePool.ValidateGenesis()
}

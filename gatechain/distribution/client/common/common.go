package common

import (
	"fmt"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	types2 "github.com/gatechain/gatechainsdk/gatechain/distribution/types"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// QueryParams actually queries distribution params.
func QueryParams(cliCtx context.NodeVaultQuerierImpl, queryRoute string) (PrettyParams, error) {
	route := fmt.Sprintf("custom/%s/params/%s", queryRoute, types2.ParamCommunityTax)

	retCommunityTax, _, err := cliCtx.QueryWithData(route, []byte{})
	if err != nil {
		return PrettyParams{}, err
	}

	route = fmt.Sprintf("custom/%s/params/%s", queryRoute, types2.ParamWithdrawAddrEnabled)
	retWithdrawAddrEnabled, _, err := cliCtx.QueryWithData(route, []byte{})
	if err != nil {
		return PrettyParams{}, err
	}
	route = fmt.Sprintf("custom/%s/params/%s", queryRoute, types2.ParamFirstCommitteeReward)
	retFirstCommitteeReward, _, err := cliCtx.QueryWithData(route, []byte{})
	if err != nil {
		return PrettyParams{}, err
	}
	route = fmt.Sprintf("custom/%s/params/%s", queryRoute, types2.ParamSecondCommitteeReward)
	retSecondCommitteeReward, _, err := cliCtx.QueryWithData(route, []byte{})
	if err != nil {
		return PrettyParams{}, err
	}
	route = fmt.Sprintf("custom/%s/params/%s", queryRoute, types2.ParamThirdCommitteeReward)
	retThirdCommitteeReward, _, err := cliCtx.QueryWithData(route, []byte{})
	if err != nil {
		return PrettyParams{}, err
	}

	return NewPrettyParams(
		retCommunityTax, retWithdrawAddrEnabled,
		retFirstCommitteeReward, retSecondCommitteeReward, retThirdCommitteeReward,
	), nil
}

// QueryDelegatorTotalRewards queries delegator total rewards.
func QueryDelegatorTotalRewards(cliCtx context.NodeVaultQuerierImpl, queryRoute, delAddr string) ([]byte, error) {
	delegatorAddr, err := types.AccAddressFromBech32(delAddr)
	if err != nil {
		return nil, err
	}

	res, _, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", queryRoute, types2.QueryDelegatorTotalRewards),
		cliCtx.Codec.MustMarshalJSON(types2.NewQueryDelegatorParams(delegatorAddr)),
	)
	return res, err
}

// QueryDelegationRewards queries a delegation rewards.
func QueryDelegationRewards(cliCtx context.NodeVaultQuerierImpl, queryRoute, delAddr, valAddr string) ([]byte, error) {
	delegatorAddr, err := types.AccAddressFromBech32(delAddr)
	if err != nil {
		return nil, err
	}

	validatorAddr, err := types.ValAddressFromBech32(valAddr)
	if err != nil {
		return nil, err
	}

	res, _, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", queryRoute, types2.QueryDelegationRewards),
		cliCtx.Codec.MustMarshalJSON(types2.NewQueryDelegationRewardsParams(delegatorAddr, validatorAddr)),
	)
	return res, err
}

// QueryDelegatorValidators returns delegator's list of validators
// it submitted delegations to.
func QueryDelegatorValidators(cliCtx context.NodeVaultQuerierImpl, queryRoute string, delegatorAddr types.AccAddress) ([]byte, error) {
	res, _, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", queryRoute, types2.QueryDelegatorValidators),
		cliCtx.Codec.MustMarshalJSON(types2.NewQueryDelegatorParams(delegatorAddr)),
	)
	return res, err
}

// QueryValidatorCommission returns a validator's commission.
func QueryValidatorCommission(cliCtx context.NodeVaultQuerierImpl, queryRoute string, validatorAddr types.ValAddress) ([]byte, error) {
	res, _, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", queryRoute, types2.QueryValidatorCommission),
		cliCtx.Codec.MustMarshalJSON(types2.NewQueryValidatorCommissionParams(validatorAddr)),
	)
	return res, err
}

// WithdrawAllDelegatorRewards builds a multi-message slice to be used
// to withdraw all delegations rewards for the given delegator.
func WithdrawAllDelegatorRewards(cliCtx context.NodeVaultQuerierImpl, queryRoute string, delegatorAddr types.AccAddress) ([]types.Msg, error) {
	// retrieve the comprehensive list of all validators which the
	// delegator had submitted delegations to
	bz, err := QueryDelegatorValidators(cliCtx, queryRoute, delegatorAddr)
	if err != nil {
		return nil, err
	}

	var validators []types.ValAddress
	if err := cliCtx.Codec.UnmarshalJSON(bz, &validators); err != nil {
		return nil, err
	}

	// build multi-message transaction
	var msgs []types.Msg
	for _, valAddr := range validators {
		msg := types2.NewMsgWithdrawDelegatorReward(delegatorAddr, valAddr)
		if err := msg.ValidateBasic(); err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}

	return msgs, nil
}

// WithdrawValidatorRewardsAndCommission builds a two-message message slice to be
// used to withdraw both validation's commission and self-delegation reward.
func WithdrawValidatorRewardsAndCommission(validatorAddr types.ValAddress) ([]types.Msg, error) {
	commissionMsg := types2.NewMsgWithdrawValidatorCommission(validatorAddr)
	if err := commissionMsg.ValidateBasic(); err != nil {
		return nil, err
	}

	// build and validate MsgWithdrawDelegatorReward
	rewardMsg := types2.NewMsgWithdrawDelegatorReward(types.AccAddress(validatorAddr.Bytes()), validatorAddr)
	if err := rewardMsg.ValidateBasic(); err != nil {
		return nil, err
	}

	return []types.Msg{commissionMsg, rewardMsg}, nil
}

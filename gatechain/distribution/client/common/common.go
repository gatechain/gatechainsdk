package common

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// QueryParams actually queries distribution params.
func QueryParams(cliCtx context.NodeVaultQuerierImpl, queryRoute string) (PrettyParams, error) {
	route := fmt.Sprintf("custom/%s/params/%s", queryRoute, types.ParamCommunityTax)

	retCommunityTax, _, err := cliCtx.QueryWithData(route, []byte{})
	if err != nil {
		return PrettyParams{}, err
	}

	route = fmt.Sprintf("custom/%s/params/%s", queryRoute, types.ParamWithdrawAddrEnabled)
	retWithdrawAddrEnabled, _, err := cliCtx.QueryWithData(route, []byte{})
	if err != nil {
		return PrettyParams{}, err
	}
	route = fmt.Sprintf("custom/%s/params/%s", queryRoute, types.ParamFirstCommitteeReward)
	retFirstCommitteeReward, _, err := cliCtx.QueryWithData(route, []byte{})
	if err != nil {
		return PrettyParams{}, err
	}
	route = fmt.Sprintf("custom/%s/params/%s", queryRoute, types.ParamSecondCommitteeReward)
	retSecondCommitteeReward, _, err := cliCtx.QueryWithData(route, []byte{})
	if err != nil {
		return PrettyParams{}, err
	}
	route = fmt.Sprintf("custom/%s/params/%s", queryRoute, types.ParamThirdCommitteeReward)
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
	delegatorAddr, err := sdk.AccAddressFromBech32(delAddr)
	if err != nil {
		return nil, err
	}

	res, _, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryDelegatorTotalRewards),
		cliCtx.Codec.MustMarshalJSON(types.NewQueryDelegatorParams(delegatorAddr)),
	)
	return res, err
}

// QueryDelegationRewards queries a delegation rewards.
func QueryDelegationRewards(cliCtx context.NodeVaultQuerierImpl, queryRoute, delAddr, valAddr string) ([]byte, error) {
	delegatorAddr, err := sdk.AccAddressFromBech32(delAddr)
	if err != nil {
		return nil, err
	}

	validatorAddr, err := sdk.ValAddressFromBech32(valAddr)
	if err != nil {
		return nil, err
	}

	res, _, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryDelegationRewards),
		cliCtx.Codec.MustMarshalJSON(types.NewQueryDelegationRewardsParams(delegatorAddr, validatorAddr)),
	)
	return res, err
}

// QueryDelegatorValidators returns delegator's list of validators
// it submitted delegations to.
func QueryDelegatorValidators(cliCtx context.NodeVaultQuerierImpl, queryRoute string, delegatorAddr sdk.AccAddress) ([]byte, error) {
	res, _, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryDelegatorValidators),
		cliCtx.Codec.MustMarshalJSON(types.NewQueryDelegatorParams(delegatorAddr)),
	)
	return res, err
}

// QueryValidatorCommission returns a validator's commission.
func QueryValidatorCommission(cliCtx context.NodeVaultQuerierImpl, queryRoute string, validatorAddr sdk.ValAddress) ([]byte, error) {
	res, _, err := cliCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryValidatorCommission),
		cliCtx.Codec.MustMarshalJSON(types.NewQueryValidatorCommissionParams(validatorAddr)),
	)
	return res, err
}

// WithdrawAllDelegatorRewards builds a multi-message slice to be used
// to withdraw all delegations rewards for the given delegator.
func WithdrawAllDelegatorRewards(cliCtx context.NodeVaultQuerierImpl, queryRoute string, delegatorAddr sdk.AccAddress) ([]sdk.Msg, error) {
	// retrieve the comprehensive list of all validators which the
	// delegator had submitted delegations to
	bz, err := QueryDelegatorValidators(cliCtx, queryRoute, delegatorAddr)
	if err != nil {
		return nil, err
	}

	var validators []sdk.ValAddress
	if err := cliCtx.Codec.UnmarshalJSON(bz, &validators); err != nil {
		return nil, err
	}

	// build multi-message transaction
	var msgs []sdk.Msg
	for _, valAddr := range validators {
		msg := types.NewMsgWithdrawDelegatorReward(delegatorAddr, valAddr)
		msgs = append(msgs, msg)
	}

	return msgs, nil
}

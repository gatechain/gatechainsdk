package cli

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/staking/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// QueryValidatorUnbondingDelegations implements the query all unbonding delegatations from a validator command.
func QueryValidatorUnbondingDelegations(ctx *context.NodeVaultQuerierImpl, ValidatorAddr string) (types.UnbondingDelegations, error) {

	ValidatorAccAddr, err := sdk.ValAddressFromBech32(ValidatorAddr)
	if err != nil {
		return types.UnbondingDelegations{}, err
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryValidatorParams(ValidatorAccAddr))
	if err != nil {
		return types.UnbondingDelegations{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryValidatorUnbondingDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		return types.UnbondingDelegations{}, err
	}

	var ubds types.UnbondingDelegations
	ctx.GetCodec().MustUnmarshalJSON(res, &ubds)
	ctx.PrintOutput(ubds)
	return ubds, nil
}

// QueryValidatorRedelegations implements the query all redelegatations
// from a validator command.
func QueryValidatorRedelegations(ctx *context.NodeVaultQuerierImpl, SrcValidatorAddr string) (types.RedelegationResponses, error) {
	SrcValidatorValAddr, err := sdk.ValAddressFromBech32(SrcValidatorAddr)
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.QueryRedelegationParams{SrcValidatorAddr: SrcValidatorValAddr})
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryRedelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	var resp types.RedelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		return types.RedelegationResponses{}, err
	}
	ctx.PrintOutput(resp)
	return resp, nil

}

// QueryDelegation the query delegation command.
func QueryDelegation(ctx *context.NodeVaultQuerierImpl, delegatorAddr, validatorAddr string) (types.DelegationResponse, error) {

	delegatorAccAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return types.DelegationResponse{}, err
	}

	validatorAccAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return types.DelegationResponse{}, err
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryBondsParams(delegatorAccAddr, validatorAccAddr))
	if err != nil {
		return types.DelegationResponse{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryDelegation)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		return types.DelegationResponse{}, err
	}

	var resp types.DelegationResponse
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		return types.DelegationResponse{}, err
	}
	ctx.PrintOutput(resp)
	return resp, nil
}

// QueryDelegations implements the command to query all the delegations
// made from one delegator.
func QueryDelegations(ctx *context.NodeVaultQuerierImpl, delegatorAddr string) (types.DelegationResponses, error) {
	delegatorAccAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return types.DelegationResponses{}, err
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryDelegatorParams(delegatorAccAddr))
	if err != nil {
		return types.DelegationResponses{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryDelegatorDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		return types.DelegationResponses{}, err
	}

	var resp types.DelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		return types.DelegationResponses{}, err
	}
	ctx.PrintOutput(resp)
	return resp, nil
}

// QueryValidatorDelegations implements the command to query all the
// delegations to a specific validator.
// , cdc *codec.Codec
func QueryValidatorDelegations(ctx *context.NodeVaultQuerierImpl, delegatorAddr string) (types.DelegationResponses, error) {

	delegatorAccAddr, err := sdk.ValAddressFromBech32(delegatorAddr)
	if err != nil {
		return types.DelegationResponses{}, err
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryValidatorParams(delegatorAccAddr))
	if err != nil {
		return types.DelegationResponses{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryValidatorDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		return types.DelegationResponses{}, err
	}

	var resp types.DelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		return types.DelegationResponses{}, err
	}
	ctx.PrintOutput(resp)
	return resp, nil
}

// QueryUnbondingDelegation implements the command to query a single
// unbonding-delegation record.
func QueryUnbondingDelegation(ctx *context.NodeVaultQuerierImpl, delegatorAddr, validatorAddr string) (types.UnbondingDelegation, error) {

	validatorAccAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return types.UnbondingDelegation{}, err
	}

	delAccAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return types.UnbondingDelegation{}, err
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryBondsParams(delAccAddr, validatorAccAddr))
	if err != nil {
		return types.UnbondingDelegation{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryUnbondingDelegation)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		return types.UnbondingDelegation{}, err
	}

	var resp types.UnbondingDelegation
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		return types.UnbondingDelegation{}, err
	}
	ctx.PrintOutput(resp)
	return resp, nil
}

// QueryUnbondingDelegations implements the command to query all the
// unbonding-delegation records for a delegator.
func QueryUnbondingDelegations(ctx *context.NodeVaultQuerierImpl, delegatorAddr string) (types.UnbondingDelegations, error) {

	delegatorAccAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return types.UnbondingDelegations{}, err
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryDelegatorParams(delegatorAccAddr))
	if err != nil {
		return types.UnbondingDelegations{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryDelegatorUnbondingDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		return types.UnbondingDelegations{}, err
	}

	var resp types.UnbondingDelegations
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		return types.UnbondingDelegations{}, err
	}
	ctx.PrintOutput(resp)
	return resp, nil
}

func QueryRedelegation(ctx *context.NodeVaultQuerierImpl, args []string) (types.RedelegationResponses, error) {

	delAddr, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	valSrcAddr, err := sdk.ValAddressFromBech32(args[1])
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	valDstAddr, err := sdk.ValAddressFromBech32(args[2])
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryRedelegationParams(delAddr, valSrcAddr, valDstAddr))
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryRedelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	var resp types.RedelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		return types.RedelegationResponses{}, err
	}
	ctx.PrintOutput(resp)
	return resp, nil
}

// QueryRedelegations implements the command to query all the
// redelegation records for a delegator.
func QueryRedelegations(ctx *context.NodeVaultQuerierImpl, delegatorAddr string) (types.RedelegationResponses, error) {
	delAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.QueryRedelegationParams{DelegatorAddr: delAddr})
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryRedelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		return types.RedelegationResponses{}, err
	}

	var resp types.RedelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		return types.RedelegationResponses{}, err
	}
	ctx.PrintOutput(resp)
	return resp, nil
}

// QueryPool implements the pool query command.
func QueryPool(ctx *context.NodeVaultQuerierImpl) (types.Pool, error) {

	bz, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/pool", types.ModuleName), nil)
	if err != nil {
		return types.Pool{}, err
	}

	var pool types.Pool
	if err := ctx.GetCodec().UnmarshalJSON(bz, &pool); err != nil {
		return types.Pool{}, err
	}
	ctx.PrintOutput(pool)
	return pool, nil
}

// QueryParams implements the params query command.
func QueryParams(ctx *context.NodeVaultQuerierImpl) (types.Params, error) {
	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryParameters)
	bz, _, err := ctx.QueryWithData(route, nil)
	if err != nil {
		return types.Params{}, err
	}

	var params types.Params
	ctx.GetCodec().MustUnmarshalJSON(bz, &params)
	ctx.PrintOutput(params)

	return types.Params{}, nil
}

package cli

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/staking/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// QueryValidatorUnbondingDelegations implements the query all unbonding delegatations from a validator command.
func QueryValidatorUnbondingDelegations(ctx *context.NodeVaultQuerierImpl, ValidatorAddr string) {

	ValidatorAccAddr, err := sdk.ValAddressFromBech32(ValidatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryValidatorParams(ValidatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryValidatorUnbondingDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return
	}

	var ubds types.UnbondingDelegations
	ctx.GetCodec().MustUnmarshalJSON(res, &ubds)
	ctx.PrintOutput(ubds)
}

// QueryValidatorRedelegations implements the query all redelegatations
// from a validator command.
func QueryValidatorRedelegations(ctx *context.NodeVaultQuerierImpl, SrcValidatorAddr string) {
	SrcValidatorValAddr, err := sdk.ValAddressFromBech32(SrcValidatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.QueryRedelegationParams{SrcValidatorAddr: SrcValidatorValAddr})
	if err != nil {
		fmt.Println(err)
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryRedelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp types.RedelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
		return
	}
	ctx.PrintOutput(resp)
}

// QueryDelegation the query delegation command.
func QueryDelegation(ctx *context.NodeVaultQuerierImpl, delegatorAddr, validatorAddr string) {

	delegatorAccAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	validatorAccAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryBondsParams(delegatorAccAddr, validatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryDelegation)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp types.DelegationResponse
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
		return
	}

	ctx.PrintOutput(resp)
}

// QueryDelegations implements the command to query all the delegations
// made from one delegator.
func QueryDelegations(ctx *context.NodeVaultQuerierImpl, delegatorAddr string) {
	delegatorAccAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryDelegatorParams(delegatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryDelegatorDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp types.DelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
		return
	}
	ctx.PrintOutput(resp)
}

// QueryValidatorDelegations implements the command to query all the
// delegations to a specific validator.
// , cdc *codec.Codec
func QueryValidatorDelegations(ctx *context.NodeVaultQuerierImpl, delegatorAddr string) {

	delegatorAccAddr, err := sdk.ValAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryValidatorParams(delegatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryValidatorDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp types.DelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
		return
	}
	ctx.PrintOutput(resp)
}

// QueryUnbondingDelegation implements the command to query a single
// unbonding-delegation record.
func QueryUnbondingDelegation(ctx *context.NodeVaultQuerierImpl, delegatorAddr, validatorAddr string) {

	validatorAccAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	delAccAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryBondsParams(delAccAddr, validatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryUnbondingDelegation)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp types.UnbondingDelegation
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
		return
	}
	ctx.PrintOutput(resp)
}

// QueryUnbondingDelegations implements the command to query all the
// unbonding-delegation records for a delegator.
func QueryUnbondingDelegations(ctx *context.NodeVaultQuerierImpl, delegatorAddr string) {

	delegatorAccAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryDelegatorParams(delegatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryDelegatorUnbondingDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp types.UnbondingDelegations
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
		return
	}
	ctx.PrintOutput(resp)
}

func QueryRedelegation(ctx *context.NodeVaultQuerierImpl, args []string) {

	delAddr, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	valSrcAddr, err := sdk.ValAddressFromBech32(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	valDstAddr, err := sdk.ValAddressFromBech32(args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.NewQueryRedelegationParams(delAddr, valSrcAddr, valDstAddr))
	if err != nil {
		fmt.Println(err)
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryRedelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp types.RedelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
		return
	}
	ctx.PrintOutput(resp)
}

// QueryRedelegations implements the command to query all the
// redelegation records for a delegator.
func QueryRedelegations(ctx *context.NodeVaultQuerierImpl, delegatorAddr string) {
	delAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	bz, err := ctx.GetCodec().MarshalJSON(types.QueryRedelegationParams{DelegatorAddr: delAddr})
	if err != nil {
		fmt.Println(err)
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryRedelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp types.RedelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
		return
	}
	ctx.PrintOutput(resp)
}

// QueryPool implements the pool query command.
func QueryPool(ctx *context.NodeVaultQuerierImpl) {

	bz, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/pool", types.ModuleName), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	var pool types.Pool
	if err := ctx.GetCodec().UnmarshalJSON(bz, &pool); err != nil {
		fmt.Println(err)
		return
	}

	ctx.PrintOutput(pool)
}

// QueryParams implements the params query command.
func QueryParams(ctx *context.NodeVaultQuerierImpl) {
	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryParameters)
	bz, _, err := ctx.QueryWithData(route, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	var params types.Params
	ctx.GetCodec().MustUnmarshalJSON(bz, &params)
	ctx.PrintOutput(params)
}

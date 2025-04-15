package cli

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/context"
	types3 "github.com/gatechain/gatechainsdk/gatechain/staking/types"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

const (
	flagPage  = "page"
	flagLimit = "limit"
)

var (
	API_TOKEN = "a7de43d74b7323e82f6c561336b0d5b01d7f452e1cf4af23cd3e30f8b8582383"
	endpoint  = "http://124.243.187.49:80"
)

// GetCmdQueryValidatorUnbondingDelegations implements the query all unbonding delegatations from a validator command.
func GetCmdQueryValidatorUnbondingDelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, ValidatorAddr string) {

	ValidatorAccAddr, err := types.ValAddressFromBech32(ValidatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	bz, err := ctx.GetCodec().MarshalJSON(types3.NewQueryValidatorParams(ValidatorAccAddr))
	if err != nil {
		fmt.Println(err)
	}

	route := fmt.Sprintf("custom/%s/%s", queryRoute, types3.QueryValidatorUnbondingDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
	}

	var ubds types3.UnbondingDelegations
	ctx.GetCodec().MustUnmarshalJSON(res, &ubds)
	ctx.PrintOutput(ubds)
}

// GetCmdQueryValidatorRedelegations implements the query all redelegatations
// from a validator command.
func GetCmdQueryValidatorRedelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, SrcValidatorAddr string) {
	SrcValidatorValAddr, err := types.ValAddressFromBech32(SrcValidatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	bz, err := ctx.GetCodec().MarshalJSON(types3.QueryRedelegationParams{SrcValidatorAddr: SrcValidatorValAddr})
	if err != nil {
		fmt.Println(err)
	}

	route := fmt.Sprintf("custom/%s/%s", queryRoute, types3.QueryRedelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
	}

	var resp types3.RedelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
	}
	ctx.PrintOutput(resp)
}

// GetCmdQueryDelegation the query delegation command.
func GetCmdQueryDelegation(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr, validatorAddr string) {

	delegatorAccAddr, err := types.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	validatorAccAddr, err := types.ValAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	bz, err := ctx.GetCodec().MarshalJSON(types3.NewQueryBondsParams(delegatorAccAddr, validatorAccAddr))
	if err != nil {
		fmt.Println(err)
	}

	route := fmt.Sprintf("custom/%s/%s", queryRoute, types3.QueryDelegation)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
	}

	var resp types3.DelegationResponse
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
	}

	ctx.PrintOutput(resp)
}

// GetCmdQueryDelegations implements the command to query all the delegations
// made from one delegator.
func GetCmdQueryDelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr string) {
	delegatorAccAddr, err := types.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	bz, err := ctx.GetCodec().MarshalJSON(types3.NewQueryDelegatorParams(delegatorAccAddr))
	if err != nil {
		fmt.Println(err)
	}

	route := fmt.Sprintf("custom/%s/%s", queryRoute, types3.QueryDelegatorDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
	}

	var resp types3.DelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
	}
	ctx.PrintOutput(resp)
}

// GetCmdQueryValidatorDelegations implements the command to query all the
// delegations to a specific validator.
// , cdc *codec.Codec
func GetCmdQueryValidatorDelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr string) {

	delegatorAccAddr, err := types.ValAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	bz, err := ctx.GetCodec().MarshalJSON(types3.NewQueryValidatorParams(delegatorAccAddr))
	if err != nil {
		fmt.Println(err)
	}

	route := fmt.Sprintf("custom/%s/%s", queryRoute, types3.QueryValidatorDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
	}

	var resp types3.DelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
	}
	ctx.PrintOutput(resp)
}

// GetCmdQueryUnbondingDelegation implements the command to query a single
// unbonding-delegation record.
func GetCmdQueryUnbondingDelegation(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr, validatorAddr string) {

	validatorAccAddr, err := types.ValAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	delAccAddr, err := types.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	bz, err := ctx.GetCodec().MarshalJSON(types3.NewQueryBondsParams(delAccAddr, validatorAccAddr))
	if err != nil {
		fmt.Println(err)
	}

	route := fmt.Sprintf("custom/%s/%s", queryRoute, types3.QueryUnbondingDelegation)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
	}

	var resp types3.UnbondingDelegation
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
	}
	ctx.PrintOutput(resp)
}

// GetCmdQueryUnbondingDelegations implements the command to query all the
// unbonding-delegation records for a delegator.
func GetCmdQueryUnbondingDelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr string) {

	delegatorAccAddr, err := types.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	bz, err := ctx.GetCodec().MarshalJSON(types3.NewQueryDelegatorParams(delegatorAccAddr))
	if err != nil {
		fmt.Println(err)
	}

	route := fmt.Sprintf("custom/%s/%s", queryRoute, types3.QueryDelegatorUnbondingDelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
	}

	var resp types3.UnbondingDelegations
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
	}
	ctx.PrintOutput(resp)
}

// GetCmdQueryRedelegation implements the command to query a single
// redelegation record.
// gatecli staking redelegation gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4
// gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4
// gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
func GetCmdQueryRedelegation(ctx *context.NodeVaultQuerierImpl, queryRoute string, args []string) {

	delAddr, err := types.AccAddressFromBech32(args[0])
	if err != nil {
		fmt.Println(err)
	}

	valSrcAddr, err := types.ValAddressFromBech32(args[1])
	if err != nil {
		fmt.Println(err)
	}

	valDstAddr, err := types.ValAddressFromBech32(args[2])
	if err != nil {
		fmt.Println(err)
	}

	bz, err := ctx.GetCodec().MarshalJSON(types3.NewQueryRedelegationParams(delAddr, valSrcAddr, valDstAddr))
	if err != nil {
		fmt.Println(err)
	}

	route := fmt.Sprintf("custom/%s/%s", queryRoute, types3.QueryRedelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
	}

	var resp types3.RedelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
	}
	ctx.PrintOutput(resp)
}

// GetCmdQueryRedelegations implements the command to query all the
// redelegation records for a delegator.
func GetCmdQueryRedelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr string) {
	delAddr, err := types.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	bz, err := ctx.GetCodec().MarshalJSON(types3.QueryRedelegationParams{DelegatorAddr: delAddr})
	if err != nil {
		fmt.Println(err)
	}

	route := fmt.Sprintf("custom/%s/%s", queryRoute, types3.QueryRedelegations)
	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
	}

	var resp types3.RedelegationResponses
	if err := ctx.GetCodec().UnmarshalJSON(res, &resp); err != nil {
		fmt.Println(err)
	}
	ctx.PrintOutput(resp)
}

// GetCmdQueryPool implements the pool query command.
func GetCmdQueryPool(ctx *context.NodeVaultQuerierImpl, storeName string) {

	bz, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/pool", storeName), nil)
	if err != nil {
		fmt.Println(err)
	}

	var pool types3.Pool
	if err := ctx.GetCodec().UnmarshalJSON(bz, &pool); err != nil {
		fmt.Println(err)
	}

	ctx.PrintOutput(pool)
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams(ctx *context.NodeVaultQuerierImpl, storeName string) {
	route := fmt.Sprintf("custom/%s/%s", storeName, types3.QueryParameters)
	bz, _, err := ctx.QueryWithData(route, nil)
	if err != nil {
		fmt.Println(err)
	}

	var params types3.Params
	ctx.GetCodec().MustUnmarshalJSON(bz, &params)
	ctx.PrintOutput(params)
}

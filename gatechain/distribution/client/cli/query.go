package cli

import (
	"fmt"
	"strconv"

	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/client/common"
	types2 "github.com/gatechain/gatechainsdk/gatechain/distribution/types"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// gatecli distribution params
// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams(ctx *context.NodeVaultQuerierImpl, queryRoute string) {
	params, err := common.QueryParams(*ctx, queryRoute)
	if err != nil {
		fmt.Println(err)
	}
	ctx.PrintOutput(params)
}

// GetCmdQueryValidatorOutstandingRewards implements the query validator outstanding rewards command.
func GetCmdQueryValidatorOutstandingRewards(ctx *context.NodeVaultQuerierImpl, queryRoute, ValidatorAddress string) {

	validatorValAddress, err := types.ValAddressFromBech32(ValidatorAddress)
	if err != nil {
		fmt.Println(err)
	}

	params := types2.NewQueryValidatorOutstandingRewardsParams(validatorValAddress)
	bz, err := ctx.GetCodec().MarshalJSON(params)
	if err != nil {
		fmt.Println(err)
	}

	resp, _, err := ctx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", queryRoute, types2.QueryValidatorOutstandingRewards),
		bz,
	)
	if err != nil {
		fmt.Println(err)
	}

	var outstandingRewards types2.ValidatorOutstandingRewards
	if err := ctx.GetCodec().UnmarshalJSON(resp, &outstandingRewards); err != nil {
		fmt.Println(err)
	}

	ctx.PrintOutput(outstandingRewards)
}

// GetCmdQueryValidatorCommission implements the query validator commission command.
func GetCmdQueryValidatorCommission(ctx *context.NodeVaultQuerierImpl, queryRoute, validatorAddr string) {
	validatorValAddr, err := types.ValAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	res, err := common.QueryValidatorCommission(*ctx, queryRoute, validatorValAddr)
	if err != nil {
		fmt.Println(err)
	}

	var valCom types2.ValidatorAccumulatedCommission
	ctx.GetCodec().MustUnmarshalJSON(res, &valCom)
	ctx.PrintOutput(valCom)
}

// GetCmdQueryValidatorSlashes implements the query validator slashes command.
func GetCmdQueryValidatorSlashes(ctx *context.NodeVaultQuerierImpl, queryRoute string, validatorAddress, startHeightStr, endHeightStr string) {
	validatorValAddr, err := types.ValAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
	}

	startHeight, err := strconv.ParseUint(startHeightStr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	endHeight, err := strconv.ParseUint(endHeightStr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	params := types2.NewQueryValidatorSlashesParams(validatorValAddr, startHeight, endHeight)
	bz, err := ctx.GetCodec().MarshalJSON(params)
	if err != nil {
		fmt.Println(err)
	}

	res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/validator_slashes", queryRoute), bz)
	if err != nil {
		fmt.Println(err)
	}

	var slashes types2.ValidatorSlashEvents
	ctx.GetCodec().MustUnmarshalJSON(res, &slashes)
	ctx.PrintOutput(slashes)
}

// GetCmdQueryDelegatorRewards implements the query delegator rewards command.
func GetCmdQueryDelegatorRewards(ctx *context.NodeVaultQuerierImpl, queryRoute string, args []string) {
	if len(args) == 2 {
		// query for rewards from a particular delegation
		resp, err := common.QueryDelegationRewards(*ctx, queryRoute, args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}

		var result types.DecCoins
		ctx.GetCodec().MustUnmarshalJSON(resp, &result)
		ctx.PrintOutput(result)
	}

	// query for delegator total rewards
	resp, err := common.QueryDelegatorTotalRewards(*ctx, queryRoute, args[0])
	if err != nil {
		fmt.Println(err)
	}

	var result types2.QueryDelegatorTotalRewardsResponse
	ctx.GetCodec().MustUnmarshalJSON(resp, &result)
	ctx.PrintOutput(result)
}

// GetCmdQueryCommunityPool returns the command for fetching community pool info
func GetCmdQueryCommunityPool(ctx *context.NodeVaultQuerierImpl, queryRoute string) {
	res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/community_pool", queryRoute), nil)
	if err != nil {
		fmt.Println(err)
	}

	var result types.DecCoins
	ctx.GetCodec().MustUnmarshalJSON(res, &result)
	ctx.PrintOutput(result)
}

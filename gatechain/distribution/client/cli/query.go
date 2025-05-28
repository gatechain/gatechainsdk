package cli

import (
	"fmt"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/client/common"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

func QueryParams(ctx *context.NodeVaultQuerierImpl) error {
	params, err := common.QueryParams(*ctx, types.ModuleName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return ctx.PrintOutput(params)
}

func QueryValidatorOutstandingRewards(ctx *context.NodeVaultQuerierImpl, ValidatorAddress string) error {

	validatorValAddress, err := sdk.ValAddressFromBech32(ValidatorAddress)
	if err != nil {
		fmt.Println(err)
		return err
	}

	params := types.NewQueryValidatorOutstandingRewardsParams(validatorValAddress)
	bz, err := ctx.GetCodec().MarshalJSON(params)
	if err != nil {
		fmt.Println(err)
		return err
	}

	resp, _, err := ctx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryValidatorOutstandingRewards),
		bz,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var outstandingRewards types.ValidatorOutstandingRewards
	if err := ctx.GetCodec().UnmarshalJSON(resp, &outstandingRewards); err != nil {
		fmt.Println(err)
		return err
	}
	return ctx.PrintOutput(outstandingRewards)
}

func QueryValidatorCommission(ctx *context.NodeVaultQuerierImpl, validatorAddr string) error {
	validatorValAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res, err := common.QueryValidatorCommission(*ctx, types.ModuleName, validatorValAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var valCom types.ValidatorAccumulatedCommission
	ctx.GetCodec().MustUnmarshalJSON(res, &valCom)
	return ctx.PrintOutput(valCom)
}

// QueryDelegatorRewards implements the query delegator rewards command.
func QueryDelegatorRewards(ctx *context.NodeVaultQuerierImpl, args []string) error {
	if len(args) == 2 {
		// query for rewards from a particular delegation
		resp, err := common.QueryDelegationRewards(*ctx, types.ModuleName, args[0], args[1])
		if err != nil {
			fmt.Println(err)
			return err
		}

		var result sdk.DecCoins
		ctx.GetCodec().MustUnmarshalJSON(resp, &result)

		return ctx.PrintOutput(result)
	}

	// query for delegator total rewards
	resp, err := common.QueryDelegatorTotalRewards(*ctx, types.ModuleName, args[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	var result types.QueryDelegatorTotalRewardsResponse
	ctx.GetCodec().MustUnmarshalJSON(resp, &result)
	return ctx.PrintOutput(result)
}

// QueryCommunityPool returns the command for fetching community pool info
func QueryCommunityPool(ctx *context.NodeVaultQuerierImpl) error {
	res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/community_pool", types.ModuleName), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var result sdk.DecCoins
	ctx.GetCodec().MustUnmarshalJSON(res, &result)
	return ctx.PrintOutput(result)
}

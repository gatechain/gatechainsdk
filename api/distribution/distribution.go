package distribution

import (
	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	cli2 "github.com/gatechain/gatechainsdk/gatechain/distribution/client/cli"
)

// gatecli distribution params
// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams(ctx *context.NodeVaultQuerierImpl, queryRoute string) {
	cli2.GetCmdQueryParams(ctx, queryRoute)
}

// GetCmdQueryValidatorOutstandingRewards implements the query validator outstanding rewards command.
func GetCmdQueryValidatorOutstandingRewards(ctx *context.NodeVaultQuerierImpl, queryRoute, ValidatorAddress string) {
	cli2.GetCmdQueryValidatorOutstandingRewards(ctx, queryRoute, ValidatorAddress)
}

// GetCmdQueryValidatorCommission implements the query validator commission command.
func GetCmdQueryValidatorCommission(ctx *context.NodeVaultQuerierImpl, queryRoute, validatorAddr string) {
	cli2.GetCmdQueryValidatorCommission(ctx, queryRoute, validatorAddr)
}

// GetCmdQueryValidatorSlashes implements the query validator slashes command.
func GetCmdQueryValidatorSlashes(ctx *context.NodeVaultQuerierImpl, queryRoute string, validatorAddress, startHeightStr, endHeightStr string) {
	cli2.GetCmdQueryValidatorSlashes(ctx, queryRoute, validatorAddress, startHeightStr, endHeightStr)
}

// GetCmdQueryDelegatorRewards implements the query delegator rewards command.
func GetCmdQueryDelegatorRewards(ctx *context.NodeVaultQuerierImpl, queryRoute string, args []string) {
	cli2.GetCmdQueryDelegatorRewards(ctx, queryRoute, args)
}

// GetCmdQueryCommunityPool returns the command for fetching community pool info
func GetCmdQueryCommunityPool(ctx *context.NodeVaultQuerierImpl, queryRoute string) {
	cli2.GetCmdQueryCommunityPool(ctx, queryRoute)
}

// command to withdraw rewards
func GetCmdWithdrawRewards(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, delegatorAddress, validatorAddress string, bComission bool) {
	cli2.GetCmdWithdrawRewards(ctx, txBldr, delegatorAddress, validatorAddress, bComission)
}

// command to withdraw all rewards
func GetGetCmdWithdrawAllRewards(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, queryRoute, delegatorAddr string) {
	cli2.GetGetCmdWithdrawAllRewards(ctx, txBldr, queryRoute, delegatorAddr)
}

// command to replace a delegator's withdrawal address
func SetWithdrawAddr(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, delegatorAddress, withdrawAddress string) {
	cli2.SetWithdrawAddr(ctx, txBldr, delegatorAddress, withdrawAddress)
}

// DelegatorAddress: delAddr,
//
//	validatorAddress: valAddr,
func GetCmdRewardReinvestment(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, delegatorAddress, validatorAddress string) {
	cli2.GetCmdRewardReinvestment(ctx, txBldr, delegatorAddress, validatorAddress)
}

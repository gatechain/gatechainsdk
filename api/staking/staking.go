package staking

import (
	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	cli2 "github.com/gatechain/gatechainsdk/gatechain/staking/client/cli"
)

// GetCmdQueryValidatorUnbondingDelegations implements the query all unbonding delegatations from a validator command.
func GetCmdQueryValidatorUnbondingDelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, ValidatorAddr string) {
	cli2.GetCmdQueryValidatorUnbondingDelegations(ctx, queryRoute, ValidatorAddr)
}

// GetCmdQueryValidatorRedelegations implements the query all redelegatations
// from a validator command.
func GetCmdQueryValidatorRedelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, SrcValidatorAddr string) {
	cli2.GetCmdQueryValidatorRedelegations(ctx, queryRoute, SrcValidatorAddr)
}

// GetCmdQueryDelegation the query delegation command.
func GetCmdQueryDelegation(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr, validatorAddr string) {
	cli2.GetCmdQueryDelegation(ctx, queryRoute, delegatorAddr, validatorAddr)
}

// GetCmdQueryDelegations implements the command to query all the delegations
// made from one delegator.
func GetCmdQueryDelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr string) {
	cli2.GetCmdQueryDelegations(ctx, queryRoute, delegatorAddr)
}

// GetCmdQueryValidatorDelegations implements the command to query all the
// delegations to a specific validator.
// , cdc *codec.Codec
func GetCmdQueryValidatorDelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr string) {
	cli2.GetCmdQueryValidatorDelegations(ctx, queryRoute, delegatorAddr)
}

// GetCmdQueryUnbondingDelegation implements the command to query a single
// unbonding-delegation record.
func GetCmdQueryUnbondingDelegation(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr, validatorAddr string) {
	cli2.GetCmdQueryUnbondingDelegation(ctx, queryRoute, delegatorAddr, validatorAddr)
}

// GetCmdQueryUnbondingDelegations implements the command to query all the
// unbonding-delegation records for a delegator.
func GetCmdQueryUnbondingDelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr string) {
	cli2.GetCmdQueryUnbondingDelegations(ctx, queryRoute, delegatorAddr)
}

// GetCmdQueryRedelegation implements the command to query a single
// redelegation record.
// gatecli staking redelegation gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4
// gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4
// gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
func GetCmdQueryRedelegation(ctx *context.NodeVaultQuerierImpl, queryRoute string, args []string) {
	cli2.GetCmdQueryRedelegation(ctx, queryRoute, args)
}

// GetCmdQueryRedelegations implements the command to query all the
// redelegation records for a delegator.
func GetCmdQueryRedelegations(ctx *context.NodeVaultQuerierImpl, queryRoute, delegatorAddr string) {
	cli2.GetCmdQueryRedelegations(ctx, queryRoute, delegatorAddr)
}

// GetCmdQueryPool implements the pool query command.
func GetCmdQueryPool(ctx *context.NodeVaultQuerierImpl, storeName string) {
	cli2.GetCmdQueryPool(ctx, storeName)
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams(ctx *context.NodeVaultQuerierImpl, storeName string) {
	cli2.GetCmdQueryParams(ctx, storeName)
}

// GetCmdDelegate implements the delegate command.
func GetCmdDelegate(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, amountStr, delegatorAddress, validatorAddr string) {
	cli2.GetCmdDelegate(ctx, txBldr, amountStr, delegatorAddress, validatorAddr)
}

// GetCmdRedelegate the begin redelegation command.
func GetCmdRedelegate(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, delAddr string, args []string) {
	cli2.GetCmdRedelegate(ctx, txBldr, delAddr, args)
}

// GetCmdUnbond implements the unbond validator command.
func GetCmdUnbond(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, amountStr, delegatorAddress, validatorAddr string) {
	cli2.GetCmdUnbond(ctx, txBldr, amountStr, delegatorAddress, validatorAddr)
}

// GetCmdUnbond implements the unbond validator command by SecurityAddress
func GetCmdUnbondBySecAddr(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, args []string) {
	cli2.GetCmdUnbondBySecAddr(ctx, txBldr, args)
}

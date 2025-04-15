package distribution

import (
	"fmt"
	"testing"

	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

//Distribution
//done
// gatecli distribution params
/*
gatecli distribution params
Distribution Params:
  Community Tax:          "0.000000000000000000"
  Withdraw Addr Enabled:  true
  First CommitteeReward:  "0.600000000000000000"
  Second CommitteeReward:  "0.400000000000000000"
  Third CommitteeReward:  "-0.060000000000000000"
*/
func TestGetCmdQueryParams(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	GetCmdQueryParams(ctx, common.Distribution)
}

// done
/*
	Use:   "outstanding-rewards [con-account]",
    Short: "Query distribution outstanding (un-withdrawn) rewards for a con-account and all their delegations",

gatecli distribution outstanding-rewards gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
1486785158870.699999609700000000NANOGT
*/
func TestGetCmdQueryValidatorOutstandingRewards(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	validatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9" //args[0]

	GetCmdQueryValidatorOutstandingRewards(ctx, common.Distribution, validatorAddress)
}

// todo not equal  test case 872282736722.700000000000000000NANOGT
// Use:   "commission [con-account]",
// Short: "Query distribution con-account commission",
//
/*
gatecli distribution commission gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
1439105158870.699999608040806648NANOGT
*/
func TestGetCmdQueryValidatorCommission(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	validatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9" //args[0]
	GetCmdQueryValidatorCommission(ctx, common.Distribution, validatorAddr)
}

// done
/*
	Use:   "slashes [validator] [start-height] [end-height]",
    Short: "Query distribution con-account slashes",

	 “Query the penalties of a consensus account”
	 gatecli distribution slashes gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9  0 4139617
*/
func TestGetCmdQueryValidatorSlashes(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	startHeightStr := "0"
	endHeightStr := "4139617"
	validatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	GetCmdQueryValidatorSlashes(ctx, common.Distribution, validatorAddress, startHeightStr, endHeightStr)
}

/*
	Use:   "rewards [delegator-addr] [<con-account_addr>]",
	Short: "query for rewards from a particular delegation",

gatecli staking delegations gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz
Delegation:

	Delegator:   gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz
	Con-account: gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
	Shares:      100.000000000000000000
	Balance:   100

Delegation:

	Delegator:   gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz
	Con-account: gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
	Shares:      1452000360.000000000000000000
	Balance:   1452000360

gatecli distribution rewards gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9 --chain-id gate-66
0.000000000000068800NANOGT
*/
func TestGetCmdQueryDelegatorRewards(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	args := []string{
		"gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz", //Delegator
		"gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9", // Con-account
	}
	GetCmdQueryDelegatorRewards(ctx, common.Distribution, args)
}

// done
/*
	Use:   "community-pool",
	Short: "Query the amount of coins in the community pool",

	gatecli distribution community-pool
	1117139129987.300003121906121500NANOGT
*/
func TestGetCmdQueryCommunityPool(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	GetCmdQueryCommunityPool(ctx, common.Distribution)
}

func TestGetCmdWithdrawRewards(t *testing.T) {
	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	validatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"

	runWithdrawTest := func(bCommission bool) {
		cdc := utils.MakeCodec()
		from := ""
		if bCommission {
			from = validatorAddress
		} else {
			from = delegatorAddress
		}
		ctx := context.NewCLIContextWithFrom(from, common.EndPoint, common.APIToken).
			WithCodec(cdc)

		txBldr := auth2.NewTxBuilderFromCLI().
			WithTxEncoder(sdk.DefaultTxEncoder(cdc)).
			WithFees("100000000NANOGT").
			WithGas(200000).
			WithChainID("gate-66")

		var err error
		txBldr, err = auth2.UpdateValidHeight(ctx, txBldr)
		if err != nil {
			t.Fatalf("failed to update valid height: %v", err)
		}
		fmt.Printf("TxBuilder: %+v\n", txBldr)
		GetCmdWithdrawRewards(ctx, txBldr, delegatorAddress, validatorAddress, bCommission)
	}

	t.Run("WithdrawWithCommission", func(t *testing.T) {
		runWithdrawTest(true)
	})

	t.Run("WithdrawWithoutCommission", func(t *testing.T) {
		runWithdrawTest(false)
	})
}

// GetCmdWithdrawAllRewards
func TestGetCmdWithdrawAllRewards(t *testing.T) {
	delegatorAddr := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(delegatorAddr, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	txBldr := auth2.NewTxBuilderFromCLI()
	txBldr = txBldr.WithTxEncoder(sdk.DefaultTxEncoder(cdc))
	txBldr = txBldr.WithFees("100000000NANOGT")
	txBldr = txBldr.WithGas(200000)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	txBldr = txBldr.WithChainID("gate-66")
	fmt.Println(txBldr, err)

	GetGetCmdWithdrawAllRewards(ctx, txBldr, common.Distribution, delegatorAddr)
}

/*
gatecli distribution set-withdraw-addr gt115rr8vfavgxa0kz5exqm7f8e6x8jg7f9gyrhzaa33hh478suwukagv26skgd98j3mza79gf --from validator1 --chain-id gate-66 --fees 1000000NANOGT --yes
*/
func TestCmdSetWithdrawAddr(t *testing.T) {
	delegatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"

	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(delegatorAddress, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	txBldr := auth2.NewTxBuilderFromCLI()
	txBldr = txBldr.WithTxEncoder(sdk.DefaultTxEncoder(cdc))
	txBldr = txBldr.WithFees("100000000NANOGT")
	txBldr = txBldr.WithGas(200000)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	txBldr = txBldr.WithChainID("gate-66")

	fmt.Println(txBldr, err)
	withdrawAddress := "gt115rr8vfavgxa0kz5exqm7f8e6x8jg7f9gyrhzaa33hh478suwukagv26skgd98j3mza79gf"
	SetWithdrawAddr(ctx, txBldr, delegatorAddress, withdrawAddress)
}

func TestGetCmdRewardReinvestment(t *testing.T) {
	cdc := utils.MakeCodec()
	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"

	ctx := context.NewCLIContextWithFrom(delegatorAddress, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	txBldr := auth2.NewTxBuilderFromCLI()
	txBldr = txBldr.WithTxEncoder(sdk.DefaultTxEncoder(cdc))
	txBldr = txBldr.WithFees("100000000NANOGT")

	txBldr = txBldr.WithGas(200000)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	txBldr = txBldr.WithChainID("gate-66")
	fmt.Println(txBldr, err)

	validatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"

	GetCmdRewardReinvestment(ctx, txBldr, delegatorAddress, validatorAddress)
}

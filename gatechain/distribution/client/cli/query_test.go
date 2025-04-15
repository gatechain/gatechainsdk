package cli

import (
	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"testing"
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

package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestGetCmdQueryDelegation(t *testing.T) {
	delegatorAddr := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	validatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryDelegation(delegatorAddr, validatorAddr)

}

func TestGetCmdQueryDelegations(t *testing.T) {
	delegatorAddr := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryDelegations(delegatorAddr)
}

func TestGetCmdQueryValidatorUnbondingDelegations(t *testing.T) {
	ValidatorAddr := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryValidatorUnbondingDelegations(ValidatorAddr)
}

func TestGetCmdQueryPool(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryPool()
}

func TestGetCmdQueryParams(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryParams()
}

func TestGetCmdQueryRedelegations(t *testing.T) {
	delegatorAddr := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryRedelegations(delegatorAddr)
}

func TestGetCmdQueryRedelegation(t *testing.T) {
	args := []string{
		"gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2", //delAddr
		"gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9", //valSrcAddr
		"gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh", //valDstAddr
	}
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryRedelegation(args)
}

func TestGetCmdQueryValidatorDelegations(t *testing.T) {
	delegatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryValidatorDelegations(delegatorAddr)
}

// "undelegation [delegator-addr] [con-account_addr]"
func TestGetCmdQueryUnbondingDelegation(t *testing.T) {
	delegatorAddr := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	validatorAddr := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryUnbondingDelegation(delegatorAddr, validatorAddr)
}

// undelegations [delegator-addr]
func TestGetCmdQueryUnbondingDelegations(t *testing.T) {
	delegatorAddr := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryUnbondingDelegations(delegatorAddr)
}

func TestGetCmdQueryValidatorRedelegations(t *testing.T) {
	SrcValidatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryValidatorRedelegations(SrcValidatorAddr)
}

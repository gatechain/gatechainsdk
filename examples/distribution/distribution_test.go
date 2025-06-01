package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestGetCmdQueryParams(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Distribution.QueryParams()
}

func TestGetCmdQueryValidatorOutstandingRewards(t *testing.T) {
	validatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9" //args[0]
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Distribution.QueryValidatorOutstandingRewards(validatorAddress)
}

func TestGetCmdQueryValidatorCommission(t *testing.T) {
	validatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9" //args[0]
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Distribution.QueryValidatorCommission(validatorAddr)
}

// query for rewards from a particular delegation
func TestGetCmdQueryDelegatorRewards(t *testing.T) {
	args := []string{
		"gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2", //Delegator
		"gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9", // Con-account
	}
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Distribution.QueryDelegatorRewards(args)
}

// query for delegator total rewards
func TestGetCmdQueryDelegatorTotalRewards(t *testing.T) {
	args := []string{
		"gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2", //Delegator
	}
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Distribution.QueryDelegatorRewards(args)
}

func TestGetCmdQueryCommunityPool(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Distribution.QueryCommunityPool()
}

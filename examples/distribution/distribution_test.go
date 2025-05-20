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

//func TestGetCmdQueryValidatorSlashes(t *testing.T) {
//	startHeightStr := "0"
//	endHeightStr := "4139617"
//	validatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
//
//	client := client.NewClient(common.EndPoint, common.APIToken)
//	client.Distribution.QueryValidatorSlashes(validatorAddress, startHeightStr, endHeightStr)
//}

// query for rewards from a particular delegation
func TestGetCmdQueryDelegatorRewards(t *testing.T) {
	args := []string{
		"gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz", //Delegator
		"gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9", // Con-account
	}
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Distribution.QueryDelegatorRewards(args)
}

// query for delegator total rewards
func TestGetCmdQueryDelegatorTotalRewards(t *testing.T) {
	args := []string{
		"gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz", //Delegator
	}
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Distribution.QueryDelegatorRewards(args)
}

func TestGetCmdQueryCommunityPool(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Distribution.QueryCommunityPool()
}

func TestGetCmdWithdrawRewardsCommission(t *testing.T) {
	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	validatorAddress := ""

	runWithdrawTest := func(bCommission bool) {
		from := ""
		if bCommission {
			from = validatorAddress
		} else {
			from = delegatorAddress
		}
		fees := "100000000NANOGT"
		gas := uint64(200000)
		chainID := "gate-66"

		client := client.NewClientWithFrom(from, common.EndPoint, common.APIToken, common.RootDir)
		client.Distribution.WithdrawRewards(delegatorAddress, validatorAddress, bCommission, fees, chainID, gas)
	}

	runWithdrawTest(true)
}

func TestGetCmdWithdrawRewards(t *testing.T) {
	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	validatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"

	runWithdrawTest := func(bCommission bool) {
		from := ""
		if bCommission {
			from = validatorAddress
		} else {
			from = delegatorAddress
		}
		fees := "100000000NANOGT"
		gas := uint64(200000)
		chainID := "gate-66"

		client := client.NewClientWithFrom(from, common.EndPoint, common.APIToken, common.RootDir)
		client.Distribution.WithdrawRewards(delegatorAddress, validatorAddress, bCommission, fees, chainID, gas)
	}

	runWithdrawTest(false)
}

// GetCmdWithdrawAllRewards
func TestGetCmdWithdrawAllRewards(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	delegatorAddr := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	client := client.NewClientWithFrom(delegatorAddr, common.EndPoint, common.APIToken, common.RootDir)
	client.Distribution.WithdrawAllRewards(delegatorAddr, fees, chainID, gas)
}

func TestCmdSetWithdrawAddr(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	delegatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	withdrawAddress := "gt115rr8vfavgxa0kz5exqm7f8e6x8jg7f9gyrhzaa33hh478suwukagv26skgd98j3mza79gf"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	client := client.NewClientWithFrom(delegatorAddress, common.EndPoint, common.APIToken, common.RootDir)
	client.Distribution.SetWithdrawAddr(delegatorAddress, withdrawAddress, fees, chainID, gas)
}

func TestGetCmdRewardReinvestment(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	validatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClientWithFrom(delegatorAddress, common.EndPoint, common.APIToken, common.RootDir)
	client.Distribution.RewardReinvestment(delegatorAddress, validatorAddress, fees, chainID, gas)
}

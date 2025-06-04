package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
	"github.com/gatechain/gatechainsdk/examples/testutil"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
)

var kb keys.Keybase

func init() {
	kb = testutil.ImportAllPrivateKeys()
}

// TestGetCmdQueryDelegatorTotalRewards
// bCommission is true  NewMsgWithdrawValidatorCommission. false MsgWithdrawDelegatorReward
// MsgWithdrawDelegatorReward{DelegatorAddress: delAddr, ValidatorAddress: valAddr, 	}
func TestGetWithdrawRewards(t *testing.T) {
	delegatorAddress := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
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

		client := client.NewClientWithFromMemKeyBase(from, common.EndPoint, common.APIToken, kb)
		client.Distribution.WithdrawRewards(delegatorAddress, validatorAddress, bCommission, fees, chainID, gas, kb)
	}

	runWithdrawTest(false)
}

// TestGetCmdQueryDelegatorTotalRewards
// bCommission is true  NewMsgWithdrawValidatorCommission. false MsgWithdrawDelegatorReward
func TestGetCmdWithdrawRewardsCommission(t *testing.T) {
	delegatorAddress := ""
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

		client := client.NewClientWithFromMemKeyBase(from, common.EndPoint, common.APIToken, kb)
		client.Distribution.WithdrawRewards(delegatorAddress, validatorAddress, bCommission, fees, chainID, gas, kb)
	}

	runWithdrawTest(true)
}

// use TestGetCmdQueryDelegatorTotalRewards to qury
func TestGetCmdWithdrawAllRewards(t *testing.T) {
	delegatorAddr := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	client := client.NewClientWithFromMemKeyBase(delegatorAddr, common.EndPoint, common.APIToken, kb)
	client.Distribution.WithdrawAllRewards(delegatorAddr, fees, chainID, gas, kb)
}

func TestCmdSetWithdrawAddr(t *testing.T) {
	delegatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	withdrawAddress := "gt115rr8vfavgxa0kz5exqm7f8e6x8jg7f9gyrhzaa33hh478suwukagv26skgd98j3mza79gf"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	client := client.NewClientWithFromMemKeyBase(delegatorAddress, common.EndPoint, common.APIToken, kb)
	client.Distribution.SetWithdrawAddr(delegatorAddress, withdrawAddress, fees, chainID, gas, kb)
}

func TestGetCmdRewardReinvestment(t *testing.T) {
	delegatorAddress := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	validatorAddress := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClientWithFromMemKeyBase(delegatorAddress, common.EndPoint, common.APIToken, kb)
	client.Distribution.RewardReinvestment(delegatorAddress, validatorAddress, fees, chainID, gas, kb)
}

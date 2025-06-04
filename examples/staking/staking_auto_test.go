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

func TestDelegate(t *testing.T) {
	amountStr := "363000140NANOGT"
	delegatorAddress := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	validatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9" //args[0]

	from_addr := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFromMemKeyBase(from_addr, common.EndPoint, common.APIToken, kb)
	//gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2
	client.Staking.GetDelegate(amountStr, delegatorAddress, validatorAddr, fees, chainID, gas, kb)
}

func TestGetCmdRedelegate(t *testing.T) {
	args := []string{
		"gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9", //valSrcAddr
		"gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh", //valDstAddr
		"100NANOGT",
	}

	delegatorAddress := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFromMemKeyBase(delegatorAddress, common.EndPoint, common.APIToken, kb)
	client.Staking.GetRedelegate(delegatorAddress, fees, chainID, gas, args, kb)
}

func TestGetCmdUnbond(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	validatorAddress := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh"
	amountStr := "100NANOGT"
	delegatorAddress := "gt11a6javtfmhhnhg67rj44p9tpfqpq2v9shd9smxvshuw0c54rase0pcw9c3pf8fm40hm0gr2"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFromMemKeyBase(delegatorAddress, common.EndPoint, common.APIToken, kb)
	client.Staking.GetUnbond(amountStr, delegatorAddress, validatorAddress, fees, chainID, gas, kb)
}

func TestGetCmdUnbondBySecAddr(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	securityAddress := "gt116lma26rvst9xhlg6fm2facm02mpgfhtcje3l0zul2ey06kfvfdf03q9u6vyd9fz7n65l7e"
	args := []string{
		"vault11k4rhqkllr5fke2hy2pyuj92vhk3jv8u7vzsvg4y8qkxppv8m29xnz2dm9kwp6hhgg5g47t",
	}
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	client := client.NewClientWithFromMemKeyBase(securityAddress, common.EndPoint, common.APIToken, kb)
	client.Staking.GetUnbondBySecAddr(fees, chainID, gas, args, kb)
}

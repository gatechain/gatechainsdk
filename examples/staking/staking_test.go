package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestGetCmdQueryDelegation(t *testing.T) {
	delegatorAddr := "gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04"
	validatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryDelegation(delegatorAddr, validatorAddr)

}

func TestGetCmdQueryDelegations(t *testing.T) {
	delegatorAddr := "gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04"
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
	delegatorAddr := "gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryRedelegations(delegatorAddr)
}

func TestGetCmdQueryRedelegation(t *testing.T) {
	args := []string{
		"gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04",
		"gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh",
		"gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9",
	}
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryRedelegation(args)
}

func TestGetCmdQueryValidatorDelegations(t *testing.T) {
	delegatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryValidatorDelegations(delegatorAddr)
}

func TestGetCmdQueryUnbondingDelegation(t *testing.T) {
	delegatorAddr := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	validatorAddr := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryUnbondingDelegation(delegatorAddr, validatorAddr)
}

func TestGetCmdQueryUnbondingDelegations(t *testing.T) {
	delegatorAddr := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryUnbondingDelegations(delegatorAddr)
}

func TestGetCmdQueryValidatorRedelegations(t *testing.T) {
	SrcValidatorAddr := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Staking.QueryValidatorRedelegations(SrcValidatorAddr)
}

func TestGetCmdDelegate(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	amountStr := "363000140NANOGT"
	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	validatorAddr := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh" //args[0]

	from_addr := "vault11xcq2vcd2e4vastn9q9vd7kh7l8ksfrpghdkefzvte9a98h9rl8cmasr47gtzdr7ljdg3e6"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFrom(from_addr, common.EndPoint, common.APIToken, common.RootDir)
	client.Staking.GetDelegate(amountStr, delegatorAddress, validatorAddr, fees, chainID, gas)
}

func TestGetCmdRedelegate(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	args := []string{
		"gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh",
		"gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9",
		"100NANOGT",
	}

	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFrom(delegatorAddress, common.EndPoint, common.APIToken, common.RootDir)
	client.Staking.GetRedelegate(delegatorAddress, fees, chainID, gas, args)
}

func TestGetCmdUnbond(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	validatorAddress := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh"
	amountStr := "100NANOGT"
	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFrom(delegatorAddress, common.EndPoint, common.APIToken, common.RootDir)
	client.Staking.GetUnbond(amountStr, delegatorAddress, validatorAddress, fees, chainID, gas)
}

func TestGetCmdUnbondBySecAddr(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	securityAddress := "gt11vk8xsf2zfr7yzuadpa3t08e2qju4ps4x8pp6a4rkakgw3kup0gm4jy50s5uuyf45uf6ykz"
	args := []string{
		"vault11mzum4y48hn8xsw3hcdjwx7mu5zxwhnds6m3fqz4wrfteg0z9rkyt9yczrrjhjalq49935w",
	}
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	client := client.NewClientWithFrom(securityAddress, common.EndPoint, common.APIToken, common.RootDir)
	client.Staking.GetUnbondBySecAddr(fees, chainID, gas, args)
}

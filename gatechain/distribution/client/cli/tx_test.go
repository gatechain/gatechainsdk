package cli

import (
	"fmt"
	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
	"testing"
)

////withdraw-rewards
//
///*
//Use:   "withdraw-rewards [con-account]",
//Short: "Withdraw rewards from a given delegation address, and optionally withdraw con-account commission if the delegation address given is a con-account operator",
//
//gatecli distribution withdraw-rewards gt11vk8xsf2zfr7yzuadpa3t08e2qju4ps4x8pp6a4rkakgw3kup0gm4jy50s5uuyf45uf6ykz
//--from gt11mzum4y48hn8xsw3hcdjwx7mu5zxwhnds6m3fqz4wrfteg0z9rkyt9yczrrjhjalq99na4f --fees 10000000NANOGT
//--chain-id meteora --node http://139.162.15.16:8080 --home /mnt/meteora-data/ -y
//*/
//func TestGetCmdWithdrawRewards(t *testing.T) {
//	GetCmdWithdrawRewards()
//}

/*
Use:   "withdraw-rewards [con-account]",
Short: "Withdraw rewards from a given delegation address, and optionally withdraw con-account commission if the delegation address given is a con-account operator",

gatecli distribution withdraw-rewards gt11vk8xsf2zfr7yzuadpa3t08e2qju4ps4x8pp6a4rkakgw3kup0gm4jy50s5uuyf45uf6ykz
--from gt11mzum4y48hn8xsw3hcdjwx7mu5zxwhnds6m3fqz4wrfteg0z9rkyt9yczrrjhjalq99na4f --fees 10000000NANOGT
--chain-id meteora --node http://139.162.15.16:8080 --home /mnt/meteora-data/ -y
*/
func TestGetCmdWithdrawRewards(t *testing.T) {

	{
		delegatorAddress := "gt11vk8xsf2zfr7yzuadpa3t08e2qju4ps4x8pp6a4rkakgw3kup0gm4jy50s5uuyf45uf6ykz"
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
		bComission := true
		validatorAddress := "gt11mzum4y48hn8xsw3hcdjwx7mu5zxwhnds6m3fqz4wrfteg0z9rkyt9yczrrjhjalq99na4f"
		GetCmdWithdrawRewards(ctx, txBldr, delegatorAddress, validatorAddress, bComission)
	}

	{
		delegatorAddress := "gt11vk8xsf2zfr7yzuadpa3t08e2qju4ps4x8pp6a4rkakgw3kup0gm4jy50s5uuyf45uf6ykz"
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
		bComission := false
		validatorAddress := "gt11mzum4y48hn8xsw3hcdjwx7mu5zxwhnds6m3fqz4wrfteg0z9rkyt9yczrrjhjalq99na4f"
		GetCmdWithdrawRewards(ctx, txBldr, delegatorAddress, validatorAddress, bComission)
	}

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

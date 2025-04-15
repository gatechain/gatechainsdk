package vaultAccount

import (
	"fmt"
	"testing"

	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// done 2
// gatecli vault-account create [to_address] [security_address] [delay_height] [clearing_height] [amount] ([pubkey]) [flags]
func TestBroadcastMsgCreateVault(t *testing.T) {
	//pay fee
	from_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(from_addr, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	to_addr := "gt116gptv427t2tu447arzqkz2uhmwk8rmzxc09cffcmg9fsh3ntt7cszy4wjevhhvccn37m52"
	security_addr := "gt11469zdcp9ehj4xvv5q2kvs4uhs62uvuul73a364h5dwdtk8mwxfud03vvrzepgm23tccdm6"
	delayHeightStr := "10"
	clearTimeHeightStr := "20000"
	coinsStr := "100000000000NANOGT"
	pubkeyStr := ""
	fees := "100000000NANOGT"
	chainID := "gate-66"
	gas := uint64(200000)

	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}
	BroadcastMsgCreateVault(ctx, txBldr, from_addr, to_addr, security_addr, delayHeightStr, clearTimeHeightStr, coinsStr, pubkeyStr)
}

//done 2
// gatecli vault-account update-clearing-height [clearing_height] [flags]
/*
 gatecli vault-account update-clearing-height 20020  --from gt112wldf46teljw9uarndfjqcs9dynrrpkwtmltcwg9yhpun8weh2829783n4ut5hzdvwelqv
--fees 100000000NANOGT --chain-id gate-66 --yes
*/
func TestBroadcastUpdateClearingHeightTx(t *testing.T) {
	from_addr := "gt116gptv427t2tu447arzqkz2uhmwk8rmzxc09cffcmg9fsh3ntt7cszy4wjevhhvccn37m52"

	clearTimeHeight := "20020"
	vaultAddr := "vault116gptv427t2tu447arzqkz2uhmwk8rmzxc09cffcmg9fsh3ntt7cszy4wjevhhvccr3gh4d"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(from_addr, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}

	BroadcastUpdateClearingHeightTx(ctx, txBldr, clearTimeHeight, vaultAddr)
}

// done 2
// gatecli vault-account show [address] [flags]
func TestQueryVaultAccount(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	vaultAddr := "gt116gptv427t2tu447arzqkz2uhmwk8rmzxc09cffcmg9fsh3ntt7cszy4wjevhhvccn37m52"
	QueryVaultAccount(ctx, vaultAddr)
}

func TestClearVaultAccountTx(t *testing.T) {

	vaultAddresses := []string{
		"vault11zt4y23l79rwksl6w6mgrta8sarqwqe9kwyauc6g4gd8at6wuzzhsyqvkrz3fle79f42tmp",
	}

	from_addr := "gt1196669klaryfct82mtylw7hqzgyx59e08f7f8nen0p56569tkma250s22jhmyfdyhmvda0f"

	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(from_addr, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	fromAccAddress, err := types.AccAddressFromBech32(from_addr)
	if err != nil {
		fmt.Println(err)
	}
	ctx = ctx.WithFromAddress(fromAccAddress)

	fees := "100000000NANOGT"
	chainID := "gate-66"
	gas := uint64(200000)
	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr, err = auth2.UpdateValidHeight(ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}
	ClearVaultAccountTx(ctx, txBldr, from_addr, vaultAddresses)
}

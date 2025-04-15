package vaultAccount

import (
	"fmt"
	"strconv"

	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	client2 "github.com/gatechain/gatechainsdk/gatechain/revocable/client"
	types2 "github.com/gatechain/gatechainsdk/gatechain/revocable/types"
	types3 "github.com/gatechain/gatechainsdk/gatechain/types"
)

func BroadcastMsgCreateVault(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, from_addr, to_addr, security_addr, delayHeightStr, clearTimeHeightStr, coinsStr, pubkeyStr string) {

	fromAccAddress, _, err := types3.AccAddressTypeFromBech32(from_addr)
	if err != nil {
		fmt.Println(err)
	}
	toAccAddress, _, err := types3.AccAddressTypeFromBech32(to_addr)
	if err != nil {
		fmt.Println(err)
	}

	securityAccAddress, _, err := types3.AccAddressTypeFromBech32(security_addr)
	if err != nil {
		fmt.Println(err)
	}

	delayHeight, err := strconv.ParseUint(delayHeightStr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	clearTime, err := strconv.ParseUint(clearTimeHeightStr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	// parse coins trying to be sent
	coins, err := types3.ParseCoins(coinsStr)
	if err != nil {
		fmt.Println(err)
	}

	msg := types2.NewMsgCreateVault(fromAccAddress, toAccAddress, securityAccAddress, delayHeight, clearTime, coins, pubkeyStr)
	fmt.Println(msg)

	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
	fmt.Println(txbytes, err)
}

func BroadcastUpdateClearingHeightTx(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, clearTimeHeight, vaultAddr string) {
	clearTime, err := strconv.ParseUint(clearTimeHeight, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	vaultAccAddress, _, err := types3.AccAddressTypeFromBech32(vaultAddr)
	if err != nil {
		fmt.Println(err)
	}

	// build and sign the transaction, then broadcast to Tendermint
	msg := types2.NewMsgUpdateClearingHeight(vaultAccAddress, clearTime)

	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
	fmt.Println(txbytes, err)
}

func QueryVaultAccount(ctx *context.NodeVaultQuerierImpl, vaultAddr string) {
	retriever := auth2.NewVaultRetriever(ctx)
	key, err := types3.AccAddressFromBech32(vaultAddr)
	if err != nil {
		fmt.Println(err)
	}
	account, height, err := retriever.GetAccountWithHeight(key)
	fmt.Println(account, height, err)
}

func ClearVaultAccountTx(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, from_addr string, vaultAddresses []string) {
	var vaultAddress []types3.AccAddress
	for i := 0; i < len(vaultAddresses); i++ {
		address, _, err := types3.AccAddressTypeFromBech32(vaultAddresses[i])
		if err != nil {
			fmt.Println(err)
		}
		vaultAddress = append(vaultAddress, address)
	}
	if err := client2.EnsureFromVaultAccount(*ctx); err != nil {
		fmt.Println(err)
	}

	msg := types2.NewMsgClearVaultAccount(ctx.FromAddress, vaultAddress)
	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
	fmt.Println(txbytes, err)
}

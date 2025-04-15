package tx

import (
	"fmt"

	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	cli2 "github.com/gatechain/gatechainsdk/gatechain/auth/client/cli"
	"github.com/gatechain/gatechainsdk/gatechain/bank"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
)

// create unsign tx to 	unsignTxFileName
func CreateUnsignTX(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, from_addr, to_addr, amount, unsignTxFileName string) {

	toAccAddress, _, err := types2.AccAddressTypeFromBech32(to_addr)
	if err != nil {
		fmt.Println(err)
	}
	fromAccAddress, _, err := types2.AccAddressTypeFromBech32(from_addr)
	if err != nil {
		fmt.Println(err)
	}
	coins, err := types2.ParseCoins(amount)
	if err != nil {
		fmt.Println(err)
	}
	// build and sign the transaction, then broadcast to Tendermint
	msg := bank.NewMsgSend(fromAccAddress, toAccAddress, coins)
	fmt.Println(msg)
	auth2.PrintUnsignedStdTx(txBldr, ctx, []types2.Msg{msg}, unsignTxFileName)
}

// sign UnsignTxFile to SignTxFile
func CreateSignTX(UnsignTxFile, SignTxFile string, cliCtx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder) {
	cli2.MakeSignCmd(UnsignTxFile, SignTxFile, cliCtx, txBldr)
}

// BroadcastSignTx SignedFileName
func BroadcastSignTx(cliCtx *context.NodeVaultQuerierImpl, SignedFileName string) {
	cli2.GetBroadcastCommand(cliCtx, SignedFileName)
}

// create  and send tx
func SendTX(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, from_addr, to_addr, amount string) {

	toAccAddress, _, err := types2.AccAddressTypeFromBech32(to_addr)
	if err != nil {
		fmt.Println(err)
	}
	fromAccAddress, _, err := types2.AccAddressTypeFromBech32(from_addr)
	if err != nil {
		fmt.Println(err)
	}
	coins, err := types2.ParseCoins(amount)
	if err != nil {
		fmt.Println(err)
	}
	// build and sign the transaction, then broadcast to Tendermint
	msg := bank.NewMsgSend(fromAccAddress, toAccAddress, coins)
	fmt.Println(msg)
	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types2.Msg{msg}, true)
	fmt.Println(txbytes, err)
}

// query transaction hashHexStr
func QueryTX(ctx *context.NodeVaultQuerierImpl, hashHexStr string) {
	res, err := auth2.QueryTx(ctx, hashHexStr)
	if err != nil {
		fmt.Println(err)
	}
	if res.Empty() {
		fmt.Println(fmt.Errorf("No transaction found with hash %s", hashHexStr))
	}
	fmt.Println(res)
}

package cli

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// BroadcastSignTx returns the tx broadcast command.
func BroadcastSignTx(cliCtx *context.NodeVaultQuerierImpl, filename string) {
	stdTx, err := auth.ReadStdTxFromFile(cliCtx.Codec, filename)
	if err != nil {
		return
	}

	txBytes, err := cliCtx.Codec.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return
	}

	res, err := cliCtx.Client.BroadcastTx(txBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	resPos := types.NewResponseFormatBroadcastTx(&res)
	TxResponse := auth.ReConvertTxResponseFromData(cliCtx.Codec, resPos)
	TxResponse = auth.ConvertTxHashPrefixForTxResponse(TxResponse)
	cliCtx.PrintOutput(TxResponse)
}

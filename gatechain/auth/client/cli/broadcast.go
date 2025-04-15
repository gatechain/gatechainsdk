package cli

import (
	"fmt"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// GetBroadcastCommand returns the tx broadcast command.
func GetBroadcastCommand(cliCtx *context.NodeVaultQuerierImpl, filename string) {
	stdTx, err := auth2.ReadStdTxFromFile(cliCtx.Codec, filename)
	if err != nil {
		return
	}

	txBytes, err := cliCtx.Codec.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return
	}

	res, err := cliCtx.Client.BroadcastTx(txBytes)
	fmt.Println(res)
	resPos := types.NewResponseFormatBroadcastTx(&res)
	TxResponse := auth2.ReConvertTxResponseFromData(cliCtx.Codec, resPos)
	TxResponse = auth2.ConvertTxHashPrefixForTxResponse(TxResponse)
	cliCtx.PrintOutput(TxResponse)
}

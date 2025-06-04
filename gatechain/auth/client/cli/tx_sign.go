package cli

import (
	"fmt"
	"os"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

func CreateSignTX(UnsignTxFile, SignTxFile string, cliCtx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder) error {
	stdTx, err := auth.ReadStdTxFromFile(cliCtx.GetCodec(), UnsignTxFile)
	if err != nil {
		return err
	}
	txBldr = txBldr.WithValidHeight(stdTx.ValidHeight).WithNonce(stdTx.Nonces[0])
	var newTx auth.StdTx
	generateSignatureOnly := false

	appendSig := true
	newTx, err = auth.SignStdTx(txBldr, cliCtx, cliCtx.GetFromName(), stdTx, appendSig, true)

	if err != nil {
		return err
	}

	json, err := getSignatureJSON(cliCtx.GetCodec(), newTx, cliCtx.Indent, generateSignatureOnly)
	if err != nil {
		return err
	}

	fp, err := os.OpenFile(
		SignTxFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644,
	)
	if err != nil {
		return err
	}

	defer fp.Close()
	fmt.Fprintf(fp, "%s\n", json)

	return nil
}

func getSignatureJSON(cdc *codec.Codec, newTx auth.StdTx, indent, generateSignatureOnly bool) ([]byte, error) {
	switch generateSignatureOnly {
	case true:
		switch indent {
		case true:
			return cdc.MarshalJSONIndent(newTx.Signatures[0], "", "  ")

		default:
			return cdc.MarshalJSON(newTx.Signatures[0])
		}
	default:
		switch indent {
		case true:
			return cdc.MarshalJSONIndent(newTx, "", "  ")

		default:
			return cdc.MarshalJSON(newTx)
		}
	}
}

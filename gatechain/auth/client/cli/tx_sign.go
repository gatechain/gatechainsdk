package cli

import (
	"fmt"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
	"os"
)

func MakeSignCmd(UnsignTxFile, SignTxFile string, cliCtx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder) error {
	stdTx, err := auth2.ReadStdTxFromFile(cliCtx.GetCodec(), UnsignTxFile)
	if err != nil {
		return err
	}
	txBldr = txBldr.WithValidHeight(stdTx.ValidHeight).WithNonce(stdTx.Nonces[0])
	// if --signature-only is on, then override --append
	var newTx auth2.StdTx
	generateSignatureOnly := false
	multisigAddrStr := ""

	if multisigAddrStr != "" {
		var multisigAddr sdk.AccAddress

		multisigAddr, err = sdk.AccAddressFromBech32(multisigAddrStr)
		if err != nil {
			return err
		}

		newTx, err = auth2.SignStdTxWithSignerAddress(
			txBldr, cliCtx, multisigAddr, cliCtx.GetFromName(), stdTx, true,
		)
		generateSignatureOnly = true
	} else {
		appendSig := true
		newTx, err = auth2.SignStdTx(txBldr, cliCtx, cliCtx.GetFromName(), stdTx, appendSig, true)
	}

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

func getSignatureJSON(cdc *codec.Codec, newTx auth2.StdTx, indent, generateSignatureOnly bool) ([]byte, error) {
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

package auth

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// txBldr types.StdTxBuilder,
func CompleteAndBroadcastTxCLI(txBldr TxBuilder, cliCtx *context.NodeVaultQuerierImpl, msgs []types.Msg, isBroadcast bool) ([]byte, error) {

	txBldr, err := PrepareTxBuilder(txBldr, cliCtx)
	if err != nil {
		return nil, err
	}
	fromName := cliCtx.GetFromName()

	passphrase := DefaultKeyPass

	txBytes, err := txBldr.BuildAndSign(fromName, passphrase, msgs)
	if err != nil {
		return nil, err
	}

	if !isBroadcast {
		return txBytes, nil
	}

	// broadcast to a Tendermint node
	res, err := cliCtx.Client.BroadcastTx(txBytes)
	if err != nil {
		return txBytes, err
	}
	fmt.Println(res)
	resPos := types.NewResponseFormatBroadcastTx(&res)
	TxResponse := ReConvertTxResponseFromData(cliCtx.Codec, resPos)
	// format prefix txhash
	TxResponse = ConvertTxHashPrefixForTxResponse(TxResponse)

	return txBytes, cliCtx.PrintOutput(TxResponse)
}

// PrepareTxBuilder populates a TxBuilder in preparation for the build of a Tx.
func PrepareTxBuilder(txBldr TxBuilder, cliCtx *context.NodeVaultQuerierImpl) (TxBuilder, error) {
	from := cliCtx.GetFromAddress()

	accGetter := NewAccountRetriever(cliCtx, cliCtx.Codec)
	if err := accGetter.EnsureExists(from); err != nil {
		return txBldr, err
	}

	txbldrAccNum := txBldr.AccountNumber()
	// TODO: (ref #1903) Allow for user supplied account number without
	// automatically doing a manual lookup.
	if txbldrAccNum == 0 {
		num, err := NewAccountRetriever(cliCtx, cliCtx.Codec).GetAccountNumber(from)
		if err != nil {
			return txBldr, err
		}

		if txbldrAccNum == 0 {
			txBldr = txBldr.WithAccountNumber(num)
		}
	}

	return txBldr, nil
}

func UpdateValidHeight(cliCtx *context.NodeVaultQuerierImpl, txBldr TxBuilder) (TxBuilder, error) {
	currentHeight, err := cliCtx.GetChainHeight()
	if err != nil {
		return txBldr, err
	}

	txBldr = txBldr.WithValidHeight([]uint64{uint64(currentHeight) - 10, uint64(currentHeight) + 200})

	return txBldr, nil
}

// PrintUnsignedStdTx builds an unsigned StdTx and prints it to os.Stdout.
func PrintUnsignedStdTx(txBldr TxBuilder, cliCtx *context.NodeVaultQuerierImpl, msgs []types.Msg, filename string) error {
	stdTx, err := buildUnsignedStdTxOffline(txBldr, cliCtx, msgs)
	if err != nil {
		return err
	}

	json, err := cliCtx.Codec.MarshalJSON(stdTx)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(json)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	fmt.Printf("Unsigned transaction JSON written to %s\n", filename)
	return nil
}
func buildUnsignedStdTxOffline(txBldr TxBuilder, cliCtx *context.NodeVaultQuerierImpl, msgs []types.Msg) (stdTx StdTx, err error) {
	if txBldr.SimulateAndExecute() {
		if cliCtx.GenerateOnly {
			return stdTx, errors.New("cannot estimate gas with generate-only")
		}

		txBldr, err = EnrichWithGas(txBldr, cliCtx, msgs)
		if err != nil {
			return stdTx, err
		}

		_, _ = fmt.Fprintf(os.Stderr, "estimated gas = %v\n", txBldr.Gas())
	}

	stdSignMsg, err := txBldr.BuildSignMsg(msgs)
	if err != nil {
		return stdTx, nil
	}

	return NewStdTx(stdSignMsg.Msgs, stdSignMsg.Fee, nil, stdSignMsg.Memo, [][]byte{txBldr.Nonce()}, txBldr.ValidHeight()), nil
}

// EnrichWithGas calculates the gas estimate that would be consumed by the
// transaction and set the transaction's respective value accordingly.
func EnrichWithGas(txBldr TxBuilder, cliCtx *context.NodeVaultQuerierImpl, msgs []types.Msg) (TxBuilder, error) {
	_, adjusted, err := simulateMsgs(txBldr, cliCtx, msgs)
	if err != nil {
		return txBldr, err
	}

	return txBldr.WithGas(adjusted), nil
}

// nolint
// SimulateMsgs simulates the transaction and returns the gas estimate and the adjusted value.
func simulateMsgs(txBldr TxBuilder, cliCtx *context.NodeVaultQuerierImpl, msgs []types.Msg) (estimated, adjusted uint64, err error) {
	txBytes, err := txBldr.BuildTxForSim(msgs)
	if err != nil {
		return
	}

	estimated, adjusted, err = CalculateGas(cliCtx.QueryWithData, cliCtx.Codec, txBytes, txBldr.GasAdjustment())
	return
}

// CalculateGas simulates the execution of a transaction and returns
// both the estimate obtained by the query and the adjusted amount.
func CalculateGas(
	queryFunc func(string, []byte) ([]byte, int64, error), cdc *codec.Codec,
	txBytes []byte, adjustment float64,
) (estimate, adjusted uint64, err error) {

	// run a simulation (via /app/simulate query) to
	// estimate gas and update TxBuilder accordingly
	rawRes, _, err := queryFunc("/app/simulate", txBytes)
	if err != nil {
		return estimate, adjusted, err
	}

	estimate, err = parseQueryResponse(cdc, rawRes)
	if err != nil {
		return
	}

	adjusted = adjustGasEstimate(estimate, adjustment)
	return estimate, adjusted, nil
}

func parseQueryResponse(cdc *codec.Codec, rawRes []byte) (uint64, error) {
	var simulationResult types.Result
	if err := cdc.UnmarshalBinaryLengthPrefixed(rawRes, &simulationResult); err != nil {
		return 0, err
	}

	return simulationResult.GasUsed, nil
}

func adjustGasEstimate(estimate uint64, adjustment float64) uint64 {
	return uint64(adjustment * float64(estimate))
}

// Read and decode a StdTx from the given filename.  Can pass "-" to read from stdin.
func ReadStdTxFromFile(cdc *codec.Codec, filename string) (stdTx StdTx, err error) {
	var bytes []byte

	if filename == "-" {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(filename)
	}

	if err != nil {
		return
	}

	if err = cdc.UnmarshalJSON(bytes, &stdTx); err != nil {
		return
	}

	return
}

func isTxSigner(user types.AccAddress, signers []types.AccAddress) bool {
	for _, s := range signers {
		if bytes.Equal(user.Bytes(), s.Bytes()) {
			return true
		}
	}

	return false
}

func populateAccountFromState(
	txBldr TxBuilder, cliCtx *context.NodeVaultQuerierImpl, addr types.AccAddress,
) (TxBuilder, error) {

	num, err := NewAccountRetriever(cliCtx, cliCtx.Codec).GetAccountNumber(addr)
	if err != nil {
		return txBldr, err
	}

	return txBldr.WithAccountNumber(num), nil
}

// SignStdTx appends a signature to a StdTx and returns a copy of it. If appendSig
// is false, it replaces the signatures already attached with the new signature.
// Don't perform online validation or lookups if offline is true.
func SignStdTx(
	txBldr TxBuilder, cliCtx *context.NodeVaultQuerierImpl, name string,
	stdTx StdTx, appendSig bool, offline bool,
) (StdTx, error) {

	var signedStdTx StdTx

	info, err := txBldr.Keybase().Get(name)
	if err != nil {
		return signedStdTx, err
	}

	addr := info.GetPubKey().Address()

	// check whether the address is a signer
	if !isTxSigner(types.AccAddress(addr), stdTx.GetSigners()) {
		return signedStdTx, fmt.Errorf("%s: %s", errInvalidSigner, name)
	}

	if !offline {
		txBldr, err = populateAccountFromState(txBldr, cliCtx, types.AccAddress(addr))
		if err != nil {
			return signedStdTx, err
		}
	}
	return txBldr.SignStdTx(name, DefaultKeyPass, stdTx, appendSig)
}

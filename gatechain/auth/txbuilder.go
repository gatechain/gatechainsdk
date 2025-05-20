package auth

import (
	"errors"
	"fmt"
	gmhash "github.com/gatechain/crypto"

	"github.com/gatechain/gatechainsdk/common"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/crypto"
	crkeys "github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// TxBuilder implements a transaction context created in SDK modules.
type TxBuilder struct {
	txEncoder          types.TxEncoder
	keybase            crkeys.Keybase
	accountNumber      uint64
	nonce              []byte
	gas                uint64
	gasAdjustment      float64
	simulateAndExecute bool
	chainID            string
	memo               string
	fees               types.Coins
	gasPrices          types.DecCoins
	validHeight        []uint64
}

// NewTxBuilderFromCLI returns a new initialized TxBuilder with parameters from
// the command line using Viper.
func NewTxBuilderFromCLI(rootDir, fees, chainID string, gas uint64, cdc *codec.Codec) TxBuilder {
	kb, err := crypto.NewKeyBaseFromDir(rootDir)
	if err != nil {
		panic(err)
	}

	uuid, _ := types.NewV4()
	nonce := gmhash.Hash(uuid[:])
	txbldr := TxBuilder{
		keybase: kb,
		nonce:   nonce[:],
		chainID: chainID,
	}
	txbldr = txbldr.WithTxEncoder(DefaultTxEncoder(cdc))
	txbldr = txbldr.WithFees(fees) //fees "100000000NANOGT"
	txbldr = txbldr.WithGas(gas)   //gas  200000
	return txbldr
}

// AccountNumber returns the account number
func (bldr TxBuilder) AccountNumber() uint64 { return bldr.accountNumber }

// Nonce returns the transaction nonce
func (bldr TxBuilder) Nonce() []byte { return bldr.nonce }

// ExpireHeight returns the transaction expireHeight
func (bldr TxBuilder) ValidHeight() []uint64 { return bldr.validHeight }

// Gas returns the gas for the transaction
func (bldr TxBuilder) Gas() uint64 { return bldr.gas }

// GasAdjustment returns the gas adjustment
func (bldr TxBuilder) GasAdjustment() float64 { return bldr.gasAdjustment }

// Keybase returns the keybase
func (bldr TxBuilder) Keybase() crkeys.Keybase { return bldr.keybase }

// SimulateAndExecute returns the option to simulate and then execute the transaction
// using the gas from the simulation results
func (bldr TxBuilder) SimulateAndExecute() bool { return bldr.simulateAndExecute }

// WithTxEncoder returns a copy of the context with an updated codec.
func (bldr TxBuilder) WithTxEncoder(txEncoder types.TxEncoder) TxBuilder {
	bldr.txEncoder = txEncoder
	return bldr
}

// WithGas returns a copy of the context with an updated gas.
func (bldr TxBuilder) WithGas(gas uint64) TxBuilder {
	bldr.gas = gas
	return bldr
}

// WithFees returns a copy of the context with an updated fee.
func (bldr TxBuilder) WithFees(fees string) TxBuilder {
	parsedFees, err := types.ParseCoins(fees)
	if err != nil {
		panic(err)
	}

	bldr.fees = parsedFees
	return bldr
}

// WithNonce returns a copy of the context with an updated nonce.
func (bldr TxBuilder) WithNonce(nonce []byte) TxBuilder {
	bldr.nonce = nonce
	return bldr
}

// WithExpireHeight returns a copy of the context with an updated nonce.
func (bldr TxBuilder) WithValidHeight(validHeight []uint64) TxBuilder {
	bldr.validHeight = validHeight
	return bldr
}

// WithAccountNumber returns a copy of the context with an account number.
func (bldr TxBuilder) WithAccountNumber(accnum uint64) TxBuilder {
	bldr.accountNumber = accnum
	return bldr
}

// BuildSignMsg builds a single message to be signed from a TxBuilder given a
// set of messages. It returns an error if a fee is supplied but cannot be
// parsed.
func (bldr TxBuilder) BuildSignMsg(msgs []types.Msg) (StdSignMsg, error) {
	if bldr.chainID == "" {
		return StdSignMsg{}, fmt.Errorf("chain ID required but not specified")
	}

	fees := bldr.fees
	if !bldr.gasPrices.IsZero() {
		if !fees.IsZero() {
			return StdSignMsg{}, errors.New("cannot provide both fees and gas prices")
		}

		glDec := types.NewDec(int64(bldr.gas))

		// Derive the fees based on the provided gas prices, where
		// fee = ceil(gasPrice * gasLimit).
		fees = make(types.Coins, len(bldr.gasPrices))
		for i, gp := range bldr.gasPrices {
			fee := gp.Amount.Mul(glDec)
			fees[i] = types.NewCoin(gp.Denom, fee.Ceil().RoundInt())
		}
	}

	return StdSignMsg{
		ChainID:     bldr.chainID,
		Nonces:      [][]byte{bldr.nonce},
		Memo:        bldr.memo,
		Msgs:        msgs,
		Fee:         NewStdFee(bldr.gas, fees),
		ValidHeight: bldr.validHeight,
	}, nil
}

// Sign signs a transaction given a name, passphrase, and a single message to
// signed. An error is returned if signing fails.
func (bldr TxBuilder) Sign(name, passphrase string, msg StdSignMsg) ([]byte, error) {
	sig, err := MakeSignature(bldr.keybase, name, passphrase, msg)
	if err != nil {
		return nil, err
	}

	return bldr.txEncoder(NewStdTx(msg.Msgs, msg.Fee, []StdSignature{sig}, msg.Memo, [][]byte{bldr.nonce}, bldr.validHeight))
}

// BuildAndSign builds a single message to be signed, and signs a transaction
// with the built message given a name, passphrase, and a set of messages.
func (bldr TxBuilder) BuildAndSign(name, passphrase string, msgs []types.Msg) ([]byte, error) {
	msg, err := bldr.BuildSignMsg(msgs)
	if err != nil {
		return nil, err
	}

	return bldr.Sign(name, passphrase, msg)
}

// BuildTxForSim creates a StdSignMsg and encodes a transaction with the
// StdSignMsg with a single empty StdSignature for tx simulation.
func (bldr TxBuilder) BuildTxForSim(msgs []types.Msg) ([]byte, error) {
	signMsg, err := bldr.BuildSignMsg(msgs)
	if err != nil {
		return nil, err
	}

	// the ante handler will populate with a sentinel pubkey
	sigs := []StdSignature{{}}
	return bldr.txEncoder(NewStdTx(signMsg.Msgs, signMsg.Fee, sigs, signMsg.Memo, [][]byte{bldr.nonce}, bldr.validHeight))
}

// SignStdTx appends a signature to a StdTx and returns a copy of it. If append
// is false, it replaces the signatures already attached with the new signature.
func (bldr TxBuilder) SignStdTx(name, passphrase string, stdTx StdTx, appendSig bool) (signedStdTx StdTx, err error) {
	if bldr.chainID == "" {
		return StdTx{}, fmt.Errorf("chain ID required but not specified")
	}

	stdSignature, err := MakeSignature(bldr.keybase, name, passphrase, StdSignMsg{
		ChainID:     bldr.chainID,
		Nonces:      [][]byte{bldr.nonce},
		Fee:         stdTx.Fee,
		Msgs:        stdTx.GetMsgs(),
		Memo:        stdTx.GetMemo(),
		ValidHeight: bldr.validHeight,
	})
	if err != nil {
		return
	}

	sigs := stdTx.GetSignatures()
	if len(sigs) == 0 || !appendSig {
		sigs = []StdSignature{stdSignature}
	} else {
		sigs = append(sigs, stdSignature)
	}
	signedStdTx = NewStdTx(stdTx.GetMsgs(), stdTx.Fee, sigs, stdTx.GetMemo(), [][]byte{bldr.nonce}, bldr.validHeight)
	return
}

// MakeSignature builds a StdSignature given keybase, key name, passphrase, and a StdSignMsg.
func MakeSignature(keybase crkeys.Keybase, name, passphrase string,
	msg StdSignMsg) (sig StdSignature, err error) {
	if keybase == nil {
		keybase, err = crypto.NewKeyBaseFromDir(common.RootDir)
		if err != nil {
			return
		}
	}

	sigBytes, pubkey, err := keybase.Sign(name, passphrase, msg.Bytes())
	if err != nil {
		return
	}
	return StdSignature{
		PubKey:    pubkey,
		Signature: sigBytes,
	}, nil
}

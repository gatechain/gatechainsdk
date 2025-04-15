package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gatechain/crypto"
	"github.com/gatechain/crypto/multisig"
	"github.com/gatechain/crypto/tmhash"
	"github.com/gatechain/gatechainsdk/gatechain/types"

	yaml "gopkg.in/yaml.v2"
)

var (
	_ types.Tx = (*StdTx)(nil)

	maxGasWanted = uint64((1 << 63) - 1)
)

// StdTx is a standard way to wrap a Msg with Fee and Signatures.
// NOTE: the first signature is the fee payer (Signatures must not be nil).
type StdTx struct {
	Msgs        []types.Msg    `json:"msg" yaml:"msg"`
	Fee         StdFee         `json:"fee" yaml:"fee"`
	Nonces      [][]byte       `json:"nonces" yaml:"nonces"`
	Signatures  []StdSignature `json:"signatures" yaml:"signatures"`
	Memo        string         `json:"memo" yaml:"memo"`
	ValidHeight []uint64       `json:"valid_height" yaml:"valid_height"`
}

func NewStdTx(msgs []types.Msg, fee StdFee, sigs []StdSignature, memo string, nonces [][]byte, validHeight []uint64) StdTx {
	return StdTx{
		Msgs:        msgs,
		Fee:         fee,
		Signatures:  sigs,
		Memo:        memo,
		Nonces:      nonces,
		ValidHeight: validHeight,
	}
}

// GetMsgs returns the all the transaction's messages.
func (tx StdTx) GetMsgs() []types.Msg { return tx.Msgs }

// ValidateBasic does a simple and lightweight validation check that doesn't
// require access to any other information.
func (tx StdTx) ValidateBasic() types.Error {
	if msgs := tx.GetMsgs(); len(msgs) > 0 {
		if typ := msgs[0].Type(); typ == "ethereum" || typ == "ethermint" {
			return nil
		}
	}

	stdSigs := tx.GetSignatures()

	if tx.Fee.Gas > maxGasWanted {
		return types.ErrGasOverflow(fmt.Sprintf("invalid gas supplied; %d > %d", tx.Fee.Gas, maxGasWanted))
	}
	if tx.Fee.Amount.IsAnyNegative() {
		return types.ErrInsufficientFee(fmt.Sprintf("invalid fee %s amount provided", tx.Fee.Amount))
	}
	if len(stdSigs) == 0 {
		return types.ErrNoSignatures("no signers")
	}
	if len(stdSigs) != len(tx.GetSigners()) {
		return types.ErrUnauthorized("wrong number of signers")
	}

	if len(tx.ValidHeight) < 2 {
		return types.ErrInvalidExpireHeight("expire height and valid height must set")
	}

	if tx.ValidHeight[0] < 0 || tx.ValidHeight[1] <= 0 {
		return types.ErrInvalidExpireHeight("expire height must lg 0")
	}

	if tx.ValidHeight[0] > tx.ValidHeight[1] {
		return types.ErrInvalidExpireHeight("expire height must lg start height")
	}

	return nil
}

// CountSubKeys counts the total number of keys for a multi-sig public key.
func CountSubKeys(pub crypto.PubKey) int {
	v, ok := pub.(multisig.PubKeyMultisigThreshold)
	if !ok {
		return 1
	}

	numKeys := 0
	for _, subkey := range v.PubKeys {
		numKeys += CountSubKeys(subkey)
	}

	return numKeys
}

// GetSigners returns the addresses that must sign the transaction.
// Addresses are returned in a deterministic order.
// They are accumulated from the GetSigners method for each Msg
// in the order they appear in tx.GetMsgs().
// Duplicate addresses will be omitted.
func (tx StdTx) GetSigners() []types.AccAddress {
	seen := map[string]bool{}
	var signers []types.AccAddress
	for _, msg := range tx.GetMsgs() {
		for _, addr := range msg.GetSigners() {
			if !seen[addr.String()] {
				signers = append(signers, addr)
				seen[addr.String()] = true
			}
		}
	}
	return signers
}

// GetMemo returns the memo
func (tx StdTx) GetMemo() string { return tx.Memo }

// GetSignatures returns the signature of signers who signed the Msg.
// GetSignatures returns the signature of signers who signed the Msg.
// CONTRACT: Length returned is same as length of
// pubkeys returned from MsgKeySigners, and the order
// matches.
// CONTRACT: If the signature is missing (ie the Msg is
// invalid), then the corresponding signature is
// .Empty().
func (tx StdTx) GetSignatures() []StdSignature { return tx.Signatures }

// GetNonce returns the nonce
func (tx StdTx) GetNonces() [][]byte { return tx.Nonces }

// GetNonce returns the nonce
func (tx StdTx) GetExpireHeight() []uint64 { return tx.ValidHeight }

//__________________________________________________________

// StdFee includes the amount of coins paid in fees and the maximum
// gas to be used by the transaction. The ratio yields an effective "gasprice",
// which must be above some miminum to be accepted into the mempool.
type StdFee struct {
	Amount types.Coins `json:"amount" yaml:"amount"`
	Gas    uint64      `json:"gas" yaml:"gas"`
}

// NewStdFee returns a new instance of StdFee
func NewStdFee(gas uint64, amount types.Coins) StdFee {
	return StdFee{
		Amount: amount,
		Gas:    gas,
	}
}

// Bytes for signing later
func (fee StdFee) Bytes() []byte {
	// normalize. XXX
	// this is a sign of something ugly
	// (in the lcd_test, client side its null,
	// server side its [])
	if len(fee.Amount) == 0 {
		fee.Amount = types.NewCoins()
	}
	bz, err := ModuleCdc.MarshalJSON(fee) // TODO
	if err != nil {
		panic(err)
	}
	return bz
}

// GasPrices returns the gas prices for a StdFee.
//
// NOTE: The gas prices returned are not the true gas prices that were
// originally part of the submitted transaction because the fee is computed
// as fee = ceil(gasWanted * gasPrices).
func (fee StdFee) GasPrices() types.DecCoins {
	return types.NewDecCoins(fee.Amount).QuoDec(types.NewDec(int64(fee.Gas)))
}

//__________________________________________________________

// StdSignDoc is replay-prevention structure.
// It includes the result of msg.GetSignBytes(),
// as well as the ChainID (prevent cross chain replay)
// and the Sequence numbers for each signature (prevent
// inchain replay and enforce tx ordering per account).
type StdSignDoc struct {
	ChainID     string            `json:"chain_id" yaml:"chain_id"`
	Fee         json.RawMessage   `json:"fee" yaml:"fee"`
	Memo        string            `json:"memo" yaml:"memo"`
	Msgs        []json.RawMessage `json:"msgs" yaml:"msgs"`
	Nonces      [][]byte          `json:"nonces" yaml:"nonces"`
	VaildHeight []uint64          `json:"valid_height" yaml:"valid_height"`
}

// StdSignBytes returns the bytes to sign for a transaction.
func StdSignBytes(chainID string, nonces [][]byte, fee StdFee, msgs []types.Msg, memo string, validHeight []uint64) []byte {
	var msgsBytes []json.RawMessage
	for _, msg := range msgs {
		msgsBytes = append(msgsBytes, json.RawMessage(msg.GetSignBytes()))
	}
	bz, err := ModuleCdc.MarshalJSON(StdSignDoc{
		ChainID:     chainID,
		Fee:         json.RawMessage(fee.Bytes()),
		Memo:        memo,
		Msgs:        msgsBytes,
		Nonces:      nonces,
		VaildHeight: validHeight,
	})
	if err != nil {
		panic(err)
	}
	sortJson := types.MustSortJSON(bz)
	signBytes := tmhash.Sum(sortJson)
	return signBytes[:]
}

// StdSignature represents a sig
type StdSignature struct {
	crypto.PubKey `json:"pub_key" yaml:"pub_key"` // optional
	Signature     []byte                          `json:"signature" yaml:"signature"`
}

// MarshalYAML returns the YAML representation of the signature.
func (ss StdSignature) MarshalYAML() (interface{}, error) {
	var (
		bz     []byte
		pubkey string
		err    error
	)

	if ss.PubKey != nil {
		pubkey, err = types.Bech32ifyAccPub(ss.PubKey)
		if err != nil {
			return nil, err
		}
	}

	bz, err = yaml.Marshal(struct {
		PubKey    string
		Signature string
	}{
		PubKey:    pubkey,
		Signature: fmt.Sprintf("%s", ss.Signature),
	})
	if err != nil {
		return nil, err
	}

	return string(bz), err
}

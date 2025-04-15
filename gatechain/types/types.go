package types

import "C"
import (
	"crypto/rand"
	"encoding/json"
	"github.com/gatechain/crypto"
	"github.com/gatechain/crypto/tmhash"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	logging "github.com/ipfs/go-log/v2"
	"github.com/tendermint/tendermint/crypto/merkle"
)

var (
	log = logging.Logger("types")
)

func init() {
}

// DefaultTxEncoder logic for standard transaction encoding
func DefaultTxEncoder(cdc *codec.Codec) TxEncoder {
	return func(tx Tx) ([]byte, error) {
		return cdc.MarshalBinaryLengthPrefixed(tx)
	}
}

// StdFee includes the amount of coins paid in fees and the maximum
// gas to be used by the transaction. The ratio yields an effective "gasprice",
// which must be above some miminum to be accepted into the mempool.
type StdFee struct {
	Amount Coins  `json:"amount" yaml:"amount"`
	Gas    uint64 `json:"gas" yaml:"gas"`
}

// NewStdFee returns a new instance of StdFee
func NewStdFee(gas uint64, amount Coins) StdFee {
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
		fee.Amount = NewCoins()
	}
	bz, err := codec.Cdc.MarshalJSON(fee) // TODO
	if err != nil {
		panic(err)
	}
	return bz
}

// StdSignature represents a sig
type StdSignature struct {
	crypto.PubKey `json:"pub_key" yaml:"pub_key"` // optional
	Signature     []byte                          `json:"signature" yaml:"signature"`
}

// GetPubKey returns the public key of a signature as a cryptotypes.PubKey using the
// Amino codec.
func (ss StdSignature) GetPubKey() crypto.PubKey {
	return ss.PubKey
}

type ResponseQuery struct {
	Code uint32 `json:"code"`
	// bytes data = 2; // use "value" instead.
	Log   string `json:"log"`
	Info  string `json:"info"`
	Index int64  `json:"index"`
	Key   []byte `json:"key"`
	Value []byte `json:"value"`

	//TODO need to change to gt merkle tree
	Proof     *merkle.Proof `json:"proof"`
	Height    int64         `json:"height"`
	Codespace string        `json:"codespace"`
}

//// CodespaceType - codespace identifier
//type CodespaceType string

// StdSignMsg is a convenience structure for passing along
// a Msg with the other requirements for a StdSignDoc before
// it is signed. For use in the CLI.
type StdSignMsg struct {
	ChainID     string   `json:"chain_id" yaml:"chain_id"`
	Nonces      [][]byte `json:"nonces" yaml:"nonces"`
	Fee         StdFee   `json:"fee" yaml:"fee"`
	Msgs        []Msg    `json:"msgs" yaml:"msgs"`
	Memo        string   `json:"Memo" yaml:"Memo"`
	ValidHeight []uint64 `json:"valid_height" yaml:"valid_height"`
}

// get message bytes
func (msg StdSignMsg) Bytes() []byte {
	return StdSignBytes(msg.ChainID, msg.Nonces, msg.Fee, msg.Msgs, msg.Memo, msg.ValidHeight)
}

func (msg StdSignMsg) GetSignBytes() []byte {
	return StdSignBytes(msg.ChainID, msg.Nonces, msg.Fee, msg.Msgs, msg.Memo, msg.ValidHeight)
}

// StdSignBytes returns the bytes to sign for a transaction.
func StdSignBytes(chainID string, nonces [][]byte, fee StdFee, msgs []Msg, memo string, validHeight []uint64) []byte {
	var msgsBytes []json.RawMessage
	for _, msg := range msgs {
		msgsBytes = append(msgsBytes, json.RawMessage(msg.GetSignBytes()))
	}
	bz, err := codec.Cdc.MarshalJSON(StdSignDoc{
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
	sortJson := MustSortJSON(bz)
	signBytes := tmhash.Sum(sortJson)
	return signBytes[:]
}

// StdSignDoc is replay-prevention structure.
// It includes the result of msg.GetSignBytes(),
// as well as the ChainID (prevent cross chain replay)
// and the Sequence numbers for each signature (prevent
// inchain replay and enforce tx ordering per account).
type StdSignDoc struct {
	ChainID     string            `json:"chain_id" yaml:"chain_id"`
	Fee         json.RawMessage   `json:"fee" yaml:"fee"`
	Memo        string            `json:"memo" yaml:"Memo"`
	Msgs        []json.RawMessage `json:"msgs" yaml:"msgs"`
	Nonces      [][]byte          `json:"nonces" yaml:"nonces"`
	VaildHeight []uint64          `json:"valid_height" yaml:"valid_height"`
}

//// Blob is an alias of Blob from go-square.
//type Blob = squareblob.Blob

// The UUID reserved variants.
const (
	ReservedNCS       byte = 0x80
	ReservedRFC4122   byte = 0x40
	ReservedMicrosoft byte = 0x20
	ReservedFuture    byte = 0x00
)

type UUID [16]byte

// Set the two most significant bits (bits 6 and 7) of the
// clock_seq_hi_and_reserved to zero and one, respectively.
func (u *UUID) setVariant(v byte) {
	switch v {
	case ReservedNCS:
		u[8] = (u[8] | ReservedNCS) & 0xBF
	case ReservedRFC4122:
		u[8] = (u[8] | ReservedRFC4122) & 0x7F
	case ReservedMicrosoft:
		u[8] = (u[8] | ReservedMicrosoft) & 0x3F
	}
}

// Set the four most significant bits (bits 12 through 15) of the
// time_hi_and_version field to the 4-bit version number.
func (u *UUID) setVersion(v byte) {
	u[6] = (u[6] & 0xF) | (v << 4)
}

// Generate a random UUID.
func NewV4() (u *UUID, err error) {
	u = new(UUID)
	// Set all bits to randomly (or pseudo-randomly) chosen values.
	_, err = rand.Read(u[:])
	if err != nil {
		return
	}
	u.setVariant(ReservedRFC4122)
	u.setVersion(4)
	return
}

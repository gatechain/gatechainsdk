package appinterface

import (
	"fmt"
	"github.com/gatechain/crypto"
)

// Txid is a hash used to uniquely identify individual transactions
type Txid crypto.Digest

// Tx is an arbitrary byte array.
// NOTE: Tx has no types at this level, so when wire encoded it's just length-prefixed.
// Might we want types here ?
type Tx []byte

// Hash computes the TMHASH hash of the wire encoded transaction.
func (tx Tx) Hash() []byte {
	txId := tx.ComputeID()
	return txId[:]
}

// String returns the hex-encoded transaction as a string.
func (tx Tx) String() string {
	return fmt.Sprintf("Tx{%X}", []byte(tx))
}

func (tx Tx) ComputeID() Txid {
	return Txid(crypto.Hash(tx))
}

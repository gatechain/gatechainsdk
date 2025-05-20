package types

import (
	_ "fmt"
	_ "time"

	_ "github.com/gatechain/gatechainsdk/gatechain/auth/exported"
)

// -----------------------------------------------------------------------------
// Revoke tx
type RevocableTx struct {
	Hash         []byte `json:"tx_hash" yaml:"tx_hash"`
	RevokeTxHash []byte `json:"revoke_hash" yaml:"revoke_hash"`
	RevokeHeight uint64 `json:"revoke_height" yaml:"revoke_height"`
}

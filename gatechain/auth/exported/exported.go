package exported

import (
	"github.com/gatechain/crypto"

	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// Account is an interface used to store coins at a given address within state.
// It presumes a notion of sequence numbers for replay protection,
// a notion of account numbers for replay protection for previously pruned accounts,
// and a pubkey for authentication purposes.
//
// Many complex conditions can be used in the concrete struct which implements Account.
type Account interface {
	GetPubKey() crypto.PubKey // can return nil.
	GetAccountNumber() uint64
	GetCoins() types.Coins
	// Ensure that account implements stringer
	String() string
}

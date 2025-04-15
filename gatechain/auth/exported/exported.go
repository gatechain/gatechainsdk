package exported

import (
	"github.com/gatechain/gatechainsdk/gatechain/types"
	"time"

	"github.com/gatechain/crypto"
)

// Account is an interface used to store coins at a given address within state.
// It presumes a notion of sequence numbers for replay protection,
// a notion of account numbers for replay protection for previously pruned accounts,
// and a pubkey for authentication purposes.
//
// Many complex conditions can be used in the concrete struct which implements Account.
type Account interface {
	GetAddress() types.AccAddress
	SetAddress(types.AccAddress) error // errors if already set.

	GetPubKey() crypto.PubKey // can return nil.
	SetPubKey(crypto.PubKey) error

	GetAccountNumber() uint64
	SetAccountNumber(uint64) error

	GetSequence() uint64
	SetSequence(uint64) error

	GetCoins() types.Coins
	SetCoins(types.Coins) error

	// Calculates the amount of coins that can be sent to other accounts given
	// the current time.
	SpendableCoins(blockTime time.Time) types.Coins

	// Ensure that account implements stringer
	String() string
}

// VestingAccount defines an account type that vests coins via a vesting schedule.
type VestingAccount interface {
	Account

	// Delegation and undelegation accounting that returns the resulting base
	// coins amount.
	TrackDelegation(blockTime time.Time, amount types.Coins)
	TrackUndelegation(amount types.Coins)

	GetVestedCoins(blockTime time.Time) types.Coins
	GetVestingCoins(blockTime time.Time) types.Coins

	GetStartTime() int64
	GetEndTime() int64

	GetOriginalVesting() types.Coins
	GetDelegatedFree() types.Coins
	GetDelegatedVesting() types.Coins
}

package auth

import (
	"fmt"

	"github.com/gatechain/crypto"

	"github.com/gatechain/gatechainsdk/gatechain/auth/exported"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

//-----------------------------------------------------------------------------
// BaseAccount

var _ exported.Account = (*BaseAccount)(nil)

// BaseAccount - a base account structure.
// This can be extended by embedding within in your AppAccount.
// However one doesn't have to use BaseAccount as long as your struct
// implements Account.
type BaseAccount struct {
	Address       types.AccAddress `json:"address" yaml:"address"`
	Coins         types.Coins      `json:"tokens" yaml:"tokens"`
	PubKey        crypto.PubKey    `json:"public_key" yaml:"public_key"`
	AccountNumber uint64           `json:"account_number" yaml:"account_number"`
	Sequence      uint64           `json:"sequence" yaml:"sequence"`
}

// String implements fmt.Stringer
func (acc BaseAccount) String() string {
	var pubkey string

	if acc.PubKey != nil {
		pubkey = types.MustBech32ifyAccPub(acc.PubKey)
	}

	return fmt.Sprintf(`Account:
 Address:       %s
 Pubkey:        %s
 Coins:         %s
 AccountNumber: %d
 Sequence:      %d`,
		acc.Address, pubkey, acc.Coins, acc.AccountNumber, acc.Sequence,
	)
}

// GetPubKey - Implements framework.Account.
func (acc BaseAccount) GetPubKey() crypto.PubKey {
	return acc.PubKey
}

// GetCoins - Implements framework.Account.
func (acc *BaseAccount) GetCoins() types.Coins {
	return acc.Coins
}

// SetCoins - Implements framework.Account.
func (acc *BaseAccount) SetCoins(coins types.Coins) error {
	acc.Coins = coins
	return nil
}

// GetAccountNumber - Implements Account
func (acc *BaseAccount) GetAccountNumber() uint64 {
	return acc.AccountNumber
}

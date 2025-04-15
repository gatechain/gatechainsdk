package types

import (
	"github.com/gatechain/crypto"
)

type Account struct {
	Name          string
	Coins         Coins
	PubKey        crypto.PubKey
	Address       AccAddress
	Sequence      uint64
	AccountNumber uint64
}

func NewAccount(keyName string, accountNumber, sequenceNumber uint64) *Account {
	return &Account{
		Name:          keyName,
		AccountNumber: accountNumber,
		Sequence:      sequenceNumber,
	}
}

func (a *Account) Copy() *Account {
	return &Account{
		Name:          a.Name,
		Address:       a.Address,
		PubKey:        a.PubKey,
		AccountNumber: a.AccountNumber,
		Sequence:      a.Sequence,
	}
}

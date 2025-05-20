// Package hd provides basic functionality Hierarchical Deterministic Wallets.
//
// The user must understand the overall concept of the BIP 32 and the BIP 44 specs:
//
//	https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
//	https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki
//
// In combination with the bip39 package in go-crypto this package provides the functionality for deriving keys using a
// BIP 44 HD path, or, more general, by passing a BIP 32 path.
//
// In particular, this package (together with bip39) provides all necessary functionality to derive keys from
// mnemonics generated during the  fundraiser.
package hd

import (
	"fmt"
)

// BIP44Params wraps BIP 44 params (5 level BIP 32 path).
// To receive a canonical string representation ala
// m / purpose' / coinType' / account' / change / addressIndex
// call String() on a BIP44Params instance.
type BIP44Params struct {
	Purpose      uint32 `json:"purpose"`
	CoinType     uint32 `json:"coinType"`
	Account      uint32 `json:"account"`
	Change       bool   `json:"change"`
	AddressIndex uint32 `json:"addressIndex"`
}

// NewParams creates a BIP 44 parameter object from the params:
// m / purpose' / coinType' / account' / change / addressIndex
func NewParams(purpose, coinType, account uint32, change bool, addressIdx uint32) *BIP44Params {
	return &BIP44Params{
		Purpose:      purpose,
		CoinType:     coinType,
		Account:      account,
		Change:       change,
		AddressIndex: addressIdx,
	}
}

// NewFundraiserParams creates a BIP 44 parameter object from the params:
// m / 44' / coinType' / account' / 0 / address_index
// The fixed parameters (purpose', coin_type', and change) are determined by what was used in the fundraiser.
func NewFundraiserParams(account, coinType, addressIdx uint32) *BIP44Params {
	return NewParams(44, coinType, account, false, addressIdx)
}

func (p BIP44Params) String() string {
	var changeStr string
	if p.Change {
		changeStr = "1"
	} else {
		changeStr = "0"
	}
	// m / Purpose' / coin_type' / Account' / Change / address_index
	return fmt.Sprintf("%d'/%d'/%d'/%s/%d",
		p.Purpose,
		p.CoinType,
		p.Account,
		changeStr,
		p.AddressIndex)
}

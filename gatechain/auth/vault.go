package auth

import (
	"errors"
	_ "fmt"
	"github.com/gatechain/crypto"
	_ "time"

	exported "github.com/gatechain/gatechainsdk/gatechain/auth/exported"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

//-----------------------------------------------------------------------------
// VaultAccount

var _ exported.VaultAccount = (*VaultAccount)(nil)

// VaultAccount - a account that can send innrecovable tx.
type VaultAccount struct {
	BaseAccount      `json:"base_account" yaml:"base_account"`
	Clear            exported.ClearingHeight `json:"clearing_height" yaml:"clearing_height"`
	DelayHeight      uint64                  `json:"delay_height" yaml:"delay_height"`
	RevocabledTokens []*exported.RevocableTx `json:"received_revocable_tokens" yaml:"received_revocable_tokens"`
	VaultAddress     []types.AccTypeAddress  `json:"vault_address" yaml:"vault_address"`
	SecurityAddress  types.AccTypeAddress    `json:"security_address" yaml:"security_address"`
	RevocableTokens  []*exported.RevocableTx `json:"sent_revocable_tokens" yaml:"sent_revocable_tokens"`
}

// GetSecurityAddress - Implements sdk.Account.
func (vault VaultAccount) GetSecurityAddress() types.AccTypeAddress {
	return vault.SecurityAddress
}

// GetAddress - Implements framework.Account.
func (vault VaultAccount) GetAddress() types.AccAddress {
	return vault.Address
}

// GetAccountType - Implements sdk.Account.
func (vault VaultAccount) GetAccountType() uint8 {
	isMultiSig := IsMultiSignAccount(&vault)
	isVault := IsVaultAccount(&vault)
	if isMultiSig && isVault {
		return types.MultiSignerVaultAccount
	} else if isMultiSig && !isVault {
		return types.MultiSignerStandardAccount
	} else if !isMultiSig && isVault {
		return types.VaultAccount_type
	} else if !isMultiSig && !isVault {
		return types.StandardAccount
	}
	return types.StandardAccount
}

func (vault *VaultAccount) MergeRevocableWei(height uint64) error {
	coins := types.Coins{}
	index := 0
	delays := new(exported.RevocableTx)
	for index, delays = range vault.RevocabledTokens {
		if delays.Height > height {
			break
		}
		for _, tx := range delays.Txs {
			txCoins := tx.Coins
			if len(vault.GetAddress()) == types.EthAddrLen && !txCoins.AmountOf(types.DefaultBondDenom).IsZero() {
				txCoins = txCoins.NanoToWei()
			}
			coins = coins.Add(txCoins)
		}
	}

	if (delays.Height <= height) && index == len(vault.RevocabledTokens)-1 {
		index = index + 1
	}

	switch vault.GetAccountType() {
	case types.StandardAccount, types.MultiSignerStandardAccount:
		vault.SetCoins(vault.GetCoins().Add(coins))
		vault.RevocabledTokens = vault.RevocabledTokens[index:]
	case types.VaultAccount_type, types.MultiSignerVaultAccount:
		vault.SetCoins(vault.GetCoins().Add(coins))
		vault.RevocabledTokens = vault.RevocabledTokens[index:]
		index := 0
		revocables := new(exported.RevocableTx)
		for index, revocables = range vault.RevocableTokens {
			if revocables.Height > height {
				break
			}
		}
		if (revocables.Height <= height) && index == len(vault.RevocableTokens)-1 {
			index = index + 1
		}
		vault.RevocableTokens = vault.RevocableTokens[index:]
	default:
		return errors.New("merge revocable coins failed,unknown account type")
	}
	return nil
}

// GetRevocableTokensDetail - Implements sdk.Account.
// return revocable coins detail for standard account
func (vault *VaultAccount) GetRevocableTokensDetail(height int64) (exported.RevocableTxCoinsArray, error) {
	switch vault.GetAccountType() {
	case types.StandardAccount, types.MultiSignerStandardAccount:
		return nil, errors.New("standard account can not have revocable coins")
	case types.VaultAccount_type, types.MultiSignerVaultAccount:
		var coins exported.RevocableTxCoinsArray
		for _, revocable := range vault.RevocableTokens {
			if revocable.Height <= uint64(height) {
				continue
			}

			for _, tx := range revocable.Txs {
				coins = append(coins, tx)
			}

		}
		return coins, nil
	default:
		// do nothing
	}
	return nil, errors.New("get revocable coins detail failed")
}

// GetDelayHeight - Implements Account
func (vault VaultAccount) GetDelayHeight() uint64 {
	return vault.DelayHeight
}

func (vault *VaultAccount) GetPubKey() crypto.PubKey {
	return vault.PubKey
}

func (vault *VaultAccount) GetCoins() types.Coins {
	return vault.Coins
}

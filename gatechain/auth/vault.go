package auth

import (
	"errors"
	"fmt"
	_ "fmt"
	exported2 "github.com/gatechain/gatechainsdk/gatechain/auth/exported"
	"github.com/gatechain/gatechainsdk/gatechain/types"
	_ "time"

	"github.com/gatechain/crypto"
	"gopkg.in/yaml.v2"
)

//-----------------------------------------------------------------------------
// VaultAccount

var _ exported2.VaultAccount = (*VaultAccount)(nil)

// VaultAccount - a account that can send innrecovable tx.
type VaultAccount struct {
	BaseAccount      `json:"base_account" yaml:"base_account"`
	Clear            exported2.ClearingHeight `json:"clearing_height" yaml:"clearing_height"`
	DelayHeight      uint64                   `json:"delay_height" yaml:"delay_height"`
	RevocabledTokens []*exported2.RevocableTx `json:"received_revocable_tokens" yaml:"received_revocable_tokens"`
	VaultAddress     []types.AccTypeAddress   `json:"vault_address" yaml:"vault_address"`
	SecurityAddress  types.AccTypeAddress     `json:"security_address" yaml:"security_address"`
	RevocableTokens  []*exported2.RevocableTx `json:"sent_revocable_tokens" yaml:"sent_revocable_tokens"`
}

// String implements fmt.Stringer
func (vault VaultAccount) String() string {

	address := vault.Address.TypeString(vault.GetAccountType())
	var pubkey string

	if vault.PubKey != nil {
		pubkey = types.MustBech32ifyAccPub(vault.PubKey)
	}
	if vault.GetAccountType() == types.StandardAccount {

		return fmt.Sprintf(`Account:
  Address:            %s
  Pubkey:        %s
  Tokens:        %s
  AccountNumber: %d
  Sequence:      %d
  AccountType:        %d
  VaultAddress:   %s
  ReceivedRevocableTokens:         %s`,

			address, pubkey, vault.Coins, vault.AccountNumber, vault.Sequence,
			vault.GetAccountType(), vault.VaultAddress, vault.GetRevocabledTokensTotal(),
		)
	} else if vault.GetAccountType() == types.VaultAccount_type {
		return fmt.Sprintf(`Account:
  Address:            %s
  Pubkey:        %s
  Tokens:        %s
  AccountNumber: %d
  Sequence:      %d
  AccountType:        %d
  DelayHeight:        %d
  SecurityAddress:    %s
  LastClearingHeight: %d
  LastClearingEffectHeight:   %d
  NextClearingHeight:      %d
  NextClearingEffectHeight:   %d
  ReceivedRevocableTokens:         %s
  SentRevocableTokens:      %s`,

			address, pubkey, vault.Coins, vault.AccountNumber, vault.Sequence, vault.GetAccountType(), vault.DelayHeight,
			vault.SecurityAddress, vault.Clear.LastClearingHeight, vault.Clear.LastClearingEffectHeight,
			vault.Clear.NextClearingHeight, vault.Clear.NextClearingEffectHeight, vault.GetRevocabledTokensTotal(), vault.getRevocableTokens(),
		)
	} else if vault.GetAccountType() == types.MultiSignerStandardAccount {

		return fmt.Sprintf(`Account:
  Address:            %s
  Pubkey:        %s
  Tokens:        %s
  AccountNumber: %d
  Sequence:      %d
  AccountType:        %d
  VaultAddress:   %s
  ReceivedRevocableTokens:         %s`,
			//  Signers:            %s
			//  MinSignatureCount:  %d,

			address, pubkey, vault.Coins, vault.AccountNumber, vault.Sequence, vault.GetAccountType(),
			vault.VaultAddress, vault.GetRevocabledTokensTotal(),
			// acc.Signers.Signers, acc.Signers.MinSigCount,
		)
	} else if vault.GetAccountType() == types.MultiSignerVaultAccount {
		return fmt.Sprintf(`Account:
  Address:            %s
  Pubkey:        %s
  Tokens:        %s
  AccountNumber: %d
  Sequence:      %d
  AccountType:        %d
  DelayHeight:        %d
  SecurityAddress:    %s
  LastClearingHeight:      %d
  LastClearingEffectHeight:   %d
  NextClearingHeight:      %d
  NextClearingEffectHeight:   %d
  ReceivedRevocableTokens:         %s
  SentRevocableTokens:      %s`,
			//  Signers:            %s
			//  MinSignatureCount:  %d

			address, pubkey, vault.Coins, vault.AccountNumber, vault.Sequence, vault.GetAccountType(), vault.DelayHeight,
			vault.SecurityAddress, vault.Clear.LastClearingHeight, vault.Clear.LastClearingEffectHeight,
			vault.Clear.NextClearingHeight, vault.Clear.NextClearingEffectHeight, vault.GetRevocabledTokensTotal(), vault.getRevocableTokens(),
			//			acc.Signers.Signers, acc.Signers.MinSigCount,
		)
	} else {
		return "Error account"
	}
}

func (acc VaultAccount) MarshalYAML() (interface{}, error) {
	var bs []byte
	var err error
	var pubkey string

	if acc.PubKey != nil {
		pubkey, err = types.Bech32ifyAccPub(acc.PubKey)
		if err != nil {
			return nil, err
		}
	}

	bs, err = yaml.Marshal(struct {
		Address          string
		Coins            types.Coins
		PubKey           string
		AccountNumber    uint64
		Sequence         uint64
		Type             uint8
		Clear            exported2.ClearingHeight
		DelayHeight      uint64
		RevocabledTokens []*exported2.RevocableTx
		VaultAddress     []types.AccTypeAddress
		SecurityAddress  types.AccTypeAddress
		RevocableTokens  []*exported2.RevocableTx
	}{
		Address:          acc.Address.TypeString(acc.GetAccountType()),
		Coins:            acc.Coins,
		PubKey:           pubkey,
		AccountNumber:    acc.AccountNumber,
		Sequence:         acc.Sequence,
		Type:             acc.GetAccountType(),
		Clear:            acc.Clear,
		DelayHeight:      acc.DelayHeight,
		RevocabledTokens: acc.RevocabledTokens,
		VaultAddress:     acc.VaultAddress,
		SecurityAddress:  acc.SecurityAddress,
		RevocableTokens:  acc.RevocableTokens,
	})
	if err != nil {
		return nil, err
	}

	return string(bs), err
}

// ProtoBaseAccount - a prototype function for BaseAccount
func ProtoVaultAccount() exported2.Account {
	return &VaultAccount{}
}

// NewBaseAccount creates a new BaseAccount object
func NewVaultAccount(address types.AccAddress, coins types.Coins,
	pubKey crypto.PubKey, accountNumber uint64, sequence uint64, clear exported2.ClearingHeight,
	delayHeight uint64, revocabledCoins []*exported2.RevocableTx, vaultAddress []types.AccTypeAddress,
	securityAddress types.AccTypeAddress, revocableCoins []*exported2.RevocableTx) *VaultAccount {
	baseAcc := NewBaseAccount(address, coins, pubKey, accountNumber, sequence)
	return &VaultAccount{
		BaseAccount:      *baseAcc,
		Clear:            clear,
		DelayHeight:      delayHeight,
		RevocabledTokens: revocabledCoins,
		VaultAddress:     vaultAddress,
		SecurityAddress:  securityAddress,
		RevocableTokens:  revocableCoins,
	}
}

// NewBaseAccountWithAddress - returns a new base account with a given address
func NewVaultAccountWithAddress(addr types.AccAddress) *VaultAccount {
	baseAcc := NewBaseAccountWithAddress(addr)
	return &VaultAccount{
		BaseAccount: baseAcc,
	}
}

// GetSecurityAddress - Implements sdk.Account.
func (vault VaultAccount) GetSecurityAddress() types.AccTypeAddress {
	return vault.SecurityAddress
}

// GetVaultAddress - Implements sdk.Account.
func (vault VaultAccount) GetVaultAddress() []types.AccTypeAddress {
	return vault.VaultAddress
}

// ExistRevocabledTokens - Implements sdk.Account.
// check account is exist revocabledTxCoins (standard account and vault account)
func (vault *VaultAccount) ExistRevocabledTokens(revocabledTxCoins exported2.RevocableTxCoins) bool {
	for _, delays := range vault.RevocabledTokens {
		if delays.Height != revocabledTxCoins.Height {
			continue
		}
		for _, tx := range delays.Txs {
			if (revocabledTxCoins.TxHash == tx.TxHash) &&
				(revocabledTxCoins.Index == tx.Index) &&
				(revocabledTxCoins.Height == tx.Height) &&
				(revocabledTxCoins.Coins.IsEqual(tx.Coins)) {
				return true
			}
		}
		if delays.Height > revocabledTxCoins.Height {
			return false
		}
	}

	return false
}

// ExistRevocableTokens - Implements sdk.Account.
// check account is exist revocabledTxCoins ( vault account)
func (vault *VaultAccount) ExistRevocableTokens(revocabledTxCoins exported2.RevocableTxCoins) bool {
	if vault.GetAccountType() != types.VaultAccount_type && vault.GetAccountType() != types.MultiSignerVaultAccount {
		return false
	}

	for _, revocables := range vault.RevocableTokens {
		if revocables.Height != revocabledTxCoins.Height {
			continue
		}
		for _, tx := range revocables.Txs {
			if (revocabledTxCoins.TxHash == tx.TxHash) &&
				(revocabledTxCoins.Index == tx.Index) &&
				(revocabledTxCoins.Height == tx.Height) &&
				(revocabledTxCoins.Coins.IsEqual(tx.Coins)) {
				return true
			}
		}
		if revocables.Height > revocabledTxCoins.Height {
			return false
		}
	}

	return false
}

// NeedClear  - Implements sdk.Account.
// return this vault account whether clear height is reached
func (vault *VaultAccount) NeedClear(height uint64) bool {
	switch vault.GetAccountType() {
	case types.StandardAccount, types.MultiSignerStandardAccount:
		return false
	case types.VaultAccount_type, types.MultiSignerVaultAccount:
		return vault.Clear.NeedClear(height)
	default:
		return false
	}
}

// SetClearingHeight - Implements sdk.Account.
// set the new clear height for the vault account
func (vault *VaultAccount) SetClearingHeight(clearingHeight uint64, effectHeight uint64, blockHeight uint64, delayArriveTime uint64) error {
	switch vault.GetAccountType() {
	case types.StandardAccount, types.MultiSignerStandardAccount:
		return errors.New("cannot set clear height for standard account")
	case types.VaultAccount_type, types.MultiSignerVaultAccount:
		return vault.Clear.SetClearingHeight(clearingHeight, effectHeight, blockHeight, delayArriveTime)
	default:
		return errors.New("set clear height failed")
	}
}

// IsCanSetClearingHeight - Implements sdk.Account.
// check whether can set clear height for vault account
func (vault *VaultAccount) IsCanSetClearingHeight(blockHeight uint64, effectHeight uint64) bool {
	switch vault.GetAccountType() {
	case types.StandardAccount, types.MultiSignerStandardAccount:
		return false
	case types.VaultAccount_type, types.MultiSignerVaultAccount:
		return vault.Clear.IsCanSet(blockHeight, effectHeight)
	default:
		return false
	}
}

// RemoveRevocabledTokens - Implements sdk.Account.
// remove the delay coins in account (standard account and vault account, for delay coins)
func (vault *VaultAccount) RemoveRevocabledTokens(revocabledTxCoins exported2.RevocableTxCoins) error {
	foundCoins := false
	delayIndex := 0
	txIndex := 0
	delays := new(exported2.RevocableTx)
	tx := exported2.RevocableTxCoins{}

	for delayIndex, delays = range vault.RevocabledTokens {
		if delays.Height < revocabledTxCoins.Height {
			continue
		}
		for txIndex, tx = range delays.Txs {
			if (revocabledTxCoins.TxHash == tx.TxHash) &&
				(revocabledTxCoins.Index == tx.Index) &&
				(revocabledTxCoins.Height == tx.Height) &&
				(revocabledTxCoins.Coins.IsEqual(tx.Coins)) {
				foundCoins = true
				break
			}
		}

		if foundCoins {
			break
		}

		if delays.Height > revocabledTxCoins.Height {
			return errors.New("delay coins not found")
		}
	}

	if foundCoins {
		if len(vault.RevocabledTokens[delayIndex].Txs) == 1 && txIndex == 0 {
			// remove this RevocabledTokens
			vault.RevocabledTokens = append(vault.RevocabledTokens[:delayIndex], vault.RevocabledTokens[delayIndex+1:]...)
		} else {
			vault.RevocabledTokens[delayIndex].Txs = append(vault.RevocabledTokens[delayIndex].Txs[:txIndex], vault.RevocabledTokens[delayIndex].Txs[txIndex+1:]...)
		}
		return nil
	}
	return errors.New("delay coins not found")
}

// RemoveRevocableTokens - Implements sdk.Account.
// remove the delay coins in account (standard account and vault account, for revocable coins)
func (vault *VaultAccount) RemoveRevocableTokens(revocabledTxCoins exported2.RevocableTxCoins) error {
	if vault.GetAccountType() != types.VaultAccount_type && vault.GetAccountType() != types.MultiSignerVaultAccount {
		return errors.New("standard account do not have revocable coins")
	}

	foundCoins := false
	revocableIndex := 0
	txIndex := 0
	revocables := new(exported2.RevocableTx)
	tx := exported2.RevocableTxCoins{}

	for revocableIndex, revocables = range vault.RevocableTokens {
		if revocables.Height < revocabledTxCoins.Height {
			continue
		}
		for txIndex, tx = range revocables.Txs {
			if (revocabledTxCoins.TxHash == tx.TxHash) &&
				(revocabledTxCoins.Index == tx.Index) &&
				(revocabledTxCoins.Height == tx.Height) &&
				(revocabledTxCoins.Coins.IsEqual(tx.Coins)) {
				foundCoins = true
				break
			}
		}

		if foundCoins {
			break
		}

		if revocables.Height > revocabledTxCoins.Height {
			return errors.New("revocable coins not found")
		}
	}

	if foundCoins {
		if len(vault.RevocableTokens[revocableIndex].Txs) == 1 && txIndex == 0 {
			// remove this RevocabledTokens
			vault.RevocableTokens = append(vault.RevocableTokens[:revocableIndex], vault.RevocableTokens[revocableIndex+1:]...)
		} else {
			vault.RevocableTokens[revocableIndex].Txs = append(vault.RevocableTokens[revocableIndex].Txs[:txIndex], vault.RevocableTokens[revocableIndex].Txs[txIndex+1:]...)

		}
		return nil
	}
	return errors.New("revocable coins not found")
}

// GetAddress - Implements framework.Account.
func (vault VaultAccount) GetAddress() types.AccAddress {
	return vault.Address
}

// SetAddress - Implements framework.Account.
func (vault *VaultAccount) SetAddress(addr types.AccAddress) error {
	if len(vault.Address) != 0 {
		return errors.New("cannot override BaseAccount address")
	}
	vault.Address = addr
	return nil
}

// SetAddress - Implements framework.Account.
func (vault *VaultAccount) GetClear() exported2.ClearingHeight {
	return vault.Clear
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

// MergeRevocableTokens - Implements sdk.Account.
// merge all coins for account
func (vault *VaultAccount) MergeRevocableTokens(height uint64) error {
	coins := types.Coins{}
	index := 0
	delays := new(exported2.RevocableTx)
	for index, delays = range vault.RevocabledTokens {
		if delays.Height > height {
			break
		}

		for _, tx := range delays.Txs {
			coins = coins.Add(tx.Coins)
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
		revocables := new(exported2.RevocableTx)
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

func (vault *VaultAccount) MergeRevocableWei(height uint64) error {
	coins := types.Coins{}
	index := 0
	delays := new(exported2.RevocableTx)
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
		revocables := new(exported2.RevocableTx)
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

// SetSecurityAddress - Implements sdk.Account.
func (vault *VaultAccount) SetSecurityAddress(addr types.AccTypeAddress) error {
	if !vault.SecurityAddress.Empty() {
		return errors.New("cannot override Security Address ")
	}

	if vault.Address.Equals(addr.Address) {
		return errors.New("security address invalid")
	}
	vault.SecurityAddress = addr
	return nil
}

// AddVaultAddress - Implements sdk.Account.
func (vault *VaultAccount) AddVaultAddress(addr types.AccTypeAddress) error {
	if (addr.Type != types.VaultAccount_type) && (addr.Type != types.MultiSignerVaultAccount) {
		return errors.New("vault address invalid")
	}

	// Sort vault address
	index := 0
	breakOut := false
	address := types.AccTypeAddress{}
	for index, address = range vault.VaultAddress {
		if address.Equals(addr) {
			return errors.New("vault account has been bond")
		}

		if address.String() > addr.String() {
			breakOut = true
			break
		}
	}
	if breakOut {
		rear := append([]types.AccTypeAddress{}, vault.VaultAddress[index:]...)
		vault.VaultAddress = append(append(vault.VaultAddress[:index], addr), rear...)
	} else {
		vault.VaultAddress = append(vault.VaultAddress, addr)
	}

	return nil
}

// GetRevocabledTokensTotal - Implements sdk.Account.
// return delay coins for standard account
func (vault *VaultAccount) GetRevocabledTokensTotal() types.Coins {
	coins := types.Coins{}
	for _, delay := range vault.RevocabledTokens {
		for _, delayCoin := range delay.Txs {
			coins = coins.Add(delayCoin.Coins)
		}
	}
	return coins
}

// GetRevocabledTokens - Implements sdk.Account.
// return delay coins for standard account
func (vault *VaultAccount) GetRevocabledTokens() []*exported2.RevocableTx {
	return vault.RevocabledTokens
}

// GetRevocableTokens - Implements sdk.Account.
// return revocable coins for standard account
func (vault *VaultAccount) GetRevocableTokens() []*exported2.RevocableTx {
	return vault.RevocableTokens
}

// GetRevocableTokens - Implements sdk.Account.
// return revocable coins for standard account
func (vault *VaultAccount) getRevocableTokens() types.Coins {
	switch vault.GetAccountType() {
	case types.StandardAccount, types.MultiSignerStandardAccount:
		return types.Coins{}
	case types.VaultAccount_type, types.MultiSignerVaultAccount:
		coins := types.Coins{}
		for _, revocable := range vault.RevocableTokens {
			for _, revocableCoin := range revocable.Txs {
				coins = coins.Add(revocableCoin.Coins)
			}
		}
		return coins
	default:
		return types.Coins{}
	}
}

// GetRevocableTokensDetail - Implements sdk.Account.
// return revocable coins detail for standard account
func (vault *VaultAccount) GetRevocableTokensDetail(height int64) (exported2.RevocableTxCoinsArray, error) {
	switch vault.GetAccountType() {
	case types.StandardAccount, types.MultiSignerStandardAccount:
		return nil, errors.New("standard account can not have revocable coins")
	case types.VaultAccount_type, types.MultiSignerVaultAccount:
		var coins exported2.RevocableTxCoinsArray
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

// AddRevocabledTokens - Implements Account
func (vault *VaultAccount) AddRevocabledTokens(delay exported2.RevocableTxCoins) error {
	delayLength := len(vault.RevocabledTokens)
	if delayLength > 0 {
		for index, delayCoins := range vault.RevocabledTokens {
			if delayCoins.Height > delay.Height {
				rear := append([]*exported2.RevocableTx{}, vault.RevocabledTokens[index:]...)
				vault.RevocabledTokens = append(append(vault.RevocabledTokens[:index], exported2.NewRevocableTokens(delay)), rear...)
				return nil
			} else if delayCoins.Height == delay.Height {
				if err := vault.RevocabledTokens[index].AddRevocableTokens(delay); err != nil {
					return err
				}
				return nil
			}
		}
		vault.RevocabledTokens = append(vault.RevocabledTokens, exported2.NewRevocableTokens(delay))
	} else {
		vault.RevocabledTokens = append(vault.RevocabledTokens, exported2.NewRevocableTokens(delay))
	}

	return nil
}

// AddRevocableTokens - Implements Account
func (vault *VaultAccount) AddRevocableTokens(delay exported2.RevocableTxCoins) error {
	if vault.GetAccountType() != types.VaultAccount_type && vault.GetAccountType() != types.MultiSignerVaultAccount {
		return errors.New("standard account can not have revocable coins")
	}

	revocableLength := len(vault.RevocableTokens)
	if revocableLength > 0 {
		revocableCoins := vault.RevocableTokens[revocableLength-1]
		if revocableCoins.Height < delay.Height {
			vault.RevocableTokens = append(vault.RevocableTokens, exported2.NewRevocableTokens(delay))
		} else if revocableCoins.Height == delay.Height {
			if err := vault.RevocableTokens[revocableLength-1].AddRevocableTokens(delay); err != nil {
				return err
			}
		} else {
			// Can not be true
			return errors.New("delay tx height invalid")
		}

	} else {
		vault.RevocableTokens = append(vault.RevocableTokens, exported2.NewRevocableTokens(delay))
	}

	return nil
}

// GetDelayHeight - Implements Account
func (vault VaultAccount) GetDelayHeight() uint64 {
	return vault.DelayHeight
}

// SetDelayHeight - Implements Account
func (vault *VaultAccount) SetDelayHeight(delayHeight uint64) error {
	vault.DelayHeight = delayHeight
	return nil
}

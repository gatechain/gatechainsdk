package types

import (
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// ------------------------------------------------------------------------------------------
// MsgCreateVault msg for create vault account
type MsgCreateVault struct {
	FromAddress     types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress       types.AccAddress `json:"to_address" yaml:"to_address"`
	SecurityAddress types.AccAddress `json:"security_address" yaml:"security_address"`
	DelayHeight     uint64           `json:"delay_height" yaml:"delay_height"`
	ClearingHeight  uint64           `json:"clearing_height" yaml:"clearing_height"`
	Amount          types.Coins      `json:"amount" yaml:"amount"`
	Pubkey          string           `json:"pubkey" yaml:"pubkey"`
}

var _ types.Msg = MsgCreateVault{}

// NewMsgCreateVault - construct arbitrary multi-in, multi-out send msg.
func NewMsgCreateVault(fromAddr, toAddr types.AccAddress, securityAddress types.AccAddress, delayHeight uint64, clearingHeight uint64, amount types.Coins, pubkey string) MsgCreateVault {
	return MsgCreateVault{FromAddress: fromAddr, ToAddress: toAddr, SecurityAddress: securityAddress, DelayHeight: delayHeight, ClearingHeight: clearingHeight, Amount: amount, Pubkey: pubkey}
}

// Type Implements Msg.
func (msg MsgCreateVault) Type() string { return "vault" }

// GetSignBytes Implements Msg.
func (msg MsgCreateVault) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgCreateVault) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}

// ------------------------------------------------------------------------------------------
// MsgRevocableSend - high level transaction of the coin module
// insurance account send coin to common account
type MsgRevocableSend struct {
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      types.Coins      `json:"amount" yaml:"amount"`
}

var _ types.Msg = MsgRevocableSend{}

// NewMsgRevocableSend - construct arbitrary multi-in, multi-out send msg.
func NewMsgRevocableSend(fromAddr, toAddr types.AccAddress, amount types.Coins) MsgRevocableSend {
	return MsgRevocableSend{FromAddress: fromAddr, ToAddress: toAddr, Amount: amount}
}

// Type Implements Msg.
func (msg MsgRevocableSend) Type() string { return "revocable" }

// GetSignBytes Implements Msg.
func (msg MsgRevocableSend) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgRevocableSend) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}

// ------------------------------------------------------------------------------------------
// MsgRevoke - high level transaction of the coin module
// revoke  coins to security account
type MsgRevoke struct {
	VaultAddress    types.AccAddress `json:"vault_address" yaml:"vault_address"`
	SecurityAddress types.AccAddress `json:"security_address" yaml:"security_address"`
	RevokeAddress   types.AccAddress `json:"revoke_address" yaml:"revoke_address"`
	Height          int64            `json:"height" yaml:"height"`
	TxHash          string           `json:"tx_hash" yaml:"tx_hash"`
	MsgIndex        int64            `json:"msg_index" yaml:"msg_index"`
	Amount          types.Coins      `json:"amount" yaml:"amount"`
}

var _ types.Msg = MsgRevoke{}

// MsgRevoke - construct arbitrary multi-in, multi-out send msg.
func NewMsgRevoke(vaultAddr, securityAddress types.AccAddress, revokeAddress types.AccAddress, height int64, msgIndex int64, hash string, amount types.Coins) MsgRevoke {
	return MsgRevoke{VaultAddress: vaultAddr, SecurityAddress: securityAddress,
		RevokeAddress: revokeAddress, Height: height, MsgIndex: msgIndex, TxHash: hash, Amount: amount}
}

// Type Implements Msg.
func (msg MsgRevoke) Type() string { return "revoke" }

// GetSignBytes Implements Msg.
func (msg MsgRevoke) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgRevoke) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.VaultAddress}
}

// ------------------------------------------------------------------------------------------
// MsgUpdateClearingHeight - high level transaction of the coin module
// change vault clear time
type MsgUpdateClearingHeight struct {
	VaultAddress   types.AccAddress `json:"vault_address" yaml:"vault_address"`
	ClearingHeight uint64           `json:"clearing_height" yaml:"clearing_height"`
}

var _ types.Msg = MsgUpdateClearingHeight{}

// NewMsgUpdateClearingHeight - construct arbitrary multi-in, multi-out send msg.
func NewMsgUpdateClearingHeight(vaultAddr types.AccAddress, clearTime uint64) MsgUpdateClearingHeight {
	return MsgUpdateClearingHeight{VaultAddress: vaultAddr, ClearingHeight: clearTime}
}

// Type Implements Msg.
func (msg MsgUpdateClearingHeight) Type() string { return "updateclearingheight" }

// GetSignBytes Implements Msg.
func (msg MsgUpdateClearingHeight) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgUpdateClearingHeight) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.VaultAddress}
}

// ------------------------------------------------------------------------------------------
// MsgPublishMultiSigAccount msg for create multi sig account
type MsgPublishMultiSigAccount struct {
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	Pubkey      string           `json:"pubkey" yaml:"pubkey"`
}

var _ types.Msg = MsgPublishMultiSigAccount{}

// Type Implements Msg.
func (msg MsgPublishMultiSigAccount) Type() string { return "multisig" }

// GetSignBytes Implements Msg.
func (msg MsgPublishMultiSigAccount) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgPublishMultiSigAccount) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}

// ------------------------------------------------------------------------------------------
// MsgClearVaultAccount msg for clear vault account
type MsgClearVaultAccount struct {
	FromAddress  types.AccAddress   `json:"from_address" yaml:"from_address"`
	VaultAddress []types.AccAddress `json:"vault_address" yaml:"vault_address"`
}

var _ types.Msg = MsgClearVaultAccount{}

// NewMsgClearVaultAccount - for clear vault account
func NewMsgClearVaultAccount(fromAddr types.AccAddress, vaultAddr []types.AccAddress) MsgClearVaultAccount {
	return MsgClearVaultAccount{FromAddress: fromAddr, VaultAddress: vaultAddr}
}

// Type Implements Msg.
func (msg MsgClearVaultAccount) Type() string { return "clear" }

// GetSignBytes Implements Msg.
func (msg MsgClearVaultAccount) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgClearVaultAccount) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}

package types

import (
	"bytes"

	"github.com/gatechain/crypto"
	"github.com/gatechain/crypto/multisig"
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
)

// RouterKey is they name of the bank module
const RouterKey = ModuleName

var TagKeyRecipient = "recipient"

// ------------------------------------------------------------------------------------------
// MsgCreateVault msg for create vault account
type MsgCreateVault struct {
	FromAddress     types2.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress       types2.AccAddress `json:"to_address" yaml:"to_address"`
	SecurityAddress types2.AccAddress `json:"security_address" yaml:"security_address"`
	DelayHeight     uint64            `json:"delay_height" yaml:"delay_height"`
	ClearingHeight  uint64            `json:"clearing_height" yaml:"clearing_height"`
	Amount          types2.Coins      `json:"amount" yaml:"amount"`
	Pubkey          string            `json:"pubkey" yaml:"pubkey"`
}

var _ types2.Msg = MsgCreateVault{}

// NewMsgCreateVault - construct arbitrary multi-in, multi-out send msg.
func NewMsgCreateVault(fromAddr, toAddr types2.AccAddress, securityAddress types2.AccAddress, delayHeight uint64, clearingHeight uint64, amount types2.Coins, pubkey string) MsgCreateVault {
	return MsgCreateVault{FromAddress: fromAddr, ToAddress: toAddr, SecurityAddress: securityAddress, DelayHeight: delayHeight, ClearingHeight: clearingHeight, Amount: amount, Pubkey: pubkey}
}

// Route Implements Msg.
func (msg MsgCreateVault) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgCreateVault) Type() string { return "vault" }

// ValidateBasic Implements Msg.
func (msg MsgCreateVault) ValidateBasic() types2.Error {
	if msg.FromAddress.Empty() {
		return types2.ErrInvalidAddress("missing sender address")
	}
	if msg.ToAddress.Empty() {
		return types2.ErrInvalidAddress("missing recipient address")
	}
	if msg.SecurityAddress.Empty() {
		return types2.ErrInvalidAddress("missing security address")
	}
	if msg.SecurityAddress.Equals(msg.ToAddress) {
		return types2.ErrInvalidAddress("SecurityAddress must not equal InsuranceAddress")
	}

	if msg.ClearingHeight <= 0 {
		return types2.ErrInvalidAddress("error clearing height")
	}

	if msg.DelayHeight <= 0 {
		return types2.ErrInvalidAddress("error delay height")
	}

	if msg.Pubkey != "" {
		pubkey, e := types2.GetAccMultiSigPubKeyBech32(msg.Pubkey)
		if e != nil {
			return types2.ErrInvalidAddress("not a multisig address")
		}

		multiKey, ok := pubkey.(multisig.PubKeyMultisigThreshold)
		if !ok {
			return types2.ErrInvalidAddress("not a multisig address")
		}

		if !bytes.Equal(multiKey.Address(), msg.ToAddress.Bytes()) {
			return types2.ErrInvalidAddress("pubkey addr not match")
		}

		if countSubKeys(multiKey) == 1 {
			return types2.ErrInvalidAddress("not a multisig address")
		}
	}

	if !msg.Amount.IsValid() {
		return types2.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}
	if !msg.Amount.IsAllPositive() {
		return types2.ErrInsufficientCoins("send amount must be positive")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgCreateVault) GetSignBytes() []byte {
	return types2.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgCreateVault) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{msg.FromAddress}
}

// ------------------------------------------------------------------------------------------
// MsgRevocableSend - high level transaction of the coin module
// insurance account send coin to common account
type MsgRevocableSend struct {
	FromAddress types2.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types2.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      types2.Coins      `json:"amount" yaml:"amount"`
}

var _ types2.Msg = MsgRevocableSend{}

// NewMsgRevocableSend - construct arbitrary multi-in, multi-out send msg.
func NewMsgRevocableSend(fromAddr, toAddr types2.AccAddress, amount types2.Coins) MsgRevocableSend {
	return MsgRevocableSend{FromAddress: fromAddr, ToAddress: toAddr, Amount: amount}
}

// Route Implements Msg.
func (msg MsgRevocableSend) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgRevocableSend) Type() string { return "revocable" }

// ValidateBasic Implements Msg.
func (msg MsgRevocableSend) ValidateBasic() types2.Error {
	if msg.FromAddress.Empty() {
		return types2.ErrInvalidAddress("missing sender address")
	}
	if msg.ToAddress.Empty() {
		return types2.ErrInvalidAddress("missing recipient address")
	}
	if !msg.Amount.IsValid() {
		return types2.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}
	if !msg.Amount.IsAllPositive() {
		return types2.ErrInsufficientCoins("send amount must be positive")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgRevocableSend) GetSignBytes() []byte {
	return types2.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgRevocableSend) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{msg.FromAddress}
}

// ------------------------------------------------------------------------------------------
// MsgRevoke - high level transaction of the coin module
// revoke  coins to security account
type MsgRevoke struct {
	VaultAddress    types2.AccAddress `json:"vault_address" yaml:"vault_address"`
	SecurityAddress types2.AccAddress `json:"security_address" yaml:"security_address"`
	RevokeAddress   types2.AccAddress `json:"revoke_address" yaml:"revoke_address"`
	Height          int64             `json:"height" yaml:"height"`
	TxHash          string            `json:"tx_hash" yaml:"tx_hash"`
	MsgIndex        int64             `json:"msg_index" yaml:"msg_index"`
	Amount          types2.Coins      `json:"amount" yaml:"amount"`
}

var _ types2.Msg = MsgRevoke{}

// MsgRevoke - construct arbitrary multi-in, multi-out send msg.
func NewMsgRevoke(vaultAddr, securityAddress types2.AccAddress, revokeAddress types2.AccAddress, height int64, msgIndex int64, hash string, amount types2.Coins) MsgRevoke {
	return MsgRevoke{VaultAddress: vaultAddr, SecurityAddress: securityAddress,
		RevokeAddress: revokeAddress, Height: height, MsgIndex: msgIndex, TxHash: hash, Amount: amount}
}

// Route Implements Msg.
func (msg MsgRevoke) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgRevoke) Type() string { return "revoke" }

// ValidateBasic Implements Msg.
func (msg MsgRevoke) ValidateBasic() types2.Error {
	if msg.VaultAddress.Empty() {
		return types2.ErrInvalidAddress("missing vault address")
	}
	if msg.SecurityAddress.Empty() {
		return types2.ErrInvalidAddress("missing security address")
	}
	if msg.RevokeAddress.Empty() {
		return types2.ErrInvalidAddress("missing revoke address")
	}
	if msg.SecurityAddress.Equals(msg.VaultAddress) {
		return types2.ErrInvalidAddress("securityAddress must not equal vaultAddress")
	}
	if msg.RevokeAddress.Equals(msg.VaultAddress) {
		return types2.ErrInvalidAddress("revokeAddress must not equal vaultAddress")
	}
	if msg.TxHash == "" {
		return ErrInvalidTxHash(DefaultCodespace)
	}
	if msg.Height <= 0 {
		return ErrInvalidBlockHeight(DefaultCodespace)
	}
	if msg.MsgIndex < 0 {
		return ErrInvalidTxMsgIndex(DefaultCodespace)
	}
	if !msg.Amount.IsValid() {
		return types2.ErrInvalidCoins("send amount is invalid")
	}
	if !msg.Amount.IsAllPositive() {
		return types2.ErrInsufficientCoins("send amount must be positive")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgRevoke) GetSignBytes() []byte {
	return types2.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgRevoke) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{msg.VaultAddress}
}

// ------------------------------------------------------------------------------------------
// MsgUpdateClearingHeight - high level transaction of the coin module
// change vault clear time
type MsgUpdateClearingHeight struct {
	VaultAddress   types2.AccAddress `json:"vault_address" yaml:"vault_address"`
	ClearingHeight uint64            `json:"clearing_height" yaml:"clearing_height"`
}

var _ types2.Msg = MsgUpdateClearingHeight{}

// NewMsgUpdateClearingHeight - construct arbitrary multi-in, multi-out send msg.
func NewMsgUpdateClearingHeight(vaultAddr types2.AccAddress, clearTime uint64) MsgUpdateClearingHeight {
	return MsgUpdateClearingHeight{VaultAddress: vaultAddr, ClearingHeight: clearTime}
}

// Route Implements Msg.
func (msg MsgUpdateClearingHeight) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgUpdateClearingHeight) Type() string { return "updateclearingheight" }

// ValidateBasic Implements Msg.
func (msg MsgUpdateClearingHeight) ValidateBasic() types2.Error {
	if msg.VaultAddress.Empty() {
		return types2.ErrInvalidAddress("missing vault address")
	}

	if msg.ClearingHeight <= 0 {
		return ErrInvalidClearingHeight(DefaultCodespace, "clearing height can not be zero")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgUpdateClearingHeight) GetSignBytes() []byte {
	return types2.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgUpdateClearingHeight) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{msg.VaultAddress}
}

// ------------------------------------------------------------------------------------------
// MsgPublishMultiSigAccount msg for create multi sig account
type MsgPublishMultiSigAccount struct {
	FromAddress types2.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types2.AccAddress `json:"to_address" yaml:"to_address"`
	Pubkey      string            `json:"pubkey" yaml:"pubkey"`
}

var _ types2.Msg = MsgPublishMultiSigAccount{}

// NewMsgPublishMultiSigAccount - construct arbitrary multi-in, multi-out send msg.
func NewMsgPublishMultiSigAccount(fromAddr, toAddr types2.AccAddress, pubkey string) MsgPublishMultiSigAccount {
	return MsgPublishMultiSigAccount{FromAddress: fromAddr, ToAddress: toAddr, Pubkey: pubkey}
}

// Route Implements Msg.
func (msg MsgPublishMultiSigAccount) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgPublishMultiSigAccount) Type() string { return "multisig" }

// ValidateBasic Implements Msg.
func (msg MsgPublishMultiSigAccount) ValidateBasic() types2.Error {
	if msg.FromAddress.Empty() {
		return types2.ErrInvalidAddress("missing sender address")
	}
	if msg.ToAddress.Empty() {
		return types2.ErrInvalidAddress("missing recipient address")
	}

	pubkey, e := types2.GetAccMultiSigPubKeyBech32(msg.Pubkey)
	if e != nil {
		return types2.ErrInvalidAddress("not a multisig address")
	}

	multiKey, ok := pubkey.(multisig.PubKeyMultisigThreshold)
	if !ok {
		return types2.ErrInvalidAddress("not a multisig address")
	}

	if !bytes.Equal(multiKey.Address(), msg.ToAddress.Bytes()) {
		return types2.ErrInvalidAddress("pubkey addr not match")
	}

	if countSubKeys(multiKey) == 1 {
		return types2.ErrInvalidAddress("not a multisig address")
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgPublishMultiSigAccount) GetSignBytes() []byte {
	return types2.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgPublishMultiSigAccount) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{msg.FromAddress}
}

func countSubKeys(pub crypto.PubKey) int {
	v, ok := pub.(multisig.PubKeyMultisigThreshold)
	if !ok {
		return 1
	}

	numKeys := 0
	for _, subkey := range v.PubKeys {
		numKeys += countSubKeys(subkey)
	}

	return numKeys
}

// ------------------------------------------------------------------------------------------
// MsgClearVaultAccount msg for clear vault account
type MsgClearVaultAccount struct {
	FromAddress  types2.AccAddress   `json:"from_address" yaml:"from_address"`
	VaultAddress []types2.AccAddress `json:"vault_address" yaml:"vault_address"`
}

var _ types2.Msg = MsgClearVaultAccount{}

// NewMsgClearVaultAccount - for clear vault account
func NewMsgClearVaultAccount(fromAddr types2.AccAddress, vaultAddr []types2.AccAddress) MsgClearVaultAccount {
	return MsgClearVaultAccount{FromAddress: fromAddr, VaultAddress: vaultAddr}
}

// Route Implements Msg.
func (msg MsgClearVaultAccount) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgClearVaultAccount) Type() string { return "clear" }

// ValidateBasic Implements Msg.
func (msg MsgClearVaultAccount) ValidateBasic() types2.Error {
	if msg.FromAddress.Empty() {
		return types2.ErrInvalidAddress("missing sender address")
	}
	if len(msg.VaultAddress) == 0 {
		return types2.ErrInvalidAddress("missing vault address")
	}

	for _, vault := range msg.VaultAddress {
		if vault.Empty() {
			return types2.ErrInvalidAddress("vault address can not be empty")
		}
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgClearVaultAccount) GetSignBytes() []byte {
	return types2.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgClearVaultAccount) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{msg.FromAddress}
}

package types

import (
	"bytes"
	"encoding/json"

	"github.com/gatechain/crypto"

	types "github.com/gatechain/gatechainsdk/gatechain/types"
)

// ensure Msg interface compliance at compile time
var (
	_ types.Msg = &MsgCreateValidator{}
	_ types.Msg = &MsgDelegate{}
	_ types.Msg = &MsgUndelegate{}

	_ types.Msg = &MsgBeginRedelegate{}
	_ types.Msg = &MsgUndelegateByRetrievalAccount{}
)

//______________________________________________________________________

// MsgCreateValidator - struct for bonding transactions
type MsgCreateValidator struct {
	Description      Description      `json:"description" yaml:"description"`
	Commission       CommissionRates  `json:"commission" yaml:"commission"`
	DelegatorAddress types.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	PubKey           crypto.PubKey    `json:"pubkey" yaml:"pubkey"`
	Extra            []byte           `json:"extra" yaml:"extra"`
}

type msgCreateValidatorJSON struct {
	Description      Description      `json:"description" yaml:"description"`
	Commission       CommissionRates  `json:"commission" yaml:"commission"`
	DelegatorAddress types.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	PubKey           string           `json:"pubkey" yaml:"pubkey"`
	Extra            []byte           `json:"extra" yaml:"extra"`
}

// nolint
// func (msg MsgCreateValidator) Route() string { return RouterKey }
func (msg MsgCreateValidator) Type() string { return "create_validator" }

// Return address(es) that must sign over msg.GetSignBytes()
func (msg MsgCreateValidator) GetSigners() []types.AccAddress {
	// delegator is first signer so delegator pays fees
	addrs := []types.AccAddress{msg.DelegatorAddress}

	if !bytes.Equal(msg.DelegatorAddress.Bytes(), msg.ValidatorAddress.Bytes()) {
		// if validator addr is not same as delegator addr, validator must sign
		// msg as well
		addrs = append(addrs, types.AccAddress(msg.ValidatorAddress))
	}
	return addrs
}

// MarshalJSON implements the json.Marshaler interface to provide custom JSON
// serialization of the MsgCreateValidator type.
func (msg MsgCreateValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgCreateValidatorJSON{
		Description:      msg.Description,
		Commission:       msg.Commission,
		DelegatorAddress: msg.DelegatorAddress,
		ValidatorAddress: msg.ValidatorAddress,
		PubKey:           types.MustBech32ifyConsPub(msg.PubKey),
		Extra:            msg.Extra,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface to provide custom
// JSON deserialization of the MsgCreateValidator type.
func (msg *MsgCreateValidator) UnmarshalJSON(bz []byte) error {
	var msgCreateValJSON msgCreateValidatorJSON
	if err := json.Unmarshal(bz, &msgCreateValJSON); err != nil {
		return err
	}

	msg.Description = msgCreateValJSON.Description
	msg.Commission = msgCreateValJSON.Commission
	msg.DelegatorAddress = msgCreateValJSON.DelegatorAddress
	msg.ValidatorAddress = msgCreateValJSON.ValidatorAddress
	msg.Extra = msgCreateValJSON.Extra
	var err error
	msg.PubKey, err = types.GetConsPubKeyBech32(msgCreateValJSON.PubKey)
	if err != nil {
		return err
	}
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgCreateValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types.MustSortJSON(bz)
}

// MsgDelegate - struct for bonding transactions
type MsgDelegate struct {
	DelegatorAddress types.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	Amount           types.Coin       `json:"amount" yaml:"amount"`
}

func NewMsgDelegate(delAddr types.AccAddress, valAddr types.ValAddress, amount types.Coin) MsgDelegate {
	return MsgDelegate{
		DelegatorAddress: delAddr,
		ValidatorAddress: valAddr,
		Amount:           amount,
	}
}

// nolint
// func (msg MsgDelegate) Route() string { return RouterKey }
func (msg MsgDelegate) Type() string { return "delegation" }
func (msg MsgDelegate) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.DelegatorAddress}
}

// get the bytes for the message signer to sign on
func (msg MsgDelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types.MustSortJSON(bz)
}

//______________________________________________________________________

type MsgUndelegateByRetrievalAccount struct {
	SecurityAddress  types.AccAddress   `json:"security_address" yaml:"security_address"`
	DelegatorAddress []types.AccAddress `json:"delegator_address" yaml:"delegator_address"`
}

func NewMsgUndelegateByRetrievalAccount(secAddr types.AccAddress, delAddr []types.AccAddress) MsgUndelegateByRetrievalAccount {

	return MsgUndelegateByRetrievalAccount{
		SecurityAddress:  secAddr,
		DelegatorAddress: delAddr,
	}
}

// nolint
// func (msg MsgUndelegateByRetrievalAccount) Route() string { return RouterKey }
func (msg MsgUndelegateByRetrievalAccount) Type() string { return "undelegateBySecurityAddress" }

func (msg MsgUndelegateByRetrievalAccount) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.SecurityAddress}
}

func (msg MsgUndelegateByRetrievalAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types.MustSortJSON(bz)
}

// MsgDelegate - struct for bonding transactions
type MsgBeginRedelegate struct {
	DelegatorAddress    types.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorSrcAddress types.ValAddress `json:"con-account_src_address" yaml:"con-account_src_address"`
	ValidatorDstAddress types.ValAddress `json:"con-account_dst_address" yaml:"con-account_dst_address"`
	Amount              types.Coin       `json:"amount" yaml:"amount"`
}

func NewMsgBeginRedelegate(delAddr types.AccAddress, valSrcAddr,
	valDstAddr types.ValAddress, amount types.Coin) MsgBeginRedelegate {

	return MsgBeginRedelegate{
		DelegatorAddress:    delAddr,
		ValidatorSrcAddress: valSrcAddr,
		ValidatorDstAddress: valDstAddr,
		Amount:              amount,
	}
}

// nolint
// func (msg MsgBeginRedelegate) Route() string { return RouterKey }
func (msg MsgBeginRedelegate) Type() string { return "redelegation" }
func (msg MsgBeginRedelegate) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.DelegatorAddress}
}

// get the bytes for the message signer to sign on
func (msg MsgBeginRedelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types.MustSortJSON(bz)
}

// MsgUndelegate - struct for unbonding transactions
type MsgUndelegate struct {
	DelegatorAddress types.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	Amount           types.Coin       `json:"amount" yaml:"amount"`
}

func NewMsgUndelegate(delAddr types.AccAddress, valAddr types.ValAddress, amount types.Coin) MsgUndelegate {
	return MsgUndelegate{
		DelegatorAddress: delAddr,
		ValidatorAddress: valAddr,
		Amount:           amount,
	}
}

// nolint
// func (msg MsgUndelegate) Route() string { return RouterKey }
func (msg MsgUndelegate) Type() string { return "undelegation" }
func (msg MsgUndelegate) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.DelegatorAddress}
}

// get the bytes for the message signer to sign on
func (msg MsgUndelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types.MustSortJSON(bz)
}

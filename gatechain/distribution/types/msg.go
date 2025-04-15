// nolint
package types

import (
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
)

// Verify interface at compile time
var _, _, _ types2.Msg = &MsgSetWithdrawAddress{}, &MsgWithdrawDelegatorReward{}, &MsgWithdrawValidatorCommission{}

// msg struct for changing the withdraw address for a delegator (or validator self-delegation)
type MsgSetWithdrawAddress struct {
	DelegatorAddress types2.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	WithdrawAddress  types2.AccAddress `json:"withdraw_address" yaml:"withdraw_address"`
}

func NewMsgSetWithdrawAddress(delAddr, withdrawAddr types2.AccAddress) MsgSetWithdrawAddress {
	return MsgSetWithdrawAddress{
		DelegatorAddress: delAddr,
		WithdrawAddress:  withdrawAddr,
	}
}

func (msg MsgSetWithdrawAddress) Route() string { return ModuleName }
func (msg MsgSetWithdrawAddress) Type() string  { return "set_withdraw_address" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgSetWithdrawAddress) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{types2.AccAddress(msg.DelegatorAddress)}
}

// get the bytes for the message signer to sign on
func (msg MsgSetWithdrawAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgSetWithdrawAddress) ValidateBasic() types2.Error {
	if msg.DelegatorAddress.Empty() {
		return ErrNilDelegatorAddr(DefaultCodespace)
	}
	if msg.WithdrawAddress.Empty() {
		return ErrNilWithdrawAddr(DefaultCodespace)
	}
	return nil
}

// msg struct for delegation withdraw from a single validator
type MsgWithdrawDelegatorReward struct {
	DelegatorAddress types2.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types2.ValAddress `json:"con-account_address" yaml:"con-account_address"`
}

func NewMsgWithdrawDelegatorReward(delAddr types2.AccAddress, valAddr types2.ValAddress) MsgWithdrawDelegatorReward {
	return MsgWithdrawDelegatorReward{
		DelegatorAddress: delAddr,
		ValidatorAddress: valAddr,
	}
}

func (msg MsgWithdrawDelegatorReward) Route() string { return ModuleName }
func (msg MsgWithdrawDelegatorReward) Type() string  { return "withdraw_delegator_reward" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgWithdrawDelegatorReward) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{types2.AccAddress(msg.DelegatorAddress)}
}

// get the bytes for the message signer to sign on
func (msg MsgWithdrawDelegatorReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgWithdrawDelegatorReward) ValidateBasic() types2.Error {
	if msg.DelegatorAddress.Empty() {
		return ErrNilDelegatorAddr(DefaultCodespace)
	}
	if msg.ValidatorAddress.Empty() {
		return ErrNilValidatorAddr(DefaultCodespace)
	}
	return nil
}

// msg struct for validator withdraw
type MsgWithdrawValidatorCommission struct {
	ValidatorAddress types2.ValAddress `json:"con-account_address" yaml:"con-account_address"`
}

func NewMsgWithdrawValidatorCommission(valAddr types2.ValAddress) MsgWithdrawValidatorCommission {
	return MsgWithdrawValidatorCommission{
		ValidatorAddress: valAddr,
	}
}

func (msg MsgWithdrawValidatorCommission) Route() string { return ModuleName }
func (msg MsgWithdrawValidatorCommission) Type() string  { return "withdraw_con-account_commission" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgWithdrawValidatorCommission) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{types2.AccAddress(msg.ValidatorAddress.Bytes())}
}

// get the bytes for the message signer to sign on
func (msg MsgWithdrawValidatorCommission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgWithdrawValidatorCommission) ValidateBasic() types2.Error {
	if msg.ValidatorAddress.Empty() {
		return ErrNilValidatorAddr(DefaultCodespace)
	}
	return nil
}

type MsgRewardReinvestment struct {
	DelegatorAddress types2.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types2.ValAddress `json:"con-account_address" yaml:"con-account_address"`
}

func NewMsgRewardReinvestment(delAddr types2.AccAddress, valAddr types2.ValAddress) MsgRewardReinvestment {
	return MsgRewardReinvestment{
		DelegatorAddress: delAddr,
		ValidatorAddress: valAddr,
	}
}

func (msg MsgRewardReinvestment) Route() string { return ModuleName }
func (msg MsgRewardReinvestment) Type() string  { return "delegate_reward_reinvestment" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgRewardReinvestment) GetSigners() []types2.AccAddress {
	return []types2.AccAddress{types2.AccAddress(msg.DelegatorAddress)}
}

// get the bytes for the message signer to sign on
func (msg MsgRewardReinvestment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types2.MustSortJSON(bz)
}

// quick validity check
func (msg MsgRewardReinvestment) ValidateBasic() types2.Error {
	if msg.DelegatorAddress.Empty() {
		return ErrNilDelegatorAddr(DefaultCodespace)
	}
	if msg.ValidatorAddress.Empty() {
		return ErrNilValidatorAddr(DefaultCodespace)
	}
	return nil
}

// nolint
package types

import (
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// Verify interface at compile time
var _, _, _ types.Msg = &MsgSetWithdrawAddress{}, &MsgWithdrawDelegatorReward{}, &MsgWithdrawValidatorCommission{}

// msg struct for changing the withdraw address for a delegator (or validator self-delegation)
type MsgSetWithdrawAddress struct {
	DelegatorAddress types.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	WithdrawAddress  types.AccAddress `json:"withdraw_address" yaml:"withdraw_address"`
}

func NewMsgSetWithdrawAddress(delAddr, withdrawAddr types.AccAddress) MsgSetWithdrawAddress {
	return MsgSetWithdrawAddress{
		DelegatorAddress: delAddr,
		WithdrawAddress:  withdrawAddr,
	}
}

func (msg MsgSetWithdrawAddress) Type() string { return "set_withdraw_address" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgSetWithdrawAddress) GetSigners() []types.AccAddress {
	return []types.AccAddress{types.AccAddress(msg.DelegatorAddress)}
}

// get the bytes for the message signer to sign on
func (msg MsgSetWithdrawAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types.MustSortJSON(bz)
}

// msg struct for delegation withdraw from a single validator
type MsgWithdrawDelegatorReward struct {
	DelegatorAddress types.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types.ValAddress `json:"con-account_address" yaml:"con-account_address"`
}

func NewMsgWithdrawDelegatorReward(delAddr types.AccAddress, valAddr types.ValAddress) MsgWithdrawDelegatorReward {
	return MsgWithdrawDelegatorReward{
		DelegatorAddress: delAddr,
		ValidatorAddress: valAddr,
	}
}

// func (msg MsgWithdrawDelegatorReward) Route() string { return ModuleName }
func (msg MsgWithdrawDelegatorReward) Type() string { return "withdraw_delegator_reward" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgWithdrawDelegatorReward) GetSigners() []types.AccAddress {
	return []types.AccAddress{types.AccAddress(msg.DelegatorAddress)}
}

// get the bytes for the message signer to sign on
func (msg MsgWithdrawDelegatorReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types.MustSortJSON(bz)
}

// msg struct for validator withdraw
type MsgWithdrawValidatorCommission struct {
	ValidatorAddress types.ValAddress `json:"con-account_address" yaml:"con-account_address"`
}

func NewMsgWithdrawValidatorCommission(valAddr types.ValAddress) MsgWithdrawValidatorCommission {
	return MsgWithdrawValidatorCommission{
		ValidatorAddress: valAddr,
	}
}

// func (msg MsgWithdrawValidatorCommission) Route() string { return ModuleName }
func (msg MsgWithdrawValidatorCommission) Type() string { return "withdraw_con-account_commission" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgWithdrawValidatorCommission) GetSigners() []types.AccAddress {
	return []types.AccAddress{types.AccAddress(msg.ValidatorAddress.Bytes())}
}

// get the bytes for the message signer to sign on
func (msg MsgWithdrawValidatorCommission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types.MustSortJSON(bz)
}

type MsgRewardReinvestment struct {
	DelegatorAddress types.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types.ValAddress `json:"con-account_address" yaml:"con-account_address"`
}

func NewMsgRewardReinvestment(delAddr types.AccAddress, valAddr types.ValAddress) MsgRewardReinvestment {
	return MsgRewardReinvestment{
		DelegatorAddress: delAddr,
		ValidatorAddress: valAddr,
	}
}

// func (msg MsgRewardReinvestment) Route() string { return ModuleName }
func (msg MsgRewardReinvestment) Type() string { return "delegate_reward_reinvestment" }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgRewardReinvestment) GetSigners() []types.AccAddress {
	return []types.AccAddress{types.AccAddress(msg.DelegatorAddress)}
}

// get the bytes for the message signer to sign on
func (msg MsgRewardReinvestment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return types.MustSortJSON(bz)
}

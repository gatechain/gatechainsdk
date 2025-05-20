package bank

import (
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// MsgSend - high level transaction of the coin module
type MsgSend struct {
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      types.Coins      `json:"amount" yaml:"amount"`
}

var _ types.Msg = MsgSend{}

// NewMsgSend - construct arbitrary multi-in, multi-out send msg.
func NewMsgSend(fromAddr, toAddr types.AccAddress, amount types.Coins) MsgSend {
	return MsgSend{FromAddress: fromAddr, ToAddress: toAddr, Amount: amount}
}

// Type Implements Msg.
func (msg MsgSend) Type() string { return "send" }

// GetSignBytes Implements Msg.
func (msg MsgSend) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgSend) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}

package types

import (
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// Bank errors reserve 100 ~ 199.
const (
	DefaultCodespace          sdk.CodespaceType = ModuleName
	CodeInvalidClearingHeight sdk.CodeType      = 204
	CodeInvalidTxHash         sdk.CodeType      = 205
	CodeInvalidBlockHeight    sdk.CodeType      = 206
	CodeInvalidTxMsgIndex     sdk.CodeType      = 207
)

func ErrInvalidClearingHeight(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidClearingHeight, "clearing height invalid: %s", msg)
}
func ErrInvalidTxHash(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidTxHash, "tx hash invalid")
}
func ErrInvalidBlockHeight(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidBlockHeight, "block height invalid")
}
func ErrInvalidTxMsgIndex(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidTxMsgIndex, "tx index invalid")
}

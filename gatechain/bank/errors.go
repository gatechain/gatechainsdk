package bank

import (
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// Bank errors reserve 100 ~ 199.
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeSendDisabled         sdk.CodeType = 101
	CodeInvalidInputsOutputs sdk.CodeType = 102
)

// ErrNoInputs is an error
func ErrNoInputs(codespace sdk.CodespaceType) sdk.Error {
	//return sdk.NewError(codespace, CodeInvalidInputsOutputs, "no inputs to send transaction")
	return nil
}

// ErrNoOutputs is an error
func ErrNoOutputs(codespace sdk.CodespaceType) sdk.Error {
	//return sdk.NewError(codespace, CodeInvalidInputsOutputs, "no outputs to send transaction")
	return nil

}

// ErrInputOutputMismatch is an error
func ErrInputOutputMismatch(codespace sdk.CodespaceType) sdk.Error {
	//return sdk.NewError(codespace, CodeInvalidInputsOutputs, "sum inputs != sum outputs")
	return nil

}

// ErrSendDisabled is an error
func ErrSendDisabled(codespace sdk.CodespaceType) sdk.Error {
	//return sdk.NewError(codespace, CodeSendDisabled, "send transactions are currently disabled")
	return nil

}

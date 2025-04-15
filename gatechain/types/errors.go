package types

import "fmt"

type Error interface {
	//// Implements cmn.Error
	Error() string
}

func (e *sdkError) Error() string {
	return fmt.Sprintf(`ERROR:
Codespace: %s
Code: %d
`, e.codespace, e.code)
}

const (
	// Base error codes
	CodeOK                  CodeType = 0
	CodeInternal            CodeType = 1
	CodeTxDecode            CodeType = 2
	CodeInvalidSequence     CodeType = 3
	CodeUnauthorized        CodeType = 4
	CodeInsufficientFunds   CodeType = 5
	CodeUnknownRequest      CodeType = 6
	CodeInvalidAddress      CodeType = 7
	CodeInvalidPubKey       CodeType = 8
	CodeUnknownAddress      CodeType = 9
	CodeInsufficientCoins   CodeType = 10
	CodeInvalidCoins        CodeType = 11
	CodeOutOfGas            CodeType = 12
	CodeMemoTooLarge        CodeType = 13
	CodeInsufficientFee     CodeType = 14
	CodeTooManySignatures   CodeType = 15
	CodeGasOverflow         CodeType = 16
	CodeNoSignatures        CodeType = 17
	CodeInvalidExpireHeight CodeType = 18

	// CodespaceRoot is a codespace for error codes in this file only.
	// Notice that 0 is an "unset" codespace, which can be overridden with
	// Error.WithDefaultCodespace().
	CodespaceUndefined CodespaceType = ""
	CodespaceRoot      CodespaceType = "framework"
)

// CodespaceType - codespace identifier
type CodespaceType string

// CodeType - ABCI code identifier within codespace
type CodeType uint32

// IsOK - is everything okay?
func (code CodeType) IsOK() bool {
	return code == CodeOK
}

type sdkError struct {
	codespace CodespaceType
	code      CodeType
	//cmnError
}

// nolint
func ErrInternal(msg string) Error {
	return newErrorWithRootCodespace(CodeInternal, msg)
}
func ErrTxDecode(msg string) Error {
	return newErrorWithRootCodespace(CodeTxDecode, msg)
}
func ErrInvalidSequence(msg string) Error {
	return newErrorWithRootCodespace(CodeInvalidSequence, msg)
}
func ErrUnauthorized(msg string) Error {
	return newErrorWithRootCodespace(CodeUnauthorized, msg)
}
func ErrInsufficientFunds(msg string) Error {
	return newErrorWithRootCodespace(CodeInsufficientFunds, msg)
}
func ErrUnknownRequest(msg string) Error {
	return newErrorWithRootCodespace(CodeUnknownRequest, msg)
}
func ErrInvalidAddress(msg string) Error {
	return newErrorWithRootCodespace(CodeInvalidAddress, msg)
}
func ErrInvalidExpireHeight(msg string) Error {
	return newErrorWithRootCodespace(CodeInvalidExpireHeight, msg)
}
func ErrUnknownAddress(msg string) Error {
	return newErrorWithRootCodespace(CodeUnknownAddress, msg)
}
func ErrInvalidPubKey(msg string) Error {
	return newErrorWithRootCodespace(CodeInvalidPubKey, msg)
}
func ErrInsufficientCoins(msg string) Error {
	return newErrorWithRootCodespace(CodeInsufficientCoins, msg)
}
func ErrInvalidCoins(msg string) Error {
	return newErrorWithRootCodespace(CodeInvalidCoins, msg)
}
func ErrOutOfGas(msg string) Error {
	return newErrorWithRootCodespace(CodeOutOfGas, msg)
}
func ErrMemoTooLarge(msg string) Error {
	return newErrorWithRootCodespace(CodeMemoTooLarge, msg)
}
func ErrInsufficientFee(msg string) Error {
	return newErrorWithRootCodespace(CodeInsufficientFee, msg)
}
func ErrTooManySignatures(msg string) Error {
	return newErrorWithRootCodespace(CodeTooManySignatures, msg)
}
func ErrNoSignatures(msg string) Error {
	return newErrorWithRootCodespace(CodeNoSignatures, msg)
}
func ErrGasOverflow(msg string) Error {
	return newErrorWithRootCodespace(CodeGasOverflow, msg)
}

func newErrorWithRootCodespace(code CodeType, format string, args ...interface{}) *sdkError {
	return newError(CodespaceRoot, code, format, args...)
}

func newError(codespace CodespaceType, code CodeType, format string, args ...interface{}) *sdkError {
	if format == "" {
		format = CodeToDefaultMsg(code)
	}
	return &sdkError{
		codespace: codespace,
		code:      code,
		//cmnError:  cmn.NewError(format, args...),
	}
}

func unknownCodeMsg(code CodeType) string {
	return fmt.Sprintf("unknown code %d", code)
}

// NOTE: Don't stringer this, we'll put better messages in later.
func CodeToDefaultMsg(code CodeType) string {
	switch code {
	case CodeInternal:
		return "internal error"
	case CodeTxDecode:
		return "tx parse error"
	case CodeInvalidSequence:
		return "invalid sequence"
	case CodeUnauthorized:
		return "unauthorized"
	case CodeInsufficientFunds:
		return "insufficient funds"
	case CodeUnknownRequest:
		return "unknown request"
	case CodeInvalidAddress:
		return "invalid address"
	case CodeInvalidPubKey:
		return "invalid pubkey"
	case CodeUnknownAddress:
		return "unknown address"
	case CodeInsufficientCoins:
		return "insufficient coins"
	case CodeInvalidCoins:
		return "invalid coins"
	case CodeOutOfGas:
		return "out of gas"
	case CodeMemoTooLarge:
		return "Memo too large"
	case CodeInsufficientFee:
		return "insufficient fee"
	case CodeTooManySignatures:
		return "maximum numer of signatures exceeded"
	case CodeNoSignatures:
		return "no signatures supplied"
	default:
		return unknownCodeMsg(code)
	}
}

// NewError - create an error.
func NewError(codespace CodespaceType, code CodeType, format string, args ...interface{}) Error {
	return newError(codespace, code, format, args...)
}

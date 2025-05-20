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
	CodeOK CodeType = 0

	CodeUnknownRequest CodeType = 6

	// CodespaceRoot is a codespace for error codes in this file only.
	// Notice that 0 is an "unset" codespace, which can be overridden with
	// Error.WithDefaultCodespace().
	//CodespaceUndefined CodespaceType = ""
	CodespaceRoot CodespaceType = "framework"
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
}

func ErrUnknownRequest(msg string) Error {
	return newErrorWithRootCodespace(CodeUnknownRequest, msg)
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
	case CodeUnknownRequest:
		return "unknown request"
	default:
		return unknownCodeMsg(code)
	}
}

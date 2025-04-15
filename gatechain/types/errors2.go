package types

import (
	"fmt"
)

const (
	codeKeyNotFound   = 1
	codeWrongPassword = 2
)

type keybaseError interface {
	error
	Code() int
}

type errKeyNotFound struct {
	code int
	name string
}

func (e errKeyNotFound) Code() int {
	return e.code
}

func (e errKeyNotFound) Error() string {
	return fmt.Sprintf("Key %s not found", e.name)
}

// NewErrKeyNotFound returns a standardized error reflecting that the specified key doesn't exist
func NewErrKeyNotFound(name string) error {
	return errKeyNotFound{
		code: codeKeyNotFound,
		name: name,
	}
}

type errWrongPassword struct {
	code int
}

func (e errWrongPassword) Code() int {
	return e.code
}

func (e errWrongPassword) Error() string {
	return "invalid account password"
}

// NewErrWrongPassword returns a standardized error reflecting that the specified password is wrong
func NewErrWrongPassword() error {
	return errWrongPassword{
		code: codeWrongPassword,
	}
}

package types

import (
	"encoding/json"
	"fmt"
	"math/big"
)

const maxBitLen = 255

func newIntegerFromString(s string) (*big.Int, bool) {
	return new(big.Int).SetString(s, 0)
}

func lt(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == -1 }

func add(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Add(i, i2) }

func mul(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Mul(i, i2) }

// MarshalAmino for custom encoding scheme
func marshalAmino(i *big.Int) (string, error) {
	bz, err := i.MarshalText()
	return string(bz), err
}

func unmarshalText(i *big.Int, text string) error {
	if err := i.UnmarshalText([]byte(text)); err != nil {
		return err
	}

	if i.BitLen() > maxBitLen {
		return fmt.Errorf("integer out of range: %s", text)
	}

	return nil
}

// UnmarshalAmino for custom decoding scheme
func unmarshalAmino(i *big.Int, text string) (err error) {
	return unmarshalText(i, text)
}

// MarshalJSON for custom encoding scheme
// Must be encoded as a string for JSON precision
func marshalJSON(i *big.Int) ([]byte, error) {
	text, err := i.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

// UnmarshalJSON for custom decoding scheme
// Must be encoded as a string for JSON precision
func unmarshalJSON(i *big.Int, bz []byte) error {
	var text string
	err := json.Unmarshal(bz, &text)
	if err != nil {
		return err
	}

	return unmarshalText(i, text)
}

// Int wraps integer with 256 bit range bound
// Checks overflow, underflow and division by zero
// Exists in range from -(2^maxBitLen-1) to 2^maxBitLen-1
type Int struct {
	i *big.Int
}

// BigInt converts Int to big.Int
func (i Int) BigInt() *big.Int {
	return new(big.Int).Set(i.i)
}

// NewIntFromBigInt constructs Int from big.Int
func NewIntFromBigInt(i *big.Int) Int {
	if i.BitLen() > maxBitLen {
		panic("NewIntFromBigInt() out of bound")
	}
	return Int{i}
}

// NewIntFromString constructs Int from string
func NewIntFromString(s string) (res Int, ok bool) {
	i, ok := newIntegerFromString(s)
	if !ok {
		return
	}
	// Check overflow
	if i.BitLen() > maxBitLen {
		ok = false
		return
	}
	return Int{i}, true
}

// ZeroInt returns Int value with zero
func ZeroInt() Int { return Int{big.NewInt(0)} }

// IsZero returns true if Int is zero
func (i Int) IsZero() bool {
	return i.i.Sign() == 0
}

// Sign returns sign of Int
func (i Int) Sign() int {
	return i.i.Sign()
}

// LT returns true if first Int is lesser than second
func (i Int) LT(i2 Int) bool {
	return lt(i.i, i2.i)
}

// Add adds Int from another
func (i Int) Add(i2 Int) (res Int) {
	res = Int{add(i.i, i2.i)}
	// Check overflow
	if res.i.BitLen() > maxBitLen {
		panic("Int overflow")
	}
	return
}

// Mul multiples two Ints
func (i Int) Mul(i2 Int) (res Int) {
	// Check overflow
	if i.i.BitLen()+i2.i.BitLen()-1 > maxBitLen {
		panic("Int overflow")
	}
	res = Int{mul(i.i, i2.i)}
	// Check overflow if sign of both are same
	if res.i.BitLen() > maxBitLen {
		panic("Int overflow")
	}
	return
}

// Human readable string
func (i Int) String() string {
	return i.i.String()
}

// MarshalJSON defines custom encoding scheme
func (i Int) MarshalJSON() ([]byte, error) {
	if i.i == nil { // Necessary since default Uint initialization has i.i as nil
		i.i = new(big.Int)
	}
	return marshalJSON(i.i)
}

// UnmarshalJSON defines custom decoding scheme
func (i *Int) UnmarshalJSON(bz []byte) error {
	if i.i == nil { // Necessary since default Int initialization has i.i as nil
		i.i = new(big.Int)
	}
	return unmarshalJSON(i.i, bz)
}

// MarshalAmino defines custom encoding scheme
func (i Int) MarshalAmino() (string, error) {
	if i.i == nil { // Necessary since default Uint initialization has i.i as nil
		i.i = new(big.Int)
	}
	return marshalAmino(i.i)
}

// UnmarshalAmino defines custom decoding scheme
func (i *Int) UnmarshalAmino(text string) error {
	if i.i == nil { // Necessary since default Int initialization has i.i as nil
		i.i = new(big.Int)
	}
	return unmarshalAmino(i.i, text)
}

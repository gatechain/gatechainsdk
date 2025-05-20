package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// NOTE: never use new(Dec) or else we will panic unmarshalling into the
// nil embedded big.Int
type Dec struct {
	*big.Int `json:"int"`
}

// number of decimal places
const (
	Precision = 18

	// bytes required to represent the above precision
	// Ceiling[Log2[999 999 999 999 999 999]]
	DecimalPrecisionBits = 60
)

var (
	precisionReuse       = new(big.Int).Exp(big.NewInt(10), big.NewInt(Precision), nil)
	fivePrecision        = new(big.Int).Quo(precisionReuse, big.NewInt(2))
	precisionMultipliers []*big.Int
	zeroInt              = big.NewInt(0)
	oneInt               = big.NewInt(1)
	tenInt               = big.NewInt(10)
)

// Set precision multipliers
func init() {
	precisionMultipliers = make([]*big.Int, Precision+1)
	for i := 0; i <= Precision; i++ {
		precisionMultipliers[i] = calcPrecisionMultiplier(int64(i))
	}
}

// calculate the precision multiplier
func calcPrecisionMultiplier(prec int64) *big.Int {
	if prec > Precision {
		panic(fmt.Sprintf("too much precision, maximum %v, provided %v", Precision, prec))
	}
	zerosToAdd := Precision - prec
	multiplier := new(big.Int).Exp(tenInt, big.NewInt(zerosToAdd), nil)
	return multiplier
}

// get the precision multiplier, do not mutate result
func precisionMultiplier(prec int64) *big.Int {
	if prec > Precision {
		panic(fmt.Sprintf("too much precision, maximum %v, provided %v", Precision, prec))
	}
	return precisionMultipliers[prec]
}

////______________________________________________________________________________________________

// create a new Dec from integer assuming whole number
func NewDec(i int64) Dec {
	return NewDecWithPrec(i, 0)
}

// create a new Dec from integer with decimal place at prec
// CONTRACT: prec <= Precision
func NewDecWithPrec(i, prec int64) Dec {
	return Dec{
		new(big.Int).Mul(big.NewInt(i), precisionMultiplier(prec)),
	}
}

// create a new Dec from big integer assuming whole numbers
// CONTRACT: prec <= Precision
func NewDecFromBigInt(i *big.Int) Dec {
	return NewDecFromBigIntWithPrec(i, 0)
}

// create a new Dec from big integer assuming whole numbers
// CONTRACT: prec <= Precision
func NewDecFromBigIntWithPrec(i *big.Int, prec int64) Dec {
	return Dec{
		new(big.Int).Mul(i, precisionMultiplier(prec)),
	}
}

// create a decimal from an input decimal string.
// valid must come in the form:
//
//	(-) whole integers (.) decimal integers
//
// examples of acceptable input include:
//
//	-123.456
//	456.7890
//	345
//	-456789
//
// NOTE - An error will return if more decimal places
// are provided in the string than the constant Precision.
//
// CONTRACT - This function does not mutate the input str.
func NewDecFromStr(str string) (d Dec, err error) {
	if len(str) == 0 {
		return d, ErrUnknownRequest("decimal string is empty")
	}

	// first extract any negative symbol
	neg := false
	if str[0] == '-' {
		neg = true
		str = str[1:]
	}

	if len(str) == 0 {
		return d, ErrUnknownRequest("decimal string is empty")
	}

	strs := strings.Split(str, ".")
	lenDecs := 0
	combinedStr := strs[0]

	if len(strs) == 2 { // has a decimal place
		lenDecs = len(strs[1])
		if lenDecs == 0 || len(combinedStr) == 0 {
			return d, ErrUnknownRequest("bad decimal length")
		}
		combinedStr = combinedStr + strs[1]

	} else if len(strs) > 2 {
		return d, ErrUnknownRequest("too many periods to be a decimal string")
	}

	if lenDecs > Precision {
		return d, ErrUnknownRequest(
			fmt.Sprintf("too much precision, maximum %v, len decimal %v", Precision, lenDecs))
	}

	// add some extra zero's to correct to the Precision factor
	zerosToAdd := Precision - lenDecs
	zeros := fmt.Sprintf(`%0`+strconv.Itoa(zerosToAdd)+`s`, "")
	combinedStr = combinedStr + zeros

	combined, ok := new(big.Int).SetString(combinedStr, 10) // base 10
	if !ok {
		return d, ErrUnknownRequest(fmt.Sprintf("bad string to integer conversion, combinedStr: %v", combinedStr))
	}
	if neg {
		combined = new(big.Int).Neg(combined)
	}
	return Dec{combined}, nil
}

// ______________________________________________________________________________________________
// nolint
// func (d Dec) IsNil() bool       { return d.Int == nil }                 // is decimal nil
func (d Dec) IsZero() bool     { return (d.Int).Sign() == 0 }          // is equal to zero
func (d Dec) IsNegative() bool { return (d.Int).Sign() == -1 }         // is negative
func (d Dec) Neg() Dec         { return Dec{new(big.Int).Neg(d.Int)} } // reverse the decimal sign

// multiplication
func (d Dec) Mul(d2 Dec) Dec {
	mul := new(big.Int).Mul(d.Int, d2.Int)
	chopped := chopPrecisionAndRound(mul)

	if chopped.BitLen() > 255+DecimalPrecisionBits {
		panic("Int overflow")
	}
	return Dec{chopped}
}

func (d Dec) String() string {
	if d.Int == nil {
		return d.Int.String()
	}

	isNeg := d.IsNegative()
	if d.IsNegative() {
		d = d.Neg()
	}

	bzInt, err := d.Int.MarshalText()
	if err != nil {
		return ""
	}
	inputSize := len(bzInt)

	var bzStr []byte

	// TODO: Remove trailing zeros
	// case 1, purely decimal
	if inputSize <= Precision {
		bzStr = make([]byte, Precision+2)

		// 0. prefix
		bzStr[0] = byte('0')
		bzStr[1] = byte('.')

		// set relevant digits to 0
		for i := 0; i < Precision-inputSize; i++ {
			bzStr[i+2] = byte('0')
		}

		// set final digits
		copy(bzStr[2+(Precision-inputSize):], bzInt)

	} else {

		// inputSize + 1 to account for the decimal point that is being added
		bzStr = make([]byte, inputSize+1)
		decPointPlace := inputSize - Precision

		copy(bzStr, bzInt[:decPointPlace])                   // pre-decimal digits
		bzStr[decPointPlace] = byte('.')                     // decimal point
		copy(bzStr[decPointPlace+1:], bzInt[decPointPlace:]) // post-decimal digits
	}

	if isNeg {
		return "-" + string(bzStr)
	}

	return string(bzStr)
}

//     ____
//  __|    |__   "chop 'em
//       ` \     round!"
// ___||  ~  _     -bankers
// |         |      __
// |       | |   __|__|__
// |_____:  /   | $$$    |
//              |________|

// nolint - go-cyclo
// Remove a Precision amount of rightmost digits and perform bankers rounding
// on the remainder (gaussian rounding) on the digits which have been removed.
//
// Mutates the input. Use the non-mutative version if that is undesired
func chopPrecisionAndRound(d *big.Int) *big.Int {

	// remove the negative and add it back when returning
	if d.Sign() == -1 {
		// make d positive, compute chopped value, and then un-mutate d
		d = d.Neg(d)
		d = chopPrecisionAndRound(d)
		d = d.Neg(d)
		return d
	}

	// get the truncated quotient and remainder
	quo, rem := d, big.NewInt(0)
	quo, rem = quo.QuoRem(d, precisionReuse, rem)

	if rem.Sign() == 0 { // remainder is zero
		return quo
	}

	switch rem.Cmp(fivePrecision) {
	case -1:
		return quo
	case 1:
		return quo.Add(quo, oneInt)
	default: // bankers rounding must take place
		// always round to an even number
		if quo.Bit(0) == 0 {
			return quo
		}
		return quo.Add(quo, oneInt)
	}
}

//func chopPrecisionAndRoundUp(d *big.Int) *big.Int {
//
//	// remove the negative and add it back when returning
//	if d.Sign() == -1 {
//		// make d positive, compute chopped value, and then un-mutate d
//		d = d.Neg(d)
//		// truncate since d is negative...
//		d = chopPrecisionAndTruncate(d)
//		d = d.Neg(d)
//		return d
//	}
//
//	// get the truncated quotient and remainder
//	quo, rem := d, big.NewInt(0)
//	quo, rem = quo.QuoRem(d, precisionReuse, rem)
//
//	if rem.Sign() == 0 { // remainder is zero
//		return quo
//	}
//
//	return quo.Add(quo, oneInt)
//}

func chopPrecisionAndRoundNonMutative(d *big.Int) *big.Int {
	tmp := new(big.Int).Set(d)
	return chopPrecisionAndRound(tmp)
}

// RoundInt round the decimal using bankers rounding
func (d Dec) RoundInt() Int {
	return NewIntFromBigInt(chopPrecisionAndRoundNonMutative(d.Int))
}

// ___________________________________________________________________________________
// Ceil returns the smallest interger value (as a decimal) that is greater than
// or equal to the given decimal.
func (d Dec) Ceil() Dec {
	tmp := new(big.Int).Set(d.Int)

	quo, rem := tmp, big.NewInt(0)
	quo, rem = quo.QuoRem(tmp, precisionReuse, rem)

	// no need to round with a zero remainder regardless of sign
	if rem.Cmp(zeroInt) == 0 {
		return NewDecFromBigInt(quo)
	}

	if rem.Sign() == -1 {
		return NewDecFromBigInt(quo)
	}

	return NewDecFromBigInt(quo.Add(quo, oneInt))
}

//___________________________________________________________________________________

// UnmarshalJSON defines custom decoding scheme
func (d *Dec) UnmarshalJSON(bz []byte) error {
	if d.Int == nil {
		d.Int = new(big.Int)
	}

	var text string
	err := json.Unmarshal(bz, &text)
	if err != nil {
		return err
	}
	// TODO: Reuse dec allocation
	newDec, err := NewDecFromStr(text)
	if err != nil {
		return err
	}
	d.Int = newDec.Int
	return nil
}

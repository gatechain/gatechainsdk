package types

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Decimal Coin

// Coins which can have additional decimal points
type DecCoin struct {
	Denom  string `json:"denom"`
	Amount Dec    `json:"amount"`
}

// String implements the Stringer interface for DecCoin. It returns a
// human-readable representation of a decimal coin.
func (coin DecCoin) String() string {
	return fmt.Sprintf("%v%v", coin.Amount, coin.Denom)
}

// ----------------------------------------------------------------------------
// Decimal Coins

// coins with decimal
type DecCoins []DecCoin

// String implements the Stringer interface for DecCoins. It returns a
// human-readable representation of decimal coins.
func (coins DecCoins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _, coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}

	return out[:len(out)-1]
}

// return whether all coins are zero
func (coins DecCoins) IsZero() bool {
	for _, coin := range coins {
		if !coin.Amount.IsZero() {
			return false
		}
	}
	return true
}

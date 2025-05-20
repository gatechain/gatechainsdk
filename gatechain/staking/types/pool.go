package types

import (
	"fmt"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// Pool - tracking bonded and not-bonded token supply of the bond denomination
type Pool struct {
	NotBondedTokens sdk.Int `json:"not_bonded_tokens" yaml:"not_bonded_tokens"` // tokens which are not bonded to a validator (unbonded or unbonding)
	BondedTokens    sdk.Int `json:"bonded_tokens" yaml:"bonded_tokens"`         // tokens which are currently bonded to a validator
}

// String returns a human readable string representation of a pool.
func (p Pool) String() string {
	return fmt.Sprintf(`Pool:
 Not Bonded Tokens:  %s
 Bonded Tokens:      %s`, p.NotBondedTokens,
		p.BondedTokens)
}

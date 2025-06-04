package auth

import (
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// QueryAccountParams defines the params for querying accounts.
type QueryVaultAccountParams struct {
	Address types.AccAddress
}

// NewQueryAccountParams creates a new instance of QueryAccountParams.
func NewQueryVaultAccountParams(addr types.AccAddress) QueryVaultAccountParams {
	return QueryVaultAccountParams{Address: addr}
}

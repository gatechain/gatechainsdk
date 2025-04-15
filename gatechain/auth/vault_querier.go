package auth

import (
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// query endpoints supported by the auth Querier
const (
	QueryRevocable   = "revocable"
	QueryAccountType = "check"
)

// QueryAccountParams defines the params for querying accounts.
type QueryVaultAccountParams struct {
	Address types.AccAddress
}

// NewQueryAccountParams creates a new instance of QueryAccountParams.
func NewQueryVaultAccountParams(addr types.AccAddress) QueryVaultAccountParams {
	return QueryVaultAccountParams{Address: addr}
}

// defines the params for query: "custom/account/revocable"
type QueryRevocableTokensParams struct {
	Address types.AccAddress
	Height  int64
}

func NewQueryRevocableTokensParams(addr types.AccAddress, height int64) QueryRevocableTokensParams {
	return QueryRevocableTokensParams{
		Address: addr,
		Height:  height,
	}
}

// QueryAccountTypeParams defines the params for querying accounts.
type QueryAccountTypeParams struct {
	Address types.AccAddress
}

// NewQueryAccountTypeParams creates a new instance of QueryAccountTypeParams.
func NewQueryAccountTypeParams(addr types.AccAddress) QueryAccountTypeParams {
	return QueryAccountTypeParams{Address: addr}
}

type AccountResponse struct {
	Name          string      `json:"name,omitempty"`
	Coins         types.Coins `json:"tokens,omitempty"`
	Address       string      `json:"address,omitempty"`
	AccountNumber uint64      `json:"account_number,omitempty"`
	Sequence      uint64      `json:"sequence,omitempty"`
}

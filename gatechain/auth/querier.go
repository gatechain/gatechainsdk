package auth

import (
	"github.com/gatechain/gatechainsdk/gatechain/auth/exported"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// query endpoints supported by the auth Querier
const (
	QueryAccount = "account"
)

// QueryAccountParams defines the params for querying accounts.
type QueryAccountParams struct {
	Address sdk.AccAddress
}

// NewQueryAccountParams creates a new instance of QueryAccountParams.
func NewQueryAccountParams(addr sdk.AccAddress) QueryAccountParams {
	return QueryAccountParams{Address: addr}
}

type BaseAccountResp struct {
	exported.VaultAccount `json:"account_field" yaml:"account_field"`
	AccountType           uint8 `json:"account_type" yaml:"account_type"`
}

func (bar BaseAccountResp) GetAccountType() uint8 {
	return bar.AccountType
}

package auth

import (
	exported2 "github.com/gatechain/gatechainsdk/gatechain/auth/exported"
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
	exported2.VaultAccount `json:"account_field" yaml:"account_field"`
	AccountType            uint8 `json:"account_type" yaml:"account_type"`
}
type BaseEvmAccountResp struct {
	exported2.Account `json:"account_field" yaml:"account_field"`
	AccountType       uint8 `json:"account_type" yaml:"account_type"`
}

func (bar BaseAccountResp) GetAccountType() uint8 {
	return bar.AccountType
}

// TODO FIXME
func GetAccountRespFromVault(account exported2.Account) exported2.Account {
	vault, ok := account.(exported2.VaultAccount)
	if ok {
		return &BaseAccountResp{
			VaultAccount: vault,
			AccountType:  vault.GetAccountType(),
		}
	} else {
		return &BaseEvmAccountResp{
			Account:     account,
			AccountType: sdk.EvmStandardAccount,
		}
	}

}

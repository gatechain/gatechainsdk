package auth

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/auth/exported"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// NodeVaultQuerier is an interface that is satisfied by types that provide the QueryWithData method
type NodeVaultQuerier interface {
	// QueryWithData performs a query to a Tendermint node with the provided path
	// and a data payload. It returns the result and height of the query upon success
	// or an error if the query fails.
	QueryWithData(path string, data []byte) ([]byte, int64, error)

	GetCodec() *codec.Codec
}

// VaultRetriever defines the properties of a type that can be used to
// retrieve vault accounts.
type VaultRetriever struct {
	querier NodeVaultQuerier
}

// NewAccountRetriever initialises a new AccountRetriever instance.
func NewVaultRetriever(querier NodeVaultQuerier) VaultRetriever {
	return VaultRetriever{querier: querier}
}

// GetAccount queries for an account given an address and a block height. An
// error is returned if the query or decoding fails.
func (ar VaultRetriever) GetAccount(addr sdk.AccAddress) (exported.VaultAccount, error) {
	account, _, err := ar.GetAccountWithHeight(addr)
	return account, err
}

// GetAccountWithHeight queries for an account given an address. Returns the
// height of the query with the account. An error is returned if the query
// or decoding fails.
func (ar VaultRetriever) GetAccountWithHeight(addr sdk.AccAddress) (exported.VaultAccount, int64, error) {
	bs, err := ar.querier.GetCodec().MarshalJSON(NewQueryVaultAccountParams(addr))
	if err != nil {
		return nil, 0, err
	}

	res, height, err := ar.querier.QueryWithData(fmt.Sprintf("custom/%s/%s", QuerierRoute, QueryAccount), bs)
	if err != nil {
		return nil, height, err
	}

	var account exported.Account
	if err := ar.querier.GetCodec().UnmarshalJSON(res, &account); err != nil {
		return nil, height, err
	}

	return account.(exported.VaultAccount), height, nil
}

// EnsureExists returns an error if no account exists for the given address else nil.
func (ar VaultRetriever) EnsureExists(addr sdk.AccAddress) error {
	if _, err := ar.GetAccount(addr); err != nil {
		return err
	}
	return nil
}

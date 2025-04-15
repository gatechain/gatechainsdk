package types

import (
	"fmt"
)

// NodeQuerier is an interface that is satisfied by types that provide the QueryWithData method
type NodeQuerier interface {
	// QueryWithData performs a query to a Tendermint node with the provided path
	// and a data payload. It returns the result and height of the query upon success
	// or an error if the query fails.
	QueryWithData(path string, data []byte) ([]byte, int64, error)
}

// VaultRetriever defines the properties of a type that can be used to
// retrieve vault accounts.
type RevocableRetriever struct {
	querier NodeQuerier
}

// NewRevocableRetriever initialises a new RevocableRRetriever instance.
func NewRevocableRetriever(querier NodeQuerier) RevocableRetriever {
	return RevocableRetriever{querier: querier}
}

// GetAccount queries for an account given an address and a block height. An
// error is returned if the query or decoding fails.
func (ar RevocableRetriever) GetTx(hash []byte) (*RevocableTx, error) {
	bs, err := ModuleCdc.MarshalJSON(NewQueryRevocableTxParams(hash))
	if err != nil {
		return nil, err
	}

	res, _, err := ar.querier.QueryWithData(fmt.Sprintf("custom/%s/%s", QuerierRoute, QueryTx), bs)
	if err != nil {
		return nil, err
	}

	var tx RevocableTx
	if err := ModuleCdc.UnmarshalJSON(res, &tx); err != nil {
		return nil,  err
	}

	return &tx,  nil
}

// EnsureExists returns an error if no account exists for the given address else nil.
func (ar RevocableRetriever) EnsureExists(hash []byte) error {
	if _, err := ar.GetTx(hash); err != nil {
		return err
	}
	return nil
}

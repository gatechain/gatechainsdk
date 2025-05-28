package types

// query endpoints supported by the auth Querier
const (
	QueryTx = "tx"
)

// QueryAccountParams defines the params for querying accounts.
type QueryRevocableTxParams struct {
	Hash []byte
}

// NewQueryAccountParams creates a new instance of QueryAccountParams.
func NewQueryRevocableTxParams(hash []byte) QueryRevocableTxParams {
	return QueryRevocableTxParams{Hash: hash}
}

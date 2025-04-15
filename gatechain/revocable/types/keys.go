package types

const (
	// module name
	ModuleName = "revocable"

	// StoreKey is string representation of the store key for auth
	StoreKey = "revocable"

	// QuerierRoute is the querier route for acc
	QuerierRoute = StoreKey
)

var (
	// AddressStoreKeyPrefix prefix for account-by-address store
	TxStoreKeyPrefix = []byte{'r', 'e', 'v'}

	// param key for global account number
	GlobalAccountNumberKey = []byte("globalAccountNumber")
)

// TxStoreKey turn an tx to key used to get it from the account store
func TxStoreKey(tx []byte) []byte {
	return append(TxStoreKeyPrefix, tx...)
}

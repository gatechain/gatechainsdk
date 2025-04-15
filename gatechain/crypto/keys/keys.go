package keys

// SigningAlgo defines an algorithm to derive key-pairs which can be used for cryptographic signing.
type SigningAlgo string

const (
	// Ed25519 uses the Bitcoin secp256k1 ECDSA parameters.
	Ed25519 = SigningAlgo("ed25519")
	// Secp256k1 represents the Secp256k1 signature system.
	// It is currently not supported for end-user keys (wallets/ledgers).
	Secp256k1 = SigningAlgo("secp256k1")
)

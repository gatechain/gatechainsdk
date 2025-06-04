package auth

import (
	"github.com/gatechain/crypto"
	"github.com/gatechain/crypto/multisig"
	"github.com/gatechain/gatechainsdk/gatechain/auth/exported"
)

func countSubKeys(pub crypto.PubKey) int {
	v, ok := pub.(multisig.PubKeyMultisigThreshold)
	if !ok {
		return 1
	}

	numKeys := 0
	for _, subkey := range v.PubKeys {
		numKeys += countSubKeys(subkey)
	}

	return numKeys
}

func IsMultiSignAccount(acc exported.Account) bool {
	return countSubKeys(acc.GetPubKey()) > 1
}

func IsVaultAccount(acc exported.Account) bool {
	switch acc.(type) {
	case *VaultAccount:
		return !acc.(*VaultAccount).SecurityAddress.Empty()
	default:
		return false
	}
	return false
}

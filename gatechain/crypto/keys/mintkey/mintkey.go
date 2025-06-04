package mintkey

import (
	"encoding/hex"
	"fmt"

	"github.com/gatechain/crypto"
	"github.com/gatechain/crypto/armor"
	"github.com/gatechain/crypto/bcrypt"
	cryptoAmino "github.com/gatechain/crypto/encoding/amino"
	"github.com/gatechain/crypto/xsalsa20symmetric"

	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys/keyerror"
)

const (
	blockTypePrivKey = "GATECHAIN PRIVATE KEY"
)

// Make bcrypt security parameter var, so it can be changed within the lcd test
// Making the bcrypt security parameter a var shouldn't be a security issue:
// One can't verify an invalid key by maliciously changing the bcrypt
// parameter during a runtime vulnerability. The main security
// threat this then exposes would be something that changes this during
// runtime before the user creates their key. This vulnerability must
// succeed to update this to that same value before every subsequent call
// to the keys command in future startups / or the attacker must get access
// to the filesystem. However, with a similar threat model (changing
// variables in runtime), one can cause the user to sign a different tx
// than what they see, which is a significantly cheaper attack then breaking
// a bcrypt hash. (Recall that the nonce still exists to break rainbow tables)
// For further notes on security parameter choice, see README.md
var BcryptSecurityParameter = 12

//-----------------------------------------------------------------
// encrypt/decrypt with armor

// Encrypt and armor the private key.
func EncryptArmorPrivKey(privKey crypto.PrivKey, passphrase string) string {
	saltBytes, encBytes := encryptPrivKey(privKey, passphrase)
	header := map[string]string{
		"kdf":  "bcrypt",
		"salt": fmt.Sprintf("%X", saltBytes),
	}
	armorStr := armor.EncodeArmor(blockTypePrivKey, header, encBytes)
	return armorStr
}

// encrypt the given privKey with the passphrase using a randomly
// generated salt and the xsalsa20 cipher. returns the salt and the
// encrypted priv key.
func encryptPrivKey(privKey crypto.PrivKey, passphrase string) (saltBytes []byte, encBytes []byte) {
	saltBytes = crypto.CRandBytes(16)
	key, err := bcrypt.GenerateFromPassword(saltBytes, []byte(passphrase), BcryptSecurityParameter)
	if err != nil {
		//cmn.Exit("Error generating bcrypt key from passphrase: " + err.Error())
	}
	key = crypto.Sha256(key) // get 32 bytes
	privKeyBytes := privKey.Bytes()
	return saltBytes, xsalsa20symmetric.EncryptSymmetric(privKeyBytes, key)
}

// Unarmor and decrypt the private key.
func UnarmorDecryptPrivKey(armorStr string, passphrase string) (crypto.PrivKey, error) {
	var privKey crypto.PrivKey
	blockType, header, encBytes, err := armor.DecodeArmor(armorStr)
	if err != nil {
		return privKey, err
	}
	if blockType != blockTypePrivKey {
		return privKey, fmt.Errorf("Unrecognized armor type: %v", blockType)
	}
	if header["kdf"] != "bcrypt" {
		return privKey, fmt.Errorf("Unrecognized KDF type: %v", header["KDF"])
	}
	if header["salt"] == "" {
		return privKey, fmt.Errorf("Missing salt bytes")
	}
	saltBytes, err := hex.DecodeString(header["salt"])
	if err != nil {
		return privKey, fmt.Errorf("Error decoding salt: %v", err.Error())
	}
	privKey, err = decryptPrivKey(saltBytes, encBytes, passphrase)
	return privKey, err
}

func decryptPrivKey(saltBytes []byte, encBytes []byte, passphrase string) (privKey crypto.PrivKey, err error) {
	key, err := bcrypt.GenerateFromPassword(saltBytes, []byte(passphrase), BcryptSecurityParameter)
	if err != nil {
		//cmn.Exit("Error generating bcrypt key from passphrase: " + err.Error())
	}
	key = crypto.Sha256(key) // Get 32 bytes
	privKeyBytes, err := xsalsa20symmetric.DecryptSymmetric(encBytes, key)
	if err != nil && err.Error() == "Ciphertext decryption failed" {
		return privKey, keyerror.NewErrWrongPassword()
	} else if err != nil {
		return privKey, err
	}
	privKey, err = cryptoAmino.PrivKeyFromBytes(privKeyBytes)
	return privKey, err
}

package keys

import (
	"fmt"
	"github.com/cosmos/go-bip39"
	tmcrypto "github.com/gatechain/crypto"
	"github.com/gatechain/crypto/ed25519x"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys/hd"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys/keyerror"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys/mintkey"
	types "github.com/gatechain/gatechainsdk/gatechain/types"
	dbm "github.com/tendermint/tm-db"
)

var _ Keybase = dbKeybase{}

// Language is a language to create the BIP 39 mnemonic in.
// Currently, only english is supported though.
// Find a list of all supported languages in the BIP 39 spec (word lists).
type Language int

// noinspection ALL
const (
	// English is the default language to create a mnemonic.
	// It is the only supported language by this package.
	English       Language = iota + 1
	addressSuffix          = "address"
	infoSuffix             = "info"
)

// dbKeybase combines encryption and storage implementation to provide
// a full-featured key manager
type dbKeybase struct {
	db dbm.DB
}

// newDbKeybase creates a new keybase instance using the passed DB for reading and writing keys.
func newDbKeybase(db dbm.DB) Keybase {
	return dbKeybase{
		db: db,
	}
}

// NewInMemory creates a transient keybase on top of in-memory storage
// instance useful for testing purposes and on-the-fly key generation.
func NewInMemory() Keybase { return dbKeybase{dbm.NewMemDB()} }

// CreateAccount converts a mnemonic to a private key and persists it, encrypted with the given password.
func (kb dbKeybase) CreateAccount(name, mnemonic, bip39Passwd, encryptPasswd string, account uint32, index uint32) (Info, error) {
	coinType := types.GetConfig().GetCoinType()
	hdPath := hd.NewFundraiserParams(account, coinType, index)
	return kb.Derive(name, mnemonic, bip39Passwd, encryptPasswd, *hdPath)
}

func (kb dbKeybase) Derive(name, mnemonic, bip39Passphrase, encryptPasswd string, params hd.BIP44Params) (info Info, err error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, bip39Passphrase)
	if err != nil {
		return
	}

	info, err = kb.persistDerivedKey(seed, encryptPasswd, name, params.String())
	return
}

func (kb *dbKeybase) persistDerivedKey(seed []byte, passwd, name, fullHdPath string) (info Info, err error) {
	masterPriv := ed25519x.GenerateXprvFromSeed(seed)
	derivedPriv, err := ed25519x.DerivePrivateKeyFromPath(masterPriv, fullHdPath)
	if err != nil {
		return
	}
	if passwd != "" {
		info = kb.writeLocalKey(name, derivedPriv, passwd)
	}
	return
}

// Get returns the public information about one key.
func (kb dbKeybase) Get(name string) (Info, error) {
	bs, err := kb.db.Get(infoKey(name))
	fmt.Println(err)

	if len(bs) == 0 {
		return nil, keyerror.NewErrKeyNotFound(name)
	}
	return readInfo(bs)
}

func (kb dbKeybase) GetByAddress(address types.AccAddress) (Info, error) {
	ik, err := kb.db.Get(addrKey(address))
	fmt.Println(err)

	if len(ik) == 0 {
		return nil, fmt.Errorf("key with address %s not found", address)
	}
	bs, err := kb.db.Get(ik)
	fmt.Println(err)
	return readInfo(bs)
}

// Sign signs the msg with the named key.
// It returns an error if the key doesn't exist or the decryption fails.
func (kb dbKeybase) Sign(name, passphrase string, msg []byte) (sig []byte, pub tmcrypto.PubKey, err error) {
	info, err := kb.Get(name)
	if err != nil {
		return
	}

	var priv tmcrypto.PrivKey

	switch info.(type) {
	case localInfo:
		linfo := info.(localInfo)
		if linfo.PrivKeyArmor == "" {
			err = fmt.Errorf("private key not available")
			return
		}

		priv, err = mintkey.UnarmorDecryptPrivKey(linfo.PrivKeyArmor, passphrase)
		if err != nil {
			return nil, nil, err
		}

	}

	sig, err = priv.Sign(msg)
	if err != nil {
		return nil, nil, err
	}

	pub = priv.PubKey()
	return sig, pub, nil
}

// CloseDB releases the lock and closes the storage backend.
func (kb dbKeybase) CloseDB() {
	kb.db.Close()
}

func (kb dbKeybase) writeLocalKey(name string, priv tmcrypto.PrivKey, passphrase string) Info {
	// encrypt private key using passphrase
	privArmor := mintkey.EncryptArmorPrivKey(priv, passphrase)
	// make Info
	pub := priv.PubKey()
	info := newLocalInfo(name, pub, privArmor)
	kb.writeInfo(name, info)
	return info
}

func (kb dbKeybase) writeInfo(name string, info Info) {
	// write the info by key
	key := infoKey(name)
	serializedInfo := writeInfo(info)
	kb.db.SetSync(key, serializedInfo)
	// store a pointer to the infokey by address for fast lookup
	kb.db.SetSync(addrKey(info.GetAddress()), key)
}

func addrKey(address types.AccAddress) []byte {
	return []byte(fmt.Sprintf("%s.%s", address.String(), addressSuffix))
}

func infoKey(name string) []byte {
	return []byte(fmt.Sprintf("%s.%s", name, infoSuffix))
}

package keys

import (
	"fmt"
	"os"

	"github.com/gatechain/crypto"

	"github.com/gatechain/gatechainsdk/gatechain/types"
)

var _ Keybase = lazyKeybase{}

type lazyKeybase struct {
	name string
	dir  string
}

// New creates a new instance of a lazy keybase.
func New(name, dir string) Keybase {
	if err := EnsureDir(dir, 0700); err != nil {
		panic(fmt.Sprintf("failed to create Keybase directory: %s", err))
	}

	return lazyKeybase{name: name, dir: dir}
}

func (lkb lazyKeybase) Get(name string) (Info, error) {
	db, err := types.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDbKeybase(db).Get(name)
}

func (lkb lazyKeybase) GetByAddress(address types.AccAddress) (Info, error) {
	db, err := types.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDbKeybase(db).GetByAddress(address)
}

func (lkb lazyKeybase) Sign(name, passphrase string, msg []byte) ([]byte, crypto.PubKey, error) {
	db, err := types.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	return newDbKeybase(db).Sign(name, passphrase, msg)
}

func (lkb lazyKeybase) CreateAccount(name, mnemonic, bip39Passwd, encryptPasswd string, account uint32, index uint32) (Info, error) {
	db, err := types.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDbKeybase(db).CreateAccount(name, mnemonic, bip39Passwd, encryptPasswd, account, index)
}

func EnsureDir(dir string, mode os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, mode)
		if err != nil {
			return fmt.Errorf("Could not create directory %v. %v", dir, err)
		}
	}
	return nil
}

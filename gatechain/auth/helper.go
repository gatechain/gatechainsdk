package auth

import (
	"fmt"
	"github.com/gatechain/gatechainsdk/common"
	"strconv"
	"time"

	"github.com/cosmos/go-bip39"

	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
)

const (
	MnemonicEntropySize = 256
)

func CreateAccount(name string) (keys.Info, error) {
	var kb keys.Keybase
	var err error
	var encryptPassword = common.DefaultKeyPass

	if len(name) == 0 {
		name = fmt.Sprintf("%s", strconv.FormatInt(time.Now().Unix(), 10))
	}

	showMnemonic := true

	flagPwd := common.DefaultKeyPass
	if len(flagPwd) >= common.MinKeyPassLen {
		encryptPassword = flagPwd
	} else if len(flagPwd) > 0 {
		return nil, fmt.Errorf("password must be at least %d characters", common.MinKeyPassLen)
	}

	kb = keys.NewInMemory()

	account := uint32(0)
	index := uint32(0)

	// Get bip39 mnemonic
	var mnemonic string
	var bip39Passphrase string

	if len(mnemonic) == 0 {
		// read entropy seed straight from crypto.Rand and convert to mnemonic
		entropySeed, err := bip39.NewEntropy(MnemonicEntropySize)
		if err != nil {
			return nil, err
		}

		mnemonic, err = bip39.NewMnemonic(entropySeed[:])
		if err != nil {
			return nil, err
		}
	}

	info, err := kb.CreateAccount(name, mnemonic, bip39Passphrase, encryptPassword, account, index)
	if err != nil {
		return nil, err
	}
	printAccountInfo(info, showMnemonic, mnemonic)
	return info, nil
}

func printAccountInfo(info keys.Info, showMnemonic bool, mnemonic string) error {
	out, err := keys.Bech32KeyOutput(info)
	if err != nil {
		return err
	}

	if showMnemonic {
		out.Mnemonic = mnemonic
	}

	var jsonString []byte
	jsonString, err = codec.Cdc.MarshalJSONIndent(out, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonString))
	return nil
}

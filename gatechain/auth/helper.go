package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/go-bip39"

	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/crypto"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
)

const (
	MnemonicEntropySize = 256
)

func CreateAccount(name, rootDir string) error {
	var kb keys.Keybase
	var err error
	var encryptPassword = DefaultKeyPass

	if len(name) == 0 {
		name = fmt.Sprintf("%s", strconv.FormatInt(time.Now().Unix(), 10))
	}

	showMnemonic := true

	flagPwd := DefaultKeyPass
	if len(flagPwd) >= MinKeyPassLen {
		encryptPassword = flagPwd
	} else if len(flagPwd) > 0 {
		return fmt.Errorf("password must be at least %d characters", MinKeyPassLen)
	}
	dryRun := false
	if dryRun {
		// we throw this away, so don't enforce args,
		// we want to get a new random seed phrase quickly
		kb = keys.NewInMemory()
	} else {
		kb, err = crypto.NewKeyBaseFromDir(rootDir)
		if err != nil {
			//t.Log(err)
			return err
		}
		if err != nil {
			return err
		}

		_, err = kb.Get(name)
		if err == nil {
			return err
		}
	}

	//account := uint32(viper.GetInt(flagAccount))
	//index := uint32(viper.GetInt(flagIndex))
	account := uint32(0)
	index := uint32(0)

	// Get bip39 mnemonic
	var mnemonic string
	var bip39Passphrase string

	if len(mnemonic) == 0 {
		// read entropy seed straight from crypto.Rand and convert to mnemonic
		entropySeed, err := bip39.NewEntropy(MnemonicEntropySize)
		if err != nil {
			return err
		}

		mnemonic, err = bip39.NewMnemonic(entropySeed[:])
		if err != nil {
			return err
		}
	}

	info, err := kb.CreateAccount(name, mnemonic, bip39Passphrase, encryptPassword, account, index)
	if err != nil {
		return err
	}

	return printAccountInfo(info, showMnemonic, mnemonic)
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

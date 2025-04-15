package auth

import (
	"fmt"
	"github.com/cosmos/go-bip39"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/crypto"
	keys2 "github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
	"github.com/gatechain/gatechainsdk/gatechain/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"strconv"
	"time"
)

// available output formats.
const (
	OutputFormatText = "text"
	OutputFormatJSON = "json"

	// defaultKeyDBName is the client's subdirectory where keys are stored.
	defaultKeyDBName = "keys"
)

func CreateAccount(name, rootDir string) error {
	var kb keys2.Keybase
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
		kb = keys2.NewInMemory()
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
		entropySeed, err := bip39.NewEntropy(types.MnemonicEntropySize)
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

	return PrintCreateNoCMD(info, showMnemonic, mnemonic)
}

func PrintCreateNoCMD(info keys2.Info, showMnemonic bool, mnemonic string) error {
	out, err := keys2.Bech32KeyOutput(info)
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

func PrintCreate(cmd *cobra.Command, info keys2.Info, showMnemonic bool, mnemonic string) error {
	output := viper.Get(cli.OutputFlag)

	//cliCtx := context.NewCLIContext().WithCodec(codec.Cdc)
	output = OutputFormatJSON
	switch output {
	case OutputFormatText:
		cmd.PrintErrln()
		// print mnemonic unless requested not to.
		if showMnemonic {
			cmd.PrintErrln("\n**Important** write this mnemonic phrase in a safe place.")
			cmd.PrintErrln("It is the only way to recover your account if you ever forget your password.")
			cmd.PrintErrln("")
			cmd.PrintErrln(mnemonic)
		}
	case OutputFormatJSON:
		out, err := keys2.Bech32KeyOutput(info)
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
		cmd.PrintErrln(string(jsonString))
	default:
		return fmt.Errorf("I can't speak: %s", output)
	}

	return nil
}

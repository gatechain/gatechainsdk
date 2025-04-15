package keys

import (
	"bufio"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/gatechain/gatechainsdk/api/common"
	keys2 "github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
	"os"
	"path/filepath"
)

// available output formats.
const (
	OutputFormatText = "text"
	OutputFormatJSON = "json"

	// defaultKeyDBName is the client's subdirectory where keys are stored.
	defaultKeyDBName = "keys"
)

// GetKeyInfo returns key info for a given name. An error is returned if the
// keybase cannot be retrieved or getting the info fails.
func GetKeyInfo(name string) (keys2.Info, error) {
	keybase, err := NewKeyBaseFromDir(common.RootDir)
	if err != nil {
		return nil, err
	}

	return keybase.Get(name)
}

// GetPassphrase returns a passphrase for a given name. It will first retrieve
// the key info for that name if the type is local, it'll fetch input from
// STDIN. Otherwise, an empty passphrase is returned. An error is returned if
// the key info cannot be fetched or reading from STDIN fails.
func GetPassphrase(name string) (string, error) {
	var passphrase string

	keyInfo, err := GetKeyInfo(name)
	if err != nil {
		return passphrase, err
	}

	// we only need a passphrase for locally stored keys
	// TODO: (ref: #864) address security concerns
	if keyInfo.GetType() == keys2.TypeLocal {
		passphrase, err = ReadPassphraseFromStdin(name)
		if err != nil {
			return passphrase, err
		}
	}

	return passphrase, nil
}

// ReadPassphraseFromStdin attempts to read a passphrase from STDIN return an
// error upon failure.
func ReadPassphraseFromStdin(name string) (string, error) {
	buf := bufio.NewReader(os.Stdin)
	prompt := fmt.Sprintf("Password to sign with '%s':", name)

	passphrase, err := input.GetPassword(prompt, buf)
	if err != nil {
		return passphrase, fmt.Errorf("Error reading passphrase: %v", err)
	}

	return passphrase, nil
}

// NewKeyBaseFromDir initializes a keybase at a particular dir.
func NewKeyBaseFromDir(rootDir string) (keys2.Keybase, error) {
	return getLazyKeyBaseFromDir(rootDir)
}

func getLazyKeyBaseFromDir(rootDir string) (keys2.Keybase, error) {
	return keys2.New(defaultKeyDBName, filepath.Join(rootDir, "keys")), nil
}

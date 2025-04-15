package crypto

import (
	"path/filepath"

	keys2 "github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
)

// available output formats.
const (
	OutputFormatText = "text"
	OutputFormatJSON = "json"

	// defaultKeyDBName is the client's subdirectory where keys are stored.
	defaultKeyDBName = "keys"
)

// NewKeyBaseFromDir initializes a keybase at a particular dir.
func NewKeyBaseFromDir(rootDir string) (keys2.Keybase, error) {
	return getLazyKeyBaseFromDir(rootDir)
}

func getLazyKeyBaseFromDir(rootDir string) (keys2.Keybase, error) {
	return keys2.New(defaultKeyDBName, filepath.Join(rootDir, "keys")), nil
}

package keys

import (
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// KeyOutput defines a structure wrapping around an Info object used for output
// functionality.
type KeyOutput struct {
	Name      string                 `json:"name" yaml:"name"`
	Type      string                 `json:"type" yaml:"type"`
	Address   string                 `json:"address" yaml:"address"`
	PubKey    string                 `json:"pubkey" yaml:"pubkey"`
	Mnemonic  string                 `json:"mnemonic,omitempty" yaml:"mnemonic"`
	Threshold uint                   `json:"threshold,omitempty" yaml:"threshold"`
	PubKeys   []multisigPubKeyOutput `json:"pubkeys,omitempty" yaml:"pubkeys"`
}

// NewKeyOutput creates a default KeyOutput instance without Mnemonic, Threshold and PubKeys
func NewKeyOutput(name, keyType, address, pubkey string) KeyOutput {
	return KeyOutput{
		Name:    name,
		Type:    keyType,
		Address: address,
		PubKey:  pubkey,
	}
}

type multisigPubKeyOutput struct {
	Address string `json:"address" yaml:"address"`
	PubKey  string `json:"pubkey" yaml:"pubkey"`
	Weight  uint   `json:"weight" yaml:"weight"`
}

// Bech32KeyOutput create a KeyOutput in with "acc" Bech32 prefixes. If the
// public key is a multisig public key, then the threshold and constituent
// public keys will be added.
func Bech32KeyOutput(keyInfo Info) (KeyOutput, error) {
	accAddr := sdk.AccAddress(keyInfo.GetPubKey().Address().Bytes())
	bechPubKey, err := sdk.Bech32ifyAccPub(keyInfo.GetPubKey())
	if err != nil {
		return KeyOutput{}, err
	}

	ko := NewKeyOutput(keyInfo.GetName(), keyInfo.GetType().String(), accAddr.String(), bechPubKey)

	return ko, nil
}

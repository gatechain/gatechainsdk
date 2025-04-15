package types

import (
	"encoding/hex"
	"errors"
	"fmt"
	_ "fmt"
	_ "github.com/gatechain/gatechainsdk/gatechain/auth/exported"
	"gopkg.in/yaml.v2"
	_ "time"
)

// -----------------------------------------------------------------------------
// Revoke tx
type RevocableTx struct {
	Hash         []byte `json:"tx_hash" yaml:"tx_hash"`
	RevokeTxHash []byte `json:"revoke_hash" yaml:"revoke_hash"`
	RevokeHeight uint64 `json:"revoke_height" yaml:"revoke_height"`
}

func (rx RevocableTx) MarshalYAML() (interface{}, error) {
	var bs []byte
	var err error
	if rx.RevokeHeight != 0 {
		bs, err = yaml.Marshal(struct {
			Hash         string
			RevokeTxHash string
			RevokeHeight uint64
		}{
			Hash:         "REVOCABLEPAY-" + hex.EncodeToString(rx.Hash),
			RevokeTxHash: "REVOKE-" + hex.EncodeToString(rx.RevokeTxHash),
			RevokeHeight: rx.RevokeHeight,
		})
	} else {
		bs, err = yaml.Marshal(struct {
			Hash string
		}{
			Hash: "REVOCABLEPAY-" + hex.EncodeToString(rx.Hash),
		})
	}
	if err != nil {
		return nil, err
	}

	return string(bs), err
}

func (rx RevocableTx) String() string {

	if rx.RevokeHeight != 0 {
		return fmt.Sprintf(`RevocableTx:
  Hash:                    %s
  RevokeTxHash:            %s
  RevokeHeight:            %d`,
			"REVOCABLEPAY-"+hex.EncodeToString(rx.Hash),
			"REVOKE-"+hex.EncodeToString(rx.RevokeTxHash),
			rx.RevokeHeight,
		)
	} else {
		return fmt.Sprintf(`RevocableTx:
  Hash:                    %s`,
			"REVOCABLEPAY-"+hex.EncodeToString(rx.Hash),
		)
	}
}

func NewRevocableTx(hash []byte) *RevocableTx {
	return &RevocableTx{
		Hash: hash,
	}
}

func (rx RevocableTx) IsRevocable() bool {
	if rx.RevokeHeight == uint64(0) {
		return true
	} else {
		return false
	}
}

func (rx RevocableTx) GetHash() []byte {
	return rx.Hash
}

func (rx *RevocableTx) SetRevokeTx(revokeHeight uint64, tx []byte) error {
	if !rx.IsRevocable() {
		return errors.New("this tx cannot be revoke")
	}

	rx.RevokeHeight = revokeHeight
	rx.RevokeTxHash = tx

	return nil
}

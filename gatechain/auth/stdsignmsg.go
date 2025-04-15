package auth

import (
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// StdSignMsg is a convenience structure for passing along
// a Msg with the other requirements for a StdSignDoc before
// it is signed. For use in the CLI.
type StdSignMsg struct {
	ChainID     string    `json:"chain_id" yaml:"chain_id"`
	Nonces      [][]byte  `json:"nonces" yaml:"nonces"`
	Fee         StdFee    `json:"fee" yaml:"fee"`
	Msgs        []sdk.Msg `json:"msgs" yaml:"msgs"`
	Memo        string    `json:"memo" yaml:"memo"`
	ValidHeight []uint64  `json:"valid_height" yaml:"valid_height"`
}

// get message bytes
func (msg StdSignMsg) Bytes() []byte {
	return StdSignBytes(msg.ChainID, msg.Nonces, msg.Fee, msg.Msgs, msg.Memo, msg.ValidHeight)
}

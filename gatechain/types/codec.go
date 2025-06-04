package types

import (
	"github.com/gatechain/gatechainsdk/gatechain/codec"
)

// Register the framework message type
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*Msg)(nil), nil)
	cdc.RegisterInterface((*Tx)(nil), nil)
}

package types

import (
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	amino "github.com/tendermint/go-amino"
)

func RegisterAmino(cdc *amino.Codec) {
	//RegisterEventDatas(cdc)
	// types.RegisterBlockAmino(cdc)
}

// Register the framework message type
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*Msg)(nil), nil)
	cdc.RegisterInterface((*Tx)(nil), nil)
}

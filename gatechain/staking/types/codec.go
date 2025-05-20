package types

import (
	"github.com/gatechain/gatechainsdk/gatechain/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgDelegate{}, "MsgDelegate", nil)
	cdc.RegisterConcrete(MsgUndelegate{}, "MsgUndelegate", nil)

	cdc.RegisterConcrete(MsgBeginRedelegate{}, "MsgBeginRedelegate", nil)
	cdc.RegisterConcrete(MsgUndelegateByRetrievalAccount{}, "MsgUndelegateByRetrievalAccount", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

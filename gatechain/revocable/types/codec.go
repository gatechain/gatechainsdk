package types

import (
	"github.com/gatechain/gatechainsdk/gatechain/codec"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateVault{}, "MsgCreateVault", nil)
	cdc.RegisterConcrete(MsgRevocableSend{}, "MsgRevocableSend", nil)
	cdc.RegisterConcrete(MsgRevoke{}, "MsgRevoke", nil)
	cdc.RegisterConcrete(MsgUpdateClearingHeight{}, "MsgUpdateClearingHeight", nil)
	cdc.RegisterConcrete(MsgClearVaultAccount{}, "MsgClearVaultAccount", nil)
	cdc.RegisterConcrete(MsgPublishMultiSigAccount{}, "MsgPublishMultiSigAccount", nil)
}

// module wide codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

package auth

import (
	"github.com/gatechain/gatechainsdk/gatechain/auth/exported"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(StdTx{}, "StdTx", nil)
	cdc.RegisterInterface((*exported.VaultAccount)(nil), nil)
	cdc.RegisterConcrete(&VaultAccount{}, "VaultAccount", nil)
	cdc.RegisterConcrete(&BaseAccountResp{}, "AccountResp", nil)
}

// module wide codec
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

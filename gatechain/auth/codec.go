package auth

import (
	exported2 "github.com/gatechain/gatechainsdk/gatechain/auth/exported"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
)

//// RegisterCodec registers concrete types on the codec
//func RegisterCodec(cdc *codec.Codec) {
//	cdc.RegisterConcrete(StdTx{}, "StdTx", nil)
//}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*exported2.Account)(nil), nil)
	cdc.RegisterConcrete(&BaseAccount{}, "Account", nil)
	cdc.RegisterInterface((*exported2.VestingAccount)(nil), nil)
	cdc.RegisterConcrete(&BaseVestingAccount{}, "BaseVestingAccount", nil)
	cdc.RegisterConcrete(&ContinuousVestingAccount{}, "ContinuousVestingAccount", nil)
	cdc.RegisterConcrete(&DelayedVestingAccount{}, "DelayedVestingAccount", nil)
	cdc.RegisterConcrete(StdTx{}, "StdTx", nil)
	cdc.RegisterInterface((*exported2.VaultAccount)(nil), nil)
	cdc.RegisterConcrete(&VaultAccount{}, "VaultAccount", nil)
	//cdc.RegisterInterface((*exported.EvmAccount)(nil), nil)
	cdc.RegisterConcrete(&BaseAccountResp{}, "AccountResp", nil)
	cdc.RegisterConcrete(&BaseEvmAccountResp{}, "EvmAccountResp", nil)
	//cdc.RegisterConcrete(&EthAccount{}, "EthAccount", nil)
}

// module wide codec
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

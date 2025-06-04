package utils

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/bank"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/types"
	revocable_types "github.com/gatechain/gatechainsdk/gatechain/revocable/types"
	staking_types "github.com/gatechain/gatechainsdk/gatechain/staking/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// custom tx codec
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	types.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	bank.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	revocable_types.RegisterCodec(cdc)
	staking_types.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	return cdc
}

func JsonOutPut(o interface{}, cdc *codec.Codec) error {
	var jsonString []byte
	jsonString, err := cdc.MarshalJSONIndent(o, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonString))
	return nil
}

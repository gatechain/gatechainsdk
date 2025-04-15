package utils

import (
	"fmt"
	"net/url"
	
	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/bank"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/types"
	types2 "github.com/gatechain/gatechainsdk/gatechain/revocable/types"
	"github.com/gatechain/gatechainsdk/gatechain/rpc/client"
	staking_types "github.com/gatechain/gatechainsdk/gatechain/staking/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

const (
	MaxMessagesPerTxDefault = 5
)

// custom tx codec
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	types.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	bank.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	types2.RegisterCodec(cdc)
	codec.RegisterEvidences(cdc)
	staking_types.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	return cdc
}

func MakeRestClient(endpoint, apitoken string) *client.RestClient {
	clientUrl, err := url.Parse(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	restClient := client.MakeRestClient(*clientUrl, apitoken)
	return &restClient
}

func JsonOutPut(o interface{}, cdc *codec.Codec) {
	var jsonString []byte
	jsonString, err := cdc.MarshalJSONIndent(o, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonString))
}

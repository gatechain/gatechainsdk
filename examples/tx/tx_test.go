package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestQueryTx(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	hashHexStr := "IRREVOCABLEPAY-08A79A9031ACB80A848D19ADE96FCD39FF6A9510752A36A211A4B1E029A378EC1A687DA4704012696B36FED883DC89AF"
	client.Tx.QueryTX(hashHexStr)
}

func TestBroadcastTX(t *testing.T) {
	from_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClientWithFrom(from_addr, common.EndPoint, common.APIToken, common.RootDir)

	to_addr := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f"
	amount := "100000000000NANOGT"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client.Tx.SendTX(from_addr, to_addr, amount, fees, chainID, gas)
}

func TestCreateTxJsonFile(t *testing.T) {
	from_addr := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f" //
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	to_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	amount := "100000000000NANOGT"

	client := client.NewClientWithFrom(from_addr, common.EndPoint, common.APIToken, common.RootDir)
	client.Tx.CreateUnsignTX(from_addr, to_addr, amount, common.UnsignTxFile, fees, chainID, gas)

}

func TestSignTxJsonFile(t *testing.T) {
	from_addr := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f" //
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	client := client.NewClientWithFrom(from_addr, common.EndPoint, common.APIToken, common.RootDir)
	client.Tx.CreateSignTX(common.UnsignTxFile, common.SignTxFile, fees, chainID, gas)
}

func TestBroadCastTxJsonFile(t *testing.T) {
	from_addr := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f" //
	client := client.NewClientWithFrom(from_addr, common.EndPoint, common.APIToken, common.RootDir)
	client.Tx.BroadcastSignTx(common.SignTxFile)
}

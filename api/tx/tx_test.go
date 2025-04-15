package tx

import (
	"fmt"
	"testing"

	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

// done
func TestBroadcastTX(t *testing.T) {
	from_addr := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f" //
	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(from_addr, common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)

	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}

	to_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	amount := "100000000000NANOGT"
	SendTX(ctx, txBldr, from_addr, to_addr, amount)
}

// done
func TestCreateTxJsonFile(t *testing.T) {
	from_addr := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f" //
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImplGenOnly(true, common.UnsignTxFile, common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)
	curHeight, err := ctx.GetChainHeight()
	if err != nil {
		t.Fatal(err)
	}
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	startHeight := uint64(curHeight - 10)   //uint64(currentHeight) - 10
	expireHeight := uint64(curHeight + 200) //currentHeight + 200
	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr = txBldr.WithValidHeight([]uint64{startHeight, expireHeight})

	to_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	amount := "100000000000NANOGT"
	CreateUnsignTX(ctx, txBldr, from_addr, to_addr, amount, ctx.OutputFileName)
}

func TestSignTxJsonFile(t *testing.T) {
	from_addr := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f" //

	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(from_addr, common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)

	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)

	CreateSignTX(common.UnsignTxFile, common.SignTxFile, ctx, txBldr)
}

func TestBroadCastTxJsonFile(t *testing.T) {
	from_addr := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f" //

	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(from_addr, common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)

	BroadcastSignTx(ctx, common.SignTxFile)
}

// done
func TestQueryTx(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)
	hashHexStr := "IRREVOCABLEPAY-08A79A9031ACB80A848D19ADE96FCD39FF6A9510752A36A211A4B1E029A378EC1A687DA4704012696B36FED883DC89AF"
	QueryTX(ctx, hashHexStr)
}

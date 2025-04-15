package account

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

func TestCreateAccount(t *testing.T) {
	name := ""
	rootDir := common.RootDir
	CreateAccount(name, rootDir)
}

func TestQueryAccount(t *testing.T) {
	addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)

	QueryAccount(ctx, addr)
}

func TestGetAccountBlance(t *testing.T) {
	addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)
	GetAccountBlance(ctx, addr)
}

package account

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/api/utils"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

func CreateAccount(name string, rootDir string) {
	auth2.CreateAccount(name, rootDir)
}

func QueryAccount(ctx *context.NodeVaultQuerierImpl, addr string) {

	key, err := types.AccAddressFromBech32(addr)
	if err != nil {
		fmt.Println(err)
	}
	retriever := auth2.NewAccountRetriever(ctx, ctx.GetCodec())
	account, _, err := retriever.GetAccountWithHeight(key)
	if err != nil {
		fmt.Println(err)
	} else {
		utils.JsonOutPut(account, ctx.GetCodec())
	}
}

func GetAccountBlance(ctx *context.NodeVaultQuerierImpl, addr string) {

	retriever := auth2.NewVaultRetriever(ctx)
	key, err := types.AccAddressFromBech32(addr)
	if err != nil {
		fmt.Println(err)
	}
	if err := retriever.EnsureExists(key); err != nil {
		fmt.Println(err)
	}

	vault, height, err := retriever.GetAccountWithHeight(key)
	if err != nil {
		fmt.Println(err)
	}
	ctx = ctx.WithHeight(height)
	ctx.PrintOutput(vault.GetCoins())
}

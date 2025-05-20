package account

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

type Service struct {
	ctx *context.NodeVaultQuerierImpl
}

func NewService(ctx *context.NodeVaultQuerierImpl) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) CreateAccount(name string, rootDir string) {
	auth.CreateAccount(name, rootDir)
}

func (s *Service) QueryAccount(addr string) {

	key, _, err := types.AccAddressTypeFromBech32(addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	retriever := auth.NewVaultRetriever(s.ctx)

	if err := retriever.EnsureExists(key); err != nil {
		fmt.Println(err)
		return
	}

	account, height, err := retriever.GetAccountWithHeight(key)
	s.ctx = s.ctx.WithHeight(height)
	err = account.MergeRevocableWei(uint64(s.ctx.Height))
	if err != nil {
		fmt.Println(err)
		return
	}
	s.ctx.PrintOutput(account)
}

func (s *Service) GetAccountBlance(addr string) {

	retriever := auth.NewVaultRetriever(s.ctx)
	key, err := types.AccAddressFromBech32(addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := retriever.EnsureExists(key); err != nil {
		fmt.Println(err)
		return
	}

	vault, height, err := retriever.GetAccountWithHeight(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.ctx = s.ctx.WithHeight(height)
	s.ctx.PrintOutput(vault.GetCoins())
}

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

func (s *Service) CreateAccount(name string, rootDir string) error {
	return auth.CreateAccount(name, rootDir)
}

func (s *Service) QueryAccount(addr string) error {

	key, _, err := types.AccAddressTypeFromBech32(addr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	retriever := auth.NewVaultRetriever(s.ctx)

	if err := retriever.EnsureExists(key); err != nil {
		fmt.Println(err)
		return err
	}

	account, height, err := retriever.GetAccountWithHeight(key)
	s.ctx = s.ctx.WithHeight(height)
	err = account.MergeRevocableWei(uint64(s.ctx.Height))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return s.ctx.PrintOutput(account)
}

func (s *Service) GetAccountBlance(addr string) error {

	retriever := auth.NewVaultRetriever(s.ctx)
	key, err := types.AccAddressFromBech32(addr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := retriever.EnsureExists(key); err != nil {
		fmt.Println(err)
		return err
	}

	vault, height, err := retriever.GetAccountWithHeight(key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	s.ctx = s.ctx.WithHeight(height)
	return s.ctx.PrintOutput(vault.GetCoins())
}

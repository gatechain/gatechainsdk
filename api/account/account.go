package account

import (
	"github.com/gatechain/gatechainsdk/gatechain/auth/exported"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys"

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

func (s *Service) CreateAccount(name string) (keys.Info, error) {
	return auth.CreateAccount(name)
}

func (s *Service) QueryAccount(addr string) (exported.VaultAccount, error) {

	key, _, err := types.AccAddressTypeFromBech32(addr)
	if err != nil {
		return nil, err
	}
	retriever := auth.NewVaultRetriever(s.ctx)

	if err := retriever.EnsureExists(key); err != nil {
		return nil, err
	}

	account, height, err := retriever.GetAccountWithHeight(key)
	s.ctx = s.ctx.WithHeight(height)
	err = account.MergeRevocableWei(uint64(s.ctx.Height))
	if err != nil {
		return nil, err
	}
	s.ctx.PrintOutput(account)
	return account, nil
}

func (s *Service) GetAccountBlance(addr string) (types.Coins, error) {

	retriever := auth.NewVaultRetriever(s.ctx)
	key, err := types.AccAddressFromBech32(addr)
	if err != nil {
		return types.Coins{}, err
	}
	if err := retriever.EnsureExists(key); err != nil {
		return types.Coins{}, err
	}

	vault, height, err := retriever.GetAccountWithHeight(key)
	if err != nil {
		return types.Coins{}, err
	}
	s.ctx = s.ctx.WithHeight(height)
	s.ctx.PrintOutput(vault.GetCoins())
	return vault.GetCoins(), nil
}

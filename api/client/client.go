package client

import (
	"github.com/gatechain/gatechainsdk/api/account"
	"github.com/gatechain/gatechainsdk/api/block"
	"github.com/gatechain/gatechainsdk/api/distribution"
	"github.com/gatechain/gatechainsdk/api/revocableTx"
	"github.com/gatechain/gatechainsdk/api/staking"
	"github.com/gatechain/gatechainsdk/api/status"
	"github.com/gatechain/gatechainsdk/api/tx"
	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/api/validator"
	"github.com/gatechain/gatechainsdk/api/vaultAccount"
	"github.com/gatechain/gatechainsdk/api/version"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

type Client struct {
	Ctx          *context.NodeVaultQuerierImpl
	Account      *account.Service
	Block        *block.Service
	Version      *version.Service
	Status       *status.Service
	Tx           *tx.Service
	VaultAccount *vaultAccount.Service
	RevocableTx  *revocableTx.Service
	Staking      *staking.Service
	Distribution *distribution.Service
	Validator    *validator.Service
}

func NewClient(endpoint, API_TOKEN string) *Client {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(endpoint, API_TOKEN)
	ctx.WithCodec(cdc)

	accout := account.NewService(ctx)
	block := block.NewService(ctx)
	version := version.NewService(ctx)
	status := status.NewService(ctx)
	tx := tx.NewService(ctx)
	vaultAccount := vaultAccount.NewService(ctx)
	revocTx := revocableTx.NewService(ctx)
	staking := staking.NewService(ctx)
	distribution := distribution.NewService(ctx)
	validator := validator.NewService(ctx)

	return &Client{
		Ctx:          ctx,
		Account:      accout,
		Block:        block,
		Version:      version,
		Status:       status,
		Tx:           tx,
		VaultAccount: vaultAccount,
		RevocableTx:  revocTx,
		Staking:      staking,
		Distribution: distribution,
		Validator:    validator,
	}
}

func NewClientWithFrom(from, endpoint, API_TOKEN, rootDir string) *Client {
	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(from, endpoint, API_TOKEN, rootDir)
	ctx.WithCodec(cdc)
	staking := staking.NewService(ctx)
	accout := account.NewService(ctx)
	block := block.NewService(ctx)
	version := version.NewService(ctx)
	status := status.NewService(ctx)
	tx := tx.NewService(ctx)
	vaultAccount := vaultAccount.NewService(ctx)
	revocTx := revocableTx.NewService(ctx)
	distribution := distribution.NewService(ctx)
	validator := validator.NewService(ctx)

	return &Client{
		Ctx:          ctx,
		Account:      accout,
		Block:        block,
		Version:      version,
		Status:       status,
		Tx:           tx,
		VaultAccount: vaultAccount,
		RevocableTx:  revocTx,
		Staking:      staking,
		Distribution: distribution,
		Validator:    validator,
	}
}

func NewClientGenOnly(genOnly bool, fileName, endpoint, API_TOKEN, rootDir string) *Client {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImplGenOnly(genOnly, fileName, endpoint, API_TOKEN)
	ctx.WithCodec(cdc)
	staking := staking.NewService(ctx)
	accout := account.NewService(ctx)
	block := block.NewService(ctx)
	version := version.NewService(ctx)
	status := status.NewService(ctx)
	tx := tx.NewService(ctx)
	vaultAccount := vaultAccount.NewService(ctx)
	revocTx := revocableTx.NewService(ctx)
	distribution := distribution.NewService(ctx)
	validator := validator.NewService(ctx)

	return &Client{
		Ctx:          ctx,
		Account:      accout,
		Block:        block,
		Version:      version,
		Status:       status,
		Tx:           tx,
		VaultAccount: vaultAccount,
		RevocableTx:  revocTx,
		Staking:      staking,
		Distribution: distribution,
		Validator:    validator,
	}
}

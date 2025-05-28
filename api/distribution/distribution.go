package distribution

import (
	"fmt"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/client/common"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/client/cli"
)

type Service struct {
	ctx *context.NodeVaultQuerierImpl
}

func NewService(ctx *context.NodeVaultQuerierImpl) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) QueryParams() (common.PrettyParams, error) {
	return cli.QueryParams(s.ctx)
}

// QueryValidatorOutstandingRewards implements the query validator outstanding rewards command.
func (s *Service) QueryValidatorOutstandingRewards(ValidatorAddress string) (types.ValidatorOutstandingRewards, error) {
	return cli.QueryValidatorOutstandingRewards(s.ctx, ValidatorAddress)
}

// QueryValidatorCommission implements the query validator commission command.
func (s *Service) QueryValidatorCommission(validatorAddr string) (types.ValidatorAccumulatedCommission, error) {
	return cli.QueryValidatorCommission(s.ctx, validatorAddr)
}

// QueryDelegatorRewards implements the query delegator rewards command.
func (s *Service) QueryDelegatorRewards(args []string) error {
	return cli.QueryDelegatorRewards(s.ctx, args)
}

// QueryCommunityPool returns the command for fetching community pool info
func (s *Service) QueryCommunityPool() (sdk.DecCoins, error) {
	return cli.QueryCommunityPool(s.ctx)
}

// command to withdraw rewards
func (s *Service) WithdrawRewards(delegatorAddress, validatorAddress string, bComission bool, fees, chainID string, gas uint64) error {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return cli.WithdrawRewards(s.ctx, txBldr, delegatorAddress, validatorAddress, bComission)
}

// command to withdraw all rewards
func (s *Service) WithdrawAllRewards(delegatorAddr string, fees, chainID string, gas uint64) error {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return cli.WithdrawAllRewards(s.ctx, txBldr, delegatorAddr)
}

// command to replace a delegator's withdrawal address
func (s *Service) SetWithdrawAddr(delegatorAddress, withdrawAddress string, fees, chainID string, gas uint64) error {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return cli.SetWithdrawAddr(s.ctx, txBldr, delegatorAddress, withdrawAddress)
}

func (s *Service) RewardReinvestment(delegatorAddress, validatorAddress string, fees, chainID string, gas uint64) error {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return cli.RewardReinvestment(s.ctx, txBldr, delegatorAddress, validatorAddress)
}

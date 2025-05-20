package distribution

import (
	"fmt"

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

func (s *Service) QueryParams() {
	cli.QueryParams(s.ctx)
}

// QueryValidatorOutstandingRewards implements the query validator outstanding rewards command.
func (s *Service) QueryValidatorOutstandingRewards(ValidatorAddress string) {
	cli.QueryValidatorOutstandingRewards(s.ctx, ValidatorAddress)
}

// QueryValidatorCommission implements the query validator commission command.
func (s *Service) QueryValidatorCommission(validatorAddr string) {
	cli.QueryValidatorCommission(s.ctx, validatorAddr)
}

// QueryDelegatorRewards implements the query delegator rewards command.
func (s *Service) QueryDelegatorRewards(args []string) {
	cli.QueryDelegatorRewards(s.ctx, args)
}

// QueryCommunityPool returns the command for fetching community pool info
func (s *Service) QueryCommunityPool() {
	cli.QueryCommunityPool(s.ctx)
}

// command to withdraw rewards
func (s *Service) WithdrawRewards(delegatorAddress, validatorAddress string, bComission bool, fees, chainID string, gas uint64) {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return
	}
	cli.WithdrawRewards(s.ctx, txBldr, delegatorAddress, validatorAddress, bComission)
}

// command to withdraw all rewards
func (s *Service) WithdrawAllRewards(delegatorAddr string, fees, chainID string, gas uint64) {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return
	}
	cli.WithdrawAllRewards(s.ctx, txBldr, delegatorAddr)
}

// command to replace a delegator's withdrawal address
func (s *Service) SetWithdrawAddr(delegatorAddress, withdrawAddress string, fees, chainID string, gas uint64) {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return
	}
	cli.SetWithdrawAddr(s.ctx, txBldr, delegatorAddress, withdrawAddress)
}

func (s *Service) RewardReinvestment(delegatorAddress, validatorAddress string, fees, chainID string, gas uint64) {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return
	}
	cli.RewardReinvestment(s.ctx, txBldr, delegatorAddress, validatorAddress)
}

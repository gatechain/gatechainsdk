package staking

import (
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
	"github.com/gatechain/gatechainsdk/gatechain/staking/types"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/staking/client/cli"
)

type Service struct {
	ctx *context.NodeVaultQuerierImpl
}

func NewService(ctx *context.NodeVaultQuerierImpl) *Service {
	return &Service{ctx: ctx}
}

// QueryValidatorUnbondingDelegations implements the query all unbonding delegatations from a validator command.
func (s *Service) QueryValidatorUnbondingDelegations(ValidatorAddr string) (types.UnbondingDelegations, error) {
	return cli.QueryValidatorUnbondingDelegations(s.ctx, ValidatorAddr)
}

// QueryValidatorRedelegations implements the query all redelegatations
// from a validator command.
func (s *Service) QueryValidatorRedelegations(SrcValidatorAddr string) (types.RedelegationResponses, error) {
	return cli.QueryValidatorRedelegations(s.ctx, SrcValidatorAddr)
}

// QueryDelegation the query delegation command.
func (s *Service) QueryDelegation(delegatorAddr, validatorAddr string) (types.DelegationResponse, error) {
	return cli.QueryDelegation(s.ctx, delegatorAddr, validatorAddr)
}

// QueryDelegations implements the command to query all the delegations
// made from one delegator.
func (s *Service) QueryDelegations(delegatorAddr string) (types.DelegationResponses, error) {
	return cli.QueryDelegations(s.ctx, delegatorAddr)
}

// QueryValidatorDelegations implements the command to query all the
// delegations to a specific validator.
// , cdc *codec.Codec
func (s *Service) QueryValidatorDelegations(delegatorAddr string) (types.DelegationResponses, error) {
	return cli.QueryValidatorDelegations(s.ctx, delegatorAddr)
}

// QueryUnbondingDelegation implements the command to query a single
// unbonding-delegation record.
func (s *Service) QueryUnbondingDelegation(delegatorAddr, validatorAddr string) (types.UnbondingDelegation, error) {
	return cli.QueryUnbondingDelegation(s.ctx, delegatorAddr, validatorAddr)
}

// QueryUnbondingDelegations implements the command to query all the
// unbonding-delegation records for a delegator.
func (s *Service) QueryUnbondingDelegations(delegatorAddr string) (types.UnbondingDelegations, error) {
	return cli.QueryUnbondingDelegations(s.ctx, delegatorAddr)
}

func (s *Service) QueryRedelegation(args []string) (types.RedelegationResponses, error) {
	return cli.QueryRedelegation(s.ctx, args)
}

// QueryRedelegations implements the command to query all the
// redelegation records for a delegator.
func (s *Service) QueryRedelegations(delegatorAddr string) (types.RedelegationResponses, error) {
	return cli.QueryRedelegations(s.ctx, delegatorAddr)
}

// QueryPool implements the pool query command.
func (s *Service) QueryPool() (types.Pool, error) {
	return cli.QueryPool(s.ctx)
}

// QueryParams implements the params query command.
func (s *Service) QueryParams() (types.Params, error) {
	return cli.QueryParams(s.ctx)
}

// GetDelegate implements the delegate command.
func (s *Service) GetDelegate(amountStr, delegatorAddress, validatorAddr, fees, chainID string, gas uint64, kb keys.Keybase) error {
	txBldr := auth.NewTxBuilderFromCLIMemKeyBase(fees, chainID, gas, s.ctx.Codec, kb)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		return err
	}
	return cli.GetDelegate(s.ctx, txBldr, amountStr, delegatorAddress, validatorAddr)
}

// GetRedelegate the begin redelegation command.
func (s *Service) GetRedelegate(delAddr, fees, chainID string, gas uint64, args []string, kb keys.Keybase) error {
	txBldr := auth.NewTxBuilderFromCLIMemKeyBase(fees, chainID, gas, s.ctx.Codec, kb)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		return err
	}
	return cli.GetRedelegate(s.ctx, txBldr, delAddr, args)
}

// GetUnbond implements the unbond validator command.
func (s *Service) GetUnbond(amountStr, delegatorAddress, validatorAddr, fees, chainID string, gas uint64, kb keys.Keybase) error {
	txBldr := auth.NewTxBuilderFromCLIMemKeyBase(fees, chainID, gas, s.ctx.Codec, kb)

	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		return err
	}
	return cli.GetUnbond(s.ctx, txBldr, amountStr, delegatorAddress, validatorAddr)
}

// GetUnbond implements the unbond validator command by SecurityAddress
func (s *Service) GetUnbondBySecAddr(fees, chainID string, gas uint64, args []string, kb keys.Keybase) error {
	txBldr := auth.NewTxBuilderFromCLIMemKeyBase(fees, chainID, gas, s.ctx.Codec, kb)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		return err
	}
	return cli.GetUnbondBySecAddr(s.ctx, txBldr, args)
}

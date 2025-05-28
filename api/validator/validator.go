package validator

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/staking/client"
	"github.com/gatechain/gatechainsdk/gatechain/staking/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

type Service struct {
	ctx *context.NodeVaultQuerierImpl
}

func NewService(ctx *context.NodeVaultQuerierImpl) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) Status() error {
	status, err := s.ctx.Client.Status()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return utils.JsonOutPut(status, s.ctx.GetCodec())
}

func (s *Service) QueryValidator(validatorAddr string) error {
	validatorAccAddr, err := sdk.AccAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	bz, err := s.ctx.GetCodec().MarshalJSON(types.NewQueryValidatorParams(sdk.ValAddress(validatorAccAddr)))
	if err != nil {
		fmt.Println(err)
		return err
	}
	route := fmt.Sprintf("custom/%s/%s", "staking", "validator")
	res, _, err := s.ctx.QueryWithData(route, bz)
	if len(res) == 0 {
		err := fmt.Errorf("No validator found with address %s", validatorAddr)
		fmt.Println(err)
		return err
	}
	validator := types.Validator{}
	s.ctx.GetCodec().MustUnmarshalJSON(res, &validator)

	gmValidator, err := s.ctx.Client.GetConAccount(sdk.GetAddressTO48(validatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return err
	}

	tmpValidator, err := client.GetTmpValidator(validator, gmValidator)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return s.ctx.PrintOutput(tmpValidator)
}

// QueryValidators implements the query all validators command.
// gatecli account list
func (s *Service) QueryValidators(page, limit int) error {

	params := types.NewQueryValidatorsParams(page, limit)
	bz, err := s.ctx.Codec.MarshalJSON(params)
	if err != nil {
		fmt.Println(err)
		return err
	}

	route := fmt.Sprintf("custom/%s/%s", "staking", "validators")
	res, _, err := s.ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var validators2 types.Validators
	s.ctx.GetCodec().MustUnmarshalJSON(res, &validators2)
	var tmpValidators2 client.TmpValidators = make([]client.TmpValidator, 0)
	for _, validator := range validators2 {
		gmValidator, err := s.ctx.Client.GetConAccount(sdk.GetAddressTO48(validator.OperatorAddress))
		if err != nil {
			err := fmt.Errorf("validator gm consensus data query failed:%s", err)
			fmt.Println(err)
			return err
		}
		tmpValidator, err := client.GetTmpValidator(validator, gmValidator)
		if err != nil {
			fmt.Println(err)
			return err
		}
		tmpValidators2 = append(tmpValidators2, *tmpValidator)
	}

	return s.ctx.PrintOutput(tmpValidators2)
}

func (s *Service) CreateValidator(validatorAddress string) error {
	validatorAccAddress, err := sdk.AccAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
		return err
	}

	extra, err := s.ctx.Client.GenParticipationKey(sdk.GetAddressTO48(validatorAccAddress))
	if err != nil {
		fmt.Println(err)
		return err
	}
	if extra.Code != 00 {
		fmt.Println(extra.Log)
		return err
	}

	status, err := s.ctx.Client.GetParticipationKey(sdk.GetAddressTO48(validatorAccAddress))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(status)
	return nil
}

func (s *Service) ShowValidatorKey(validatorAddr string) error {
	validatorAccAddr, err := sdk.AccAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	extra, err := s.ctx.Client.GetParticipationKey(sdk.GetAddressTO48(validatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return err
	}
	if extra.Code != 00 {
		fmt.Println(extra.Log)
		err = fmt.Errorf(extra.Log)
		return err
	}
	fmt.Println("ParticipationKeyResponse Data: ", extra)
	return nil
}

func (s *Service) GetValidatorsKey() error {
	restClient, err := s.ctx.GetNode()
	if err != nil {
		fmt.Println(err)
		return err
	}
	accounts, err := restClient.ListLocalConAccounts()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, acc := range accounts {
		if err := s.ctx.PrintOutput(acc); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

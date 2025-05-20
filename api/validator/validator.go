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

func (s *Service) Status() {
	status, err := s.ctx.Client.Status()
	if err != nil {
		fmt.Println(err)
	}
	utils.JsonOutPut(status, s.ctx.GetCodec())
}

func (s *Service) QueryValidator(validatorAddr string) {
	validatorAccAddr, err := sdk.AccAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	bz, err := s.ctx.GetCodec().MarshalJSON(types.NewQueryValidatorParams(sdk.ValAddress(validatorAccAddr)))
	if err != nil {
		fmt.Println(err)
		return
	}
	route := fmt.Sprintf("custom/%s/%s", "staking", "validator")
	res, _, err := s.ctx.QueryWithData(route, bz)
	if len(res) == 0 {
		fmt.Println(fmt.Errorf("No validator found with address %s", validatorAddr))
		return
	}
	validator := types.Validator{}
	s.ctx.GetCodec().MustUnmarshalJSON(res, &validator)

	gmValidator, err := s.ctx.Client.GetConAccount(sdk.GetAddressTO48(validatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpValidator, err := client.GetTmpValidator(validator, gmValidator)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.ctx.PrintOutput(tmpValidator)
}

// QueryValidators implements the query all validators command.
// gatecli account list
func (s *Service) QueryValidators(page, limit int) {

	params := types.NewQueryValidatorsParams(page, limit)
	bz, err := s.ctx.Codec.MarshalJSON(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	route := fmt.Sprintf("custom/%s/%s", "staking", "validators")
	res, _, err := s.ctx.QueryWithData(route, bz)

	var validators2 types.Validators
	s.ctx.GetCodec().MustUnmarshalJSON(res, &validators2)
	var tmpValidators2 client.TmpValidators = make([]client.TmpValidator, 0)
	for _, validator := range validators2 {
		gmValidator, err := s.ctx.Client.GetConAccount(sdk.GetAddressTO48(validator.OperatorAddress))
		if err != nil {
			fmt.Println(fmt.Sprintf("validator gm consensus data query failed:%s", err))
			return
		}
		tmpValidator, err := client.GetTmpValidator(validator, gmValidator)
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpValidators2 = append(tmpValidators2, *tmpValidator)
	}

	s.ctx.PrintOutput(tmpValidators2)
}

func (s *Service) CreateValidator(validatorAddress string) {
	validatorAccAddress, err := sdk.AccAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	extra, err := s.ctx.Client.GenParticipationKey(sdk.GetAddressTO48(validatorAccAddress))
	if err != nil {
		fmt.Println(err)
		return
	}
	if extra.Code != 00 {
		fmt.Println(extra.Log)
		return
	}

	status, err := s.ctx.Client.GetParticipationKey(sdk.GetAddressTO48(validatorAccAddress))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(status)
}

func (s *Service) ShowValidatorKey(validatorAddr string) {
	validatorAccAddr, err := sdk.AccAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	extra, err := s.ctx.Client.GetParticipationKey(sdk.GetAddressTO48(validatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return
	}
	if extra.Code != 00 {
		fmt.Println(extra.Log)
		return
	}
	fmt.Println("ParticipationKeyResponse Data: ", extra)
}

func (s *Service) GetValidatorsKey() {
	restClient, err := s.ctx.GetNode()
	if err != nil {
		fmt.Println(err)
		return
	}
	accounts, err := restClient.ListLocalConAccounts()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, acc := range accounts {
		if err := s.ctx.PrintOutput(acc); err != nil {
			fmt.Println(err)
			return
		}
	}
}

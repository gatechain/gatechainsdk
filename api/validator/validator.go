package validator

import (
	"fmt"
	"github.com/gatechain/gatechainsdk/gatechain/node/appinterface"
	v1 "github.com/gatechain/gatechainsdk/gatechain/rpc/spec/v1"

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

func (s *Service) QueryValidator(validatorAddr string) (*client.TmpValidator, error) {
	validatorAccAddr, err := sdk.AccAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return &client.TmpValidator{}, err
	}

	bz, err := s.ctx.GetCodec().MarshalJSON(types.NewQueryValidatorParams(sdk.ValAddress(validatorAccAddr)))
	if err != nil {
		fmt.Println(err)
		return &client.TmpValidator{}, err
	}
	route := fmt.Sprintf("custom/%s/%s", "staking", "validator")
	res, _, err := s.ctx.QueryWithData(route, bz)
	if len(res) == 0 {
		err := fmt.Errorf("No validator found with address %s", validatorAddr)
		fmt.Println(err)
		return &client.TmpValidator{}, err
	}
	validator := types.Validator{}
	s.ctx.GetCodec().MustUnmarshalJSON(res, &validator)

	gmValidator, err := s.ctx.Client.GetConAccount(sdk.GetAddressTO48(validatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return &client.TmpValidator{}, err
	}

	tmpValidator, err := client.GetTmpValidator(validator, gmValidator)
	if err != nil {
		fmt.Println(err)
		return &client.TmpValidator{}, err
	}
	s.ctx.PrintOutput(tmpValidator)
	return tmpValidator, nil
}

// QueryValidators implements the query all validators command.
// gatecli account list
func (s *Service) QueryValidators(page, limit int) (*client.TmpValidators, error) {

	params := types.NewQueryValidatorsParams(page, limit)
	bz, err := s.ctx.Codec.MarshalJSON(params)
	if err != nil {
		fmt.Println(err)
		return &client.TmpValidators{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", "staking", "validators")
	res, _, err := s.ctx.QueryWithData(route, bz)
	if err != nil {
		fmt.Println(err)
		return &client.TmpValidators{}, err
	}

	var validators2 types.Validators
	s.ctx.GetCodec().MustUnmarshalJSON(res, &validators2)
	var tmpValidators2 client.TmpValidators = make([]client.TmpValidator, 0)
	for _, validator := range validators2 {
		gmValidator, err := s.ctx.Client.GetConAccount(sdk.GetAddressTO48(validator.OperatorAddress))
		if err != nil {
			err := fmt.Errorf("validator gm consensus data query failed:%s", err)
			fmt.Println(err)
			return &client.TmpValidators{}, err
		}
		tmpValidator, err := client.GetTmpValidator(validator, gmValidator)
		if err != nil {
			fmt.Println(err)
			return &client.TmpValidators{}, err
		}
		tmpValidators2 = append(tmpValidators2, *tmpValidator)
	}
	s.ctx.PrintOutput(tmpValidators2)
	return &tmpValidators2, nil
}

func (s *Service) ShowValidatorKey(validatorAddr string) (response v1.ParticipationKeyResponse, err error) {
	validatorAccAddr, err := sdk.AccAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	extra, err := s.ctx.Client.GetParticipationKey(sdk.GetAddressTO48(validatorAccAddr))
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	if extra.Code != 00 {
		fmt.Println(extra.Log)
		err = fmt.Errorf(extra.Log)
		return response, err
	}
	fmt.Println("ParticipationKeyResponse Data: ", extra)
	return extra, nil
}

func (s *Service) GetValidatorsKey() (response []appinterface.ParticipationData, err error) {
	restClient, err := s.ctx.GetNode()
	if err != nil {
		fmt.Println(err)
		return []appinterface.ParticipationData{}, err
	}
	accounts, err := restClient.ListLocalConAccounts()
	if err != nil {
		fmt.Println(err)
		return []appinterface.ParticipationData{}, err
	}

	for _, acc := range accounts {
		if err := s.ctx.PrintOutput(acc); err != nil {
			fmt.Println(err)
			return []appinterface.ParticipationData{}, err
		}
	}
	return accounts, nil
}

package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestGetQueryValidator(t *testing.T) {
	validatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Validator.QueryValidator(validatorAddr)
}

func TestGetCmdQueryValidators(t *testing.T) {
	page := 1
	limit := 100
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Validator.QueryValidators(page, limit)
}

func TestGetCmdShowValidatorKey(t *testing.T) {
	validatorAccAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Validator.ShowValidatorKey(validatorAccAddr)
}

func TestGetValidatorsKey(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Validator.GetValidatorsKey()
}

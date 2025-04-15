package validator

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"

	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	client2 "github.com/gatechain/gatechainsdk/gatechain/staking/client"
	staking_types "github.com/gatechain/gatechainsdk/gatechain/staking/types"
)

func TestCreateValidator(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)
	validatorAddress := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f"
	CreateValidator(ctx, validatorAddress)
}

func TestGetQueryValidator(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)
	validatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	GetQueryValidator(ctx, validatorAddr)
}

func TestGetCmdQueryValidators(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)
	page := 1
	limit := 100

	GetCmdQueryValidators(ctx, page, limit)
}

// done
// gatecli con-account show-key gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
func TestGetCmdShowValidatorKey(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)
	validatorAccAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	GetCmdShowValidatorKey(ctx, validatorAccAddr)
}

// gatecli con-account list-key
func TestGetValidatorsKey(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)
	GetValidatorsKey(ctx)
}

// done
// GetCmdCreateValidator implements the create validator command handler.
// gatecli con-account online --from gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4
// --fees 1000000NANOGT --chain-id gate-66
func TestCmdOnlineValidator(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)

	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}

	validatorAccAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	OnlineValidator(ctx, txBldr, validatorAccAddr)
}

// done
// GetCmdCreateValidator implements the create validator command handler.
// gatecli con-account offline --from gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4
// --fees 1000000NANOGT --chain-id gate-66
func TestCmdOfflineValidator(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)

	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}

	validatorAccAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	OfflineValidator(ctx, txBldr, validatorAccAddr)
}

// todo
// GetCmdEditValidator implements the create edit validator command.
// TODO: add full description
func TestGetCmdEditValidator(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)

	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}
	validatorAccAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	description := staking_types.Description{
		Moniker:  viper.GetString(client2.FlagMoniker),
		Identity: viper.GetString(client2.FlagIdentity),
		Website:  viper.GetString(client2.FlagWebsite),
		Details:  viper.GetString(client2.FlagDetails),
	}
	commissionRate := ""
	GetCmdEditValidator(ctx, txBldr, validatorAccAddr, commissionRate, description)
}

// todo
func TestGetCmdEditValidatorMaxRate(t *testing.T) {
	maxRate := ""
	maxChangeRate := ""

	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)

	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}
	validatorAccAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"

	GetCmdEditValidatorMaxRate(ctx, txBldr, validatorAccAddr, maxRate, maxChangeRate)
}

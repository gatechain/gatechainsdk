package validator

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/node/gen"
	"github.com/gatechain/gatechainsdk/gatechain/staking/client"
	types2 "github.com/gatechain/gatechainsdk/gatechain/staking/types"
	types3 "github.com/gatechain/gatechainsdk/gatechain/types"
)

// gatecli con-account show gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
func GetQueryValidator(ctx *context.NodeVaultQuerierImpl, validatorAddr string) {
	validatorAccAddr, err := types3.AccAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	bz, err := ctx.GetCodec().MarshalJSON(types2.NewQueryValidatorParams(types3.ValAddress(validatorAccAddr)))
	if err != nil {
		fmt.Println(err)
	}
	route := fmt.Sprintf("custom/%s/%s", "staking", "validator")
	res, _, err := ctx.QueryWithData(route, bz)
	if len(res) == 0 {
		fmt.Println(fmt.Errorf("No validator found with address %s", validatorAddr))
	}
	validator := types2.Validator{}
	ctx.GetCodec().MustUnmarshalJSON(res, &validator)

	gmValidator, err := ctx.Client.GetConAccount(types3.GetAddressTO48(validatorAccAddr))
	if err != nil {
		fmt.Println(err)
	}

	tmpValidator, err := client.GetTmpValidator(validator, gmValidator)
	if err != nil {
		fmt.Println(err)
	}
	ctx.PrintOutput(tmpValidator)
}

// GetCmdQueryValidators implements the query all validators command.
// gatecli account list
func GetCmdQueryValidators(ctx *context.NodeVaultQuerierImpl, page, limit int) {

	params := types2.NewQueryValidatorsParams(page, limit)
	bz, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		fmt.Println(err)
	}

	route := fmt.Sprintf("custom/%s/%s", "staking", "validators")
	res, _, err := ctx.QueryWithData(route, bz)

	var validators2 types2.Validators
	ctx.GetCodec().MustUnmarshalJSON(res, &validators2)
	var tmpValidators2 client.TmpValidators = make([]client.TmpValidator, 0)
	for _, validator := range validators2 {
		gmValidator, err := ctx.Client.GetConAccount(types3.GetAddressTO48(validator.OperatorAddress))
		if err != nil {
			fmt.Println(fmt.Sprintf("validator gm consensus data query failed:%s", err))
		}
		tmpValidator, err := client.GetTmpValidator(validator, gmValidator)
		if err != nil {
			fmt.Println(err)
		}
		tmpValidators2 = append(tmpValidators2, *tmpValidator)
	}

	ctx.PrintOutput(tmpValidators2)
}

func CreateValidator(node *context.NodeVaultQuerierImpl, validatorAddress string) {
	validatorAccAddress, err := sdk.AccAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
	}

	extra, err := node.Client.GenParticipationKey(types3.GetAddressTO48(validatorAccAddress))
	if err != nil {
		fmt.Println(err)
	}
	if extra.Code != 00 {
		fmt.Println(extra.Log)
	}

	status, err := node.Client.GetParticipationKey(types3.GetAddressTO48(validatorAccAddress))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(status)
}

func GetCmdShowValidatorKey(node *context.NodeVaultQuerierImpl, validatorAddr string) {
	validatorAccAddr, err := types3.AccAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	extra, err := node.Client.GetParticipationKey(types3.GetAddressTO48(validatorAccAddr))
	if err != nil {
		fmt.Println(err)
	}
	if extra.Code != 00 {
		fmt.Println(extra.Log)
	}
	fmt.Println("ParticipationKeyResponse Data: ", extra)
}

// gatecli con-account list-key
func GetValidatorsKey(ctx *context.NodeVaultQuerierImpl) {
	restClient, err := ctx.GetNode()
	if err != nil {
		fmt.Println(err)
	}
	accounts, err := restClient.ListLocalConAccounts()
	if err != nil {
		fmt.Println(err)
	}

	for _, acc := range accounts {
		if err := ctx.PrintOutput(acc); err != nil {
			fmt.Println(err)
		}
	}
}

// done
// GetCmdCreateValidator implements the create validator command handler.
// gatecli con-account online --from gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4
// --fees 1000000NANOGT --chain-id gate-66
func OnlineValidator(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, validatorAddress string) {
	validatorAccAddress, err := types3.AccAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
	}
	res, _, err := ctx.QueryStore(types2.GetValidatorKey(types3.ValAddress(validatorAccAddress)), "staking")
	if err != nil {
		fmt.Println(err)
	}

	if len(res) == 0 {
		txBldr, msg, err := client.BuildCreateValidatorMsg(*ctx, txBldr)
		if err != nil {
			fmt.Println(err)
		}
		txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
		fmt.Println(string(txbytes))
	} else {
		node, err := ctx.GetNode()
		if err != nil {
			fmt.Println(err)
		}
		data, err := node.GetParticipationKey(types3.GetAddressTO48(validatorAccAddress))
		if err != nil {
			fmt.Println(err)
		}
		if data.Code != 00 {
			fmt.Println(data.Log)
		}
		extra := gen.MakeOnlineExtra(data.Data)
		msg := types2.NewMsgValidatorSwitchState(types3.ValAddress(validatorAccAddress), extra)
		txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
		fmt.Println(txbytes)
	}
}

// done
// GetCmdCreateValidator implements the create validator command handler.
// gatecli con-account offline --from gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4
// --fees 1000000NANOGT --chain-id gate-66
func OfflineValidator(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, validatorAddress string) {
	validatorAccAddress, err := types3.AccAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
	}

	res, _, err := ctx.QueryStore(types2.GetValidatorKey(types3.ValAddress(validatorAccAddress)), "staking")
	if err != nil {
		fmt.Println(err)
	}

	if len(res) == 0 {
		fmt.Println("con-account not found (online)")
	} else {
		node, err := ctx.GetNode()
		if err != nil {
			fmt.Println(err)
		}
		data, err := node.GetParticipationKey(types3.GetAddressTO48(validatorAccAddress))
		if err != nil {
			fmt.Println(err)
		}
		if data.Code != 00 {
			fmt.Println(data.Log)
		}
		extra := gen.MakeOfflineExtra()
		msg := types2.NewMsgValidatorSwitchState(types3.ValAddress(validatorAccAddress), extra)
		txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
		fmt.Println(txbytes)
	}
}

// todo
// GetCmdEditValidator implements the create edit validator command.
// TODO: add full description
func GetCmdEditValidator(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, validatorAddress,
	commissionRate string, description types2.Description) {
	var newRate *types3.Dec
	if commissionRate != "" {
		rate, err := types3.NewDecFromStr(commissionRate)
		if err != nil {
			fmt.Println(fmt.Sprintf("invalid new commission rate: %v", err))
		}

		newRate = &rate
	}
	validatorAccAddress, err := types3.AccAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
	}
	msg := types2.NewMsgEditValidator(types3.ValAddress(validatorAccAddress), description, newRate)

	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
	fmt.Println(string(txbytes))
}

// todo
func GetCmdEditValidatorMaxRate(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, validatorAddress, maxRate, maxChangeRate string) {

	var newMaxRate types3.Dec
	var newMaxChangeRate types3.Dec
	var err error
	if maxRate == "" || maxChangeRate == "" {
		fmt.Println(fmt.Sprintf("nil commission max rate or commission max change rate"))
	}

	newMaxRate, err = types3.NewDecFromStr(maxRate)
	if err != nil {
		fmt.Println(fmt.Sprintf("invalid new commission max rate: %v", err))
	}

	newMaxChangeRate, err = types3.NewDecFromStr(maxChangeRate)
	if err != nil {
		fmt.Println(fmt.Sprintf("invalid new commission max change rate: %v", err))
	}

	validatorAccAddress, err := types3.AccAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
	}
	msg := types2.NewMsgEditValidatorMaxRate(types3.ValAddress(validatorAccAddress), &newMaxRate, &newMaxChangeRate)

	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
	fmt.Println(string(txbytes))
}

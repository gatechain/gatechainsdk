// nolint
package cli

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/client/common"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// command to withdraw rewards
func WithdrawRewards(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, delegatorAddress, validatorAddress string, bComission bool) {
	delegatorAccAddress, _, err := sdk.AccAddressTypeFromBech32(delegatorAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	validatorAccAddress, err := sdk.ValAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	msgs := make([]sdk.Msg, 0)
	if bComission {
		msgs = append(msgs, types.NewMsgWithdrawValidatorCommission(validatorAccAddress))
	} else {
		msgs = append(msgs, types.NewMsgWithdrawDelegatorReward(delegatorAccAddress, validatorAccAddress))
	}
	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, ctx, msgs, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txbytes)
}

// command to withdraw all rewards
func WithdrawAllRewards(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, delegatorAddr string) {
	delegatorAccAddr, _, err := sdk.AccAddressTypeFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	msgs, err := common.WithdrawAllDelegatorRewards(*ctx, types.ModuleName, delegatorAccAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, ctx, msgs, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txbytes)
}

// command to replace a delegator's withdrawal address
func SetWithdrawAddr(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, delegatorAddress, withdrawAddress string) {
	delegatorAccAddr, _, err := sdk.AccAddressTypeFromBech32(delegatorAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	withdrawAddr, err := sdk.AccAddressFromBech32(withdrawAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := types.NewMsgSetWithdrawAddress(delegatorAccAddr, withdrawAddr)
	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, ctx, []sdk.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txbytes)
}

func RewardReinvestment(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, delegatorAddress, validatorAddress string) {

	delAccAddr, _, err := sdk.AccAddressTypeFromBech32(delegatorAddress)

	valAccAddr, err := sdk.ValAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	msgs := make([]sdk.Msg, 0)
	msgs = append(msgs, types.NewMsgRewardReinvestment(delAccAddr, valAccAddr))

	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, ctx, msgs, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(txbytes)
}

// nolint
package cli

import (
	"fmt"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/distribution/client/common"
	types3 "github.com/gatechain/gatechainsdk/gatechain/distribution/types"
	"github.com/gatechain/gatechainsdk/gatechain/types"
	// "github.com/gatechain/gatechain/framework/x/gov"
)

var (
	flagOnlyFromValidator = "only-from-validator"
	flagIsValidator       = "is-validator"
	flagComission         = "commission"
	flagMaxMessagesPerTx  = "max-msgs"
)

const (
	MaxMessagesPerTxDefault = 5
)

// command to withdraw rewards
func GetCmdWithdrawRewards(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, delegatorAddress, validatorAddress string, bComission bool) {
	delegatorAccAddress, _, err := types.AccAddressTypeFromBech32(delegatorAddress)
	if err != nil {
		fmt.Println(err)
	}
	validatorAccAddress, err := types.ValAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
	}
	msgs := make([]types.Msg, 0)
	if bComission {
		msgs = append(msgs, types3.NewMsgWithdrawValidatorCommission(validatorAccAddress))
	} else {
		msgs = append(msgs, types3.NewMsgWithdrawDelegatorReward(delegatorAccAddress, validatorAccAddress))
	}
	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, msgs, true)
	fmt.Println(txbytes, err)
}

// command to withdraw all rewards
func GetGetCmdWithdrawAllRewards(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, queryRoute, delegatorAddr string) {
	delegatorAccAddr, _, err := types.AccAddressTypeFromBech32(delegatorAddr)
	if err != nil {
		fmt.Println(err)
	}
	msgs, err := common.WithdrawAllDelegatorRewards(*ctx, queryRoute, delegatorAccAddr)
	if err != nil {
		fmt.Println(err)
	}
	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, msgs, true)
	fmt.Println(txbytes, err)
}

// command to replace a delegator's withdrawal address
func SetWithdrawAddr(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, delegatorAddress, withdrawAddress string) {
	delegatorAccAddr, _, err := types.AccAddressTypeFromBech32(delegatorAddress)
	if err != nil {
		fmt.Println(err)
	}
	withdrawAddr, err := types.AccAddressFromBech32(withdrawAddress)
	if err != nil {
		fmt.Println(err)
	}
	msg := types3.NewMsgSetWithdrawAddress(delegatorAccAddr, withdrawAddr)
	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types.Msg{msg}, true)
	fmt.Println(txbytes, err)
}

// DelegatorAddress: delAddr,
//
//	validatorAddress: valAddr,
func GetCmdRewardReinvestment(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, delegatorAddress, validatorAddress string) {

	delAccAddr, _, err := types.AccAddressTypeFromBech32(delegatorAddress)

	valAccAddr, err := types.ValAddressFromBech32(validatorAddress)
	if err != nil {
		fmt.Println(err)
	}
	msgs := make([]types.Msg, 0)
	msgs = append(msgs, types3.NewMsgRewardReinvestment(delAccAddr, valAccAddr))

	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, msgs, true)
	fmt.Println(txbytes, err)
}

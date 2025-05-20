package cli

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/staking/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// GetDelegate implements the delegate command.
func GetDelegate(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, amountStr, delegatorAddress, validatorAddr string) {
	amount, err := sdk.ParseCoin(amountStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	delegatorAccAddress, err := sdk.AccAddressFromBech32(delegatorAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	validatorValAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := types.NewMsgDelegate(delegatorAccAddress, validatorValAddr, amount)
	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, ctx, []sdk.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txbytes)
}

// GetRedelegate the begin redelegation command.
func GetRedelegate(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, delAddr string, args []string) {
	delegatorAccAddress, _, err := sdk.AccAddressTypeFromBech32(delAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	valSrcAddr, err := sdk.ValAddressFromBech32(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	valDstAddr, err := sdk.ValAddressFromBech32(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	amount, err := sdk.ParseCoin(args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := types.NewMsgBeginRedelegate(delegatorAccAddress, valSrcAddr, valDstAddr, amount)

	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, ctx, []sdk.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txbytes)
}

// GetUnbond implements the unbond validator command.
func GetUnbond(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, amountStr, delegatorAddress, validatorAddr string) {
	delAddr, _, err := sdk.AccAddressTypeFromBech32(delegatorAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	validatorValAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	amount, err := sdk.ParseCoin(amountStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := types.NewMsgUndelegate(delAddr, validatorValAddr, amount)
	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, ctx, []sdk.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txbytes)
}

// GetUnbond implements the unbond validator command by SecurityAddress
func GetUnbondBySecAddr(ctx *context.NodeVaultQuerierImpl, txBldr auth.TxBuilder, args []string) {
	var vaultAddress []sdk.AccAddress
	from := ctx.GetFromAddress()
	for i := 0; i < len(args); i++ {
		address, _, err := sdk.AccAddressTypeFromBech32(args[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		vaultAddress = append(vaultAddress, address)
	}
	msg := types.NewMsgUndelegateByRetrievalAccount(from, vaultAddress)
	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, ctx, []sdk.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txbytes)
}

//__________________________________________________________

package cli

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	types2 "github.com/gatechain/gatechainsdk/gatechain/staking/types"
	types3 "github.com/gatechain/gatechainsdk/gatechain/types"
)

// GetCmdDelegate implements the delegate command.
func GetCmdDelegate(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, amountStr, delegatorAddress, validatorAddr string) {
	amount, err := types3.ParseCoin(amountStr)
	if err != nil {
		fmt.Println(err)
	}
	delegatorAccAddress, err := types3.AccAddressFromBech32(delegatorAddress)
	if err != nil {
		fmt.Println(err)
	}

	validatorValAddr, err := types3.ValAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	msg := types2.NewMsgDelegate(delegatorAccAddress, validatorValAddr, amount)
	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(txbytes)
}

// GetCmdRedelegate the begin redelegation command.
func GetCmdRedelegate(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, delAddr string, args []string) {
	delegatorAccAddress, _, err := types3.AccAddressTypeFromBech32(delAddr)
	if err != nil {
		fmt.Println(err)
	}

	valSrcAddr, err := types3.ValAddressFromBech32(args[0])
	if err != nil {
		fmt.Println(err)
	}

	valDstAddr, err := types3.ValAddressFromBech32(args[1])
	if err != nil {
		fmt.Println(err)
	}

	amount, err := types3.ParseCoin(args[2])
	if err != nil {
		fmt.Println(err)
	}

	msg := types2.NewMsgBeginRedelegate(delegatorAccAddress, valSrcAddr, valDstAddr, amount)

	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(txbytes)
}

// GetCmdUnbond implements the unbond validator command.
func GetCmdUnbond(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, amountStr, delegatorAddress, validatorAddr string) {
	delAddr, _, err := types3.AccAddressTypeFromBech32(delegatorAddress)
	if err != nil {
		fmt.Println(err)
	}
	validatorValAddr, err := types3.ValAddressFromBech32(validatorAddr)
	if err != nil {
		fmt.Println(err)
	}

	amount, err := types3.ParseCoin(amountStr)
	if err != nil {
		fmt.Println(err)
	}

	msg := types2.NewMsgUndelegate(delAddr, validatorValAddr, amount)
	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(txbytes)
}

// GetCmdUnbond implements the unbond validator command by SecurityAddress
func GetCmdUnbondBySecAddr(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, args []string) {
	var vaultAddress []types3.AccAddress
	from := ctx.GetFromAddress()
	for i := 0; i < len(args); i++ {
		address, _, err := types3.AccAddressTypeFromBech32(args[i])
		if err != nil {
			fmt.Println(err)
		}
		vaultAddress = append(vaultAddress, address)
	}
	msg := types2.NewMsgUndelegateByRetrievalAccount(from, vaultAddress)
	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types3.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(txbytes)
}

//__________________________________________________________

var (
	defaultTokens                  = types3.TokensFromConsensusPower(100)
	defaultAmount                  = defaultTokens.String() + sdk.DefaultBondDenom
	defaultCommissionRate          = "0.1"
	defaultCommissionMaxRate       = "0.2"
	defaultCommissionMaxChangeRate = "0.01"
)

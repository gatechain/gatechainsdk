package vaultAccount

import (
	"fmt"
	"strconv"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/revocable/client"
	"github.com/gatechain/gatechainsdk/gatechain/revocable/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

type Service struct {
	ctx *context.NodeVaultQuerierImpl
}

func NewService(ctx *context.NodeVaultQuerierImpl) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) BroadcastMsgCreateVault(from_addr, to_addr, security_addr, delayHeightStr, clearTimeHeightStr,
	coinsStr, pubkeyStr, fees, chainID string, gas uint64) error {

	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fromAccAddress, _, err := sdk.AccAddressTypeFromBech32(from_addr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	toAccAddress, _, err := sdk.AccAddressTypeFromBech32(to_addr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	securityAccAddress, _, err := sdk.AccAddressTypeFromBech32(security_addr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	delayHeight, err := strconv.ParseUint(delayHeightStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		return err
	}

	clearTime, err := strconv.ParseUint(clearTimeHeightStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// parse coins trying to be sent
	coins, err := sdk.ParseCoins(coinsStr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	msg := types.NewMsgCreateVault(fromAccAddress, toAccAddress, securityAccAddress, delayHeight, clearTime, coins, pubkeyStr)
	fmt.Println(msg)

	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, s.ctx, []sdk.Msg{msg}, true)
	if err != nil {
		return err
	}
	fmt.Println(txbytes)
	return nil
}

func (s *Service) BroadcastUpdateClearingHeightTx(clearTimeHeight, vaultAddr, fees, chainID string, gas uint64) error {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	clearTime, err := strconv.ParseUint(clearTimeHeight, 10, 64)
	if err != nil {
		fmt.Println(err)
		return err
	}

	vaultAccAddress, _, err := sdk.AccAddressTypeFromBech32(vaultAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// build and sign the transaction, then broadcast to Tendermint
	msg := types.NewMsgUpdateClearingHeight(vaultAccAddress, clearTime)

	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, s.ctx, []sdk.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(txbytes)
	return nil
}

func (s *Service) QueryVaultAccount(vaultAddr string) error {
	retriever := auth.NewVaultRetriever(s.ctx)
	key, err := sdk.AccAddressFromBech32(vaultAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	account, height, err := retriever.GetAccountWithHeight(key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(account, height)
	return nil
}

func (s *Service) ClearVaultAccountTx(fees, chainID string, vaultAddresses []string, gas uint64) error {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var vaultAddress []sdk.AccAddress
	for i := 0; i < len(vaultAddresses); i++ {
		address, _, err := sdk.AccAddressTypeFromBech32(vaultAddresses[i])
		if err != nil {
			fmt.Println(err)
			return err
		}
		vaultAddress = append(vaultAddress, address)
	}
	if err := client.EnsureFromVaultAccount(*s.ctx); err != nil {
		fmt.Println(err)
		return err
	}

	msg := types.NewMsgClearVaultAccount(s.ctx.FromAddress, vaultAddress)
	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, s.ctx, []sdk.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(txbytes)
	return nil
}

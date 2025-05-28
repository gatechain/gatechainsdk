package revocableTx

import (
	"encoding/hex"
	"errors"
	"fmt"

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

func (s *Service) QueryTx(hashHexStr string) error {
	res, err := auth.QueryTx(s.ctx, hashHexStr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.Empty() {
		err := fmt.Errorf("No transaction found with hash %s", hashHexStr)
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	return nil
}

func (s *Service) GetAccountRevocableTx(vaultAddr string) error {
	retriever := auth.NewVaultRetriever(s.ctx)
	key, err := sdk.AccAddressFromBech32(vaultAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	account, height, err := retriever.GetAccountWithHeight(key)
	fmt.Println(account, height, err)

	key, addressType, err := sdk.AccAddressTypeFromBech32(vaultAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if addressType != sdk.VaultAccount_type && addressType != sdk.MultiSignerVaultAccount {
		err := errors.New("input account must be vault account")
		fmt.Println(err)
		return err
	}

	if err := retriever.EnsureExists(key); err != nil {
		fmt.Println(err)
		return err
	}

	height, err = s.ctx.GetChainHeight()
	if err != nil {
		fmt.Println(err)
		return err
	}

	vault, err := retriever.GetAccount(key)
	if err != nil {
		fmt.Println(err)
		return err
	}

	data, err := vault.GetRevocableTokensDetail(height)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(data)
	return nil
}

func (s *Service) RevocableTxSend(from_addr, to_addr, amount, fees, chainID string, gas uint64) error {
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

	// parse coins trying to be sent
	coins, err := sdk.ParseCoins(amount)
	if err != nil {
		fmt.Println(err)
		return err
	}

	msg := types.NewMsgRevocableSend(fromAccAddress, toAccAddress, coins)

	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, s.ctx, []sdk.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(txbytes)
	return nil
}

func (s *Service) RevokeTx(fromAddr, txHash, fees, chainID string, gas uint64) error {

	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fromAccAddress, err := sdk.AccAddressFromBech32(fromAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	s.ctx.WithFromAddress(fromAccAddress)
	if err := client.EnsureFromVaultAccount(*s.ctx); err != nil {
		fmt.Println(err)
		return err
	}

	txHash, err = auth.SplitTxHashPreFix(txHash)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// get tx detail
	tx, err := auth.QueryTx(s.ctx, txHash)
	if err != nil {
		fmt.Println(err)
		return err
	}

	vaultAccount, err := client.GetVaultAccount(*s.ctx, fromAccAddress)
	if err != nil {
		fmt.Println(err)
		return err
	}

	height, err := s.ctx.GetChainHeight()
	if err != nil {
		fmt.Println(err)
		return err
	}

	data, err := client.GetRevocableTokens(*s.ctx, fromAccAddress, height)
	if err != nil {
		fmt.Println(err)
		return err
	}

	index := uint64(0)

	delayHeight := uint64(0)
	coins := sdk.Coins{}
	for _, delay := range data {
		delayTxHash := delay.TxHash
		// TODO: fix me
		if delayTxHash == txHash {
			if index != delay.Index {
				err := errors.New("index for revoke tx input error")
				fmt.Println(err)
				return err
			}
			delayHeight = delay.Height
			coins = coins.Add(delay.Coins)
			break
		}
	}

	if delayHeight == uint64(0) {
		err := fmt.Errorf("can not find revoke tx delayHeight : " + txHash)
		return err
	}

	toAddress := sdk.AccAddress{}
	if uint64(len(tx.Tx.GetMsgs())) < index+1 {
		err := fmt.Errorf("can not find revoke tx index: " + txHash)
		return err
	}
	txMsg := tx.Tx.GetMsgs()[index]
	switch txMsg := txMsg.(type) {
	case types.MsgRevocableSend:
		toAddress = txMsg.ToAddress
	default:
		err := fmt.Errorf("can not find revoke tx MsgRevoke: " + txHash)
		return err
	}
	// build and sign the transaction, then broadcast to Tendermint
	msg := types.NewMsgRevoke(fromAccAddress, vaultAccount.GetSecurityAddress().Address, toAddress, int64(delayHeight), int64(index), txHash, coins)
	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, s.ctx, []sdk.Msg{msg}, true)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(txbytes)
	return nil
}

func (s *Service) TxStatus(hashHexStr string) error {

	revGetter := types.NewRevocableRetriever(s.ctx)
	txData, err := auth.QueryTx(s.ctx, hashHexStr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if txData.Empty() {
		err := fmt.Errorf("No transaction found with hash %s", hashHexStr)
		return err
	}
	hashStr, err := auth.SplitTxHashPreFix(hashHexStr)
	if err != nil {
		fmt.Println(err)
		return err

	}
	hash, err := hex.DecodeString(hashStr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	output := &sdk.RevocableTxResponse{}
	revocabletx, err := revGetter.GetTx(hash)
	if err != nil {
		output = sdk.NewRevocableTxResponse(sdk.IRREVOCABLEPAY, "")
	} else if revocabletx.RevokeHeight == uint64(0) {

		sender := auth.GetRevocableTxSender(txData)
		vaultGetter := auth.NewVaultRetriever(s.ctx)

		key, _, err := sdk.AccAddressTypeFromBech32(sender)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if err := vaultGetter.EnsureExists(key); err != nil {
			fmt.Println(err)
			return err
		}

		vault, height, err := vaultGetter.GetAccountWithHeight(key)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if height > txData.Height+int64(vault.GetDelayHeight()) {
			output = sdk.NewRevocableTxResponse(sdk.IRREVOCABLEPAY, "")
		} else {
			output = sdk.NewRevocableTxResponse(sdk.REVOCABLEPAY, "")
		}
	} else {
		output = sdk.NewRevocableTxResponse(sdk.REVOKED, "REVOKE-"+hex.EncodeToString(revocabletx.RevokeTxHash))
	}
	err = s.ctx.PrintOutput(output)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

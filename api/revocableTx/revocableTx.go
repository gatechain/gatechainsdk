package revocableTx

import (
	"encoding/hex"
	"errors"
	"fmt"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	client2 "github.com/gatechain/gatechainsdk/gatechain/revocable/client"
	types3 "github.com/gatechain/gatechainsdk/gatechain/revocable/types"
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
)

func QueryTxCmd(ctx *context.NodeVaultQuerierImpl, hashHexStr string) {
	res, err := auth2.QueryTx(ctx, hashHexStr)
	if err != nil {
		fmt.Println(err)
	}
	if res.Empty() {
		fmt.Println(fmt.Errorf("No transaction found with hash %s", hashHexStr))
	}
	fmt.Println(res)
}

func GetAccountRevocableTxCmd(ctx *context.NodeVaultQuerierImpl, vaultAddr string) {
	retriever := auth2.NewVaultRetriever(ctx)
	key, err := types2.AccAddressFromBech32(vaultAddr)
	if err != nil {
		fmt.Println(err)
	}
	account, height, err := retriever.GetAccountWithHeight(key)
	fmt.Println(account, height, err)

	key, addressType, err := types2.AccAddressTypeFromBech32(vaultAddr)
	if err != nil {
		fmt.Println(err)
	}

	if addressType != types2.VaultAccount_type && addressType != types2.MultiSignerVaultAccount {
		fmt.Println(errors.New("input account must be vault account"))
	}

	if err := retriever.EnsureExists(key); err != nil {
		fmt.Println(err)
	}

	height, err = ctx.GetChainHeight()
	if err != nil {
		fmt.Println(err)
	}

	vault, err := retriever.GetAccount(key)
	if err != nil {
		fmt.Println(err)
	}

	data, err := vault.GetRevocableTokensDetail(height)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

func RevocableTxSendCmd(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, from_addr, to_addr, amount string) {
	fromAccAddress, _, err := types2.AccAddressTypeFromBech32(from_addr)
	if err != nil {
		fmt.Println(err)
	}

	toAccAddress, _, err := types2.AccAddressTypeFromBech32(to_addr)
	if err != nil {
		fmt.Println(err)
	}

	// parse coins trying to be sent
	coins, err := types2.ParseCoins(amount)
	if err != nil {
		fmt.Println(err)
	}

	msg := types3.NewMsgRevocableSend(fromAccAddress, toAccAddress, coins)

	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types2.Msg{msg}, true)
	fmt.Println(txbytes, err)
}

func RevokeTxCmd(ctx *context.NodeVaultQuerierImpl, txBldr auth2.TxBuilder, fromAddr, txHash string) {
	fromAccAddress, err := types2.AccAddressFromBech32(fromAddr)
	if err != nil {
		fmt.Println(err)
	}
	ctx.WithFromAddress(fromAccAddress)
	if err := client2.EnsureFromVaultAccount(*ctx); err != nil {
		fmt.Println(err)
	}

	txHash, err = auth2.SplitTxHashPreFix(txHash)
	if err != nil {
		fmt.Println(err)
	}

	// get tx detail
	tx, err := auth2.QueryTx(ctx, txHash)
	if err != nil {
		fmt.Println(err)
	}

	vaultAccount, err := client2.GetVaultAccount(*ctx, fromAccAddress)
	if err != nil {
		fmt.Println(err)
	}

	height, err := ctx.GetChainHeight()
	if err != nil {
		fmt.Println(err)
	}

	data, err := client2.GetRevocableTokens(*ctx, fromAccAddress, height)
	if err != nil {
		fmt.Println(err)
	}

	index := uint64(0)
	//if len(args) == 2 {
	//	index, err = strconv.ParseUint(args[1], 10, 64)
	//	if err != nil {
	//		return err
	//	}
	//}

	delayHeight := uint64(0)
	coins := types2.Coins{}
	for _, delay := range data {
		delayTxHash := delay.TxHash
		// TODO: fix me
		//delayTxHash, err := sdk.SplitTxHashPreFix(delay.TxHash)
		//if err != nil {
		//	return err
		//}
		if delayTxHash == txHash {
			if index != delay.Index {
				fmt.Println(errors.New("index for revoke tx input error"))
			}
			delayHeight = delay.Height
			coins = coins.Add(delay.Coins)
			break
		}
	}

	if delayHeight == uint64(0) {
		fmt.Println("can not find revoke tx delayHeight : " + txHash)
	}

	toAddress := types2.AccAddress{}
	if uint64(len(tx.Tx.GetMsgs())) < index+1 {
		fmt.Println("can not find revoke tx index: " + txHash)
	}
	txMsg := tx.Tx.GetMsgs()[index]
	switch txMsg := txMsg.(type) {
	case types3.MsgRevocableSend:
		toAddress = txMsg.ToAddress
	default:
		fmt.Println("can not find revoke tx MsgRevoke: " + txHash)
	}
	// build and sign the transaction, then broadcast to Tendermint
	msg := types3.NewMsgRevoke(fromAccAddress, vaultAccount.GetSecurityAddress().Address, toAddress, int64(delayHeight), int64(index), txHash, coins)
	txbytes, err := auth2.CompleteAndBroadcastTxCLI(txBldr, ctx, []types2.Msg{msg}, true)
	fmt.Println(txbytes, err)
}

func TxStatusCmd(hashHexStr string, ctx *context.NodeVaultQuerierImpl) {

	revGetter := types3.NewRevocableRetriever(ctx)
	txData, err := auth2.QueryTx(ctx, hashHexStr)
	if err != nil {
		fmt.Println(err)
	}
	if txData.Empty() {
		fmt.Println(fmt.Errorf("No transaction found with hash %s", hashHexStr))
	}
	hashStr, err := auth2.SplitTxHashPreFix(hashHexStr)
	if err != nil {
		fmt.Println(err)
	}
	hash, err := hex.DecodeString(hashStr)
	if err != nil {
		fmt.Println(err)
	}

	output := &types2.RevocableTxResponse{}
	revocabletx, err := revGetter.GetTx(hash)
	if err != nil {
		output = types2.NewRevocableTxResponse(types2.IRREVOCABLEPAY, "")
	} else if revocabletx.RevokeHeight == uint64(0) {

		sender := auth2.GetRevocableTxSender(txData)
		vaultGetter := auth2.NewVaultRetriever(ctx)

		key, _, err := types2.AccAddressTypeFromBech32(sender)
		if err != nil {
			fmt.Println(err)
		}

		if err := vaultGetter.EnsureExists(key); err != nil {
			fmt.Println(err)
		}

		vault, height, err := vaultGetter.GetAccountWithHeight(key)
		if err != nil {
			fmt.Println(err)
		}

		if height > txData.Height+int64(vault.GetDelayHeight()) {
			output = types2.NewRevocableTxResponse(types2.IRREVOCABLEPAY, "")
		} else {
			output = types2.NewRevocableTxResponse(types2.REVOCABLEPAY, "")
		}
	} else {
		output = types2.NewRevocableTxResponse(types2.REVOKED, "REVOKE-"+hex.EncodeToString(revocabletx.RevokeTxHash))
	}
	err = ctx.PrintOutput(output)
	if err != nil {
		fmt.Println(err)
	}
}

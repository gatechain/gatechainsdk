package auth

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/node/appinterface"
	"github.com/gatechain/gatechainsdk/gatechain/node/basics"
	"github.com/gatechain/gatechainsdk/gatechain/rpc/spec/v1"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

func QueryTx(cliCtx *context.NodeVaultQuerierImpl, hashHexStr string) (types.TxResponse, error) {
	// convert prefix txhash
	hashStr, err := SplitTxHashPreFix(hashHexStr)
	if err != nil {
		return types.TxResponse{}, err
	}

	hash, err := hex.DecodeString(hashStr)
	if err != nil {
		return types.TxResponse{}, err
	}

	node, err := cliCtx.GetNode()
	if err != nil {
		return types.TxResponse{}, err
	}
	resTx, err := node.Tx(hash)
	if err != nil {
		return types.TxResponse{}, err
	}
	if len(resTx.Tx) == 0 {
		return types.TxResponse{}, err
	}
	fmt.Println("resTx", resTx)

	resBlocks, err := getBlocksForTxResults(cliCtx, []*appinterface.ResponseTx{&resTx})
	if err != nil {
		return types.TxResponse{}, err
	}
	fmt.Println("resBlocks", resBlocks)
	out, err := formatTxResult(cliCtx.Codec, &resTx, resBlocks[resTx.Height])
	if err != nil {
		return out, err
	}

	out = ConvertTxHashPrefixForTxResponse(out)
	return out, nil
}

func ReConvertTxResponseFromData(cdc *codec.Codec, txResponse types.TxResponse) types.TxResponse {
	if len(txResponse.Data) == 0 {
		return txResponse
	}
	txBytes, err := base64.StdEncoding.DecodeString(txResponse.Data)
	if err != nil {
		return txResponse
	}

	tx, err := parseTx(cdc, txBytes)
	if err != nil {
		return txResponse
	}
	txResponse.Tx = tx
	return txResponse
}

func getBlocksForTxResults(cliCtx *context.NodeVaultQuerierImpl, resTxs []*appinterface.ResponseTx) (map[int64]*v1.Block, error) {
	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}

	resBlocks := make(map[int64]*v1.Block)

	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := node.Block(uint64(resTx.Height))
			if err != nil {
				return nil, err
			}
			resBlock, err = ChangeGMBlockAddress(resBlock)
			if err != nil {
				return nil, err
			}

			resBlocks[resTx.Height] = &resBlock
		}
	}
	return resBlocks, nil
}

func formatTxResult(cdc *codec.Codec, resTx *appinterface.ResponseTx, resBlock *v1.Block) (types.TxResponse, error) {
	tx, err := parseTx(cdc, resTx.Tx)
	if err != nil {
		return types.TxResponse{}, err
	}

	return types.NewResponseResultTx(resTx, tx, time.Unix(resBlock.Timestamp, 0).Format(time.RFC3339)), nil
}

func parseTx(cdc *codec.Codec, txBytes []byte) (types.Tx, error) {
	var tx StdTx

	err := cdc.UnmarshalBinaryLengthPrefixed(txBytes, &tx)
	if err != nil {
		return nil, err

	}

	return tx, nil
}

func ChangeGMBlockAddress(res v1.Block) (v1.Block, error) {
	proposerGMAddress := basics.Address{}
	if err := proposerGMAddress.UnmarshalText([]byte(res.Proposer)); err != nil {
		return res, err
	}
	res.Proposer = types.AccAddress(types.GetAddressFrom48(proposerGMAddress[:])).String()

	GMaddress := basics.Address{}
	for i := 0; i < len(res.CertCommitteeInfo); i++ {
		if err := GMaddress.UnmarshalText([]byte(res.CertCommitteeInfo[i].Address)); err != nil {
			return res, err
		}
		addr := types.AccAddress(types.GetAddressFrom48(GMaddress[:])).String()
		res.CertCommitteeInfo[i].Address = addr
	}
	for j := 0; j < len(res.BlockCommitteeInfo); j++ {
		if err := GMaddress.UnmarshalText([]byte(res.BlockCommitteeInfo[j].Address)); err != nil {
			return res, err
		}
		addr := types.AccAddress(types.GetAddressFrom48(GMaddress[:])).String()
		res.BlockCommitteeInfo[j].Address = addr
	}

	return res, nil
}

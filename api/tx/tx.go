package tx

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/auth/client/cli"
	"github.com/gatechain/gatechainsdk/gatechain/bank"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

type Service struct {
	ctx *context.NodeVaultQuerierImpl
}

func NewService(ctx *context.NodeVaultQuerierImpl) *Service {
	return &Service{ctx: ctx}
}

// create unsign tx to 	unsignTxFileName
func (s *Service) CreateUnsignTX(from_addr, to_addr, amount, unsignTxFileName string, fees, chainID string, gas uint64) {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
		return
	}
	toAccAddress, _, err := types.AccAddressTypeFromBech32(to_addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fromAccAddress, _, err := types.AccAddressTypeFromBech32(from_addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	coins, err := types.ParseCoins(amount)
	if err != nil {
		fmt.Println(err)
		return
	}
	// build and sign the transaction, then broadcast to Tendermint
	msg := bank.NewMsgSend(fromAccAddress, toAccAddress, coins)
	fmt.Println(msg)
	auth.PrintUnsignedStdTx(txBldr, s.ctx, []types.Msg{msg}, unsignTxFileName)
}

// sign UnsignTxFile to SignTxFile
func (s *Service) CreateSignTX(UnsignTxFile, SignTxFile string, fees, chainID string, gas uint64) {
	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}
	cli.CreateSignTX(UnsignTxFile, SignTxFile, s.ctx, txBldr)
}

// BroadcastSignTx SignedFileName
func (s *Service) BroadcastSignTx(SignedFileName string) {
	cli.BroadcastSignTx(s.ctx, SignedFileName)
}

// create  and send tx
func (s *Service) SendTX(from_addr, to_addr, amount, fees, chainID string, gas uint64) {

	txBldr := auth.NewTxBuilderFromCLI(s.ctx.RootDir, fees, chainID, gas, s.ctx.Codec)
	txBldr, err := auth.UpdateValidHeight(s.ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}

	toAccAddress, _, err := types.AccAddressTypeFromBech32(to_addr)
	if err != nil {
		fmt.Println(err)
	}
	fromAccAddress, _, err := types.AccAddressTypeFromBech32(from_addr)
	if err != nil {
		fmt.Println(err)
	}
	coins, err := types.ParseCoins(amount)
	if err != nil {
		fmt.Println(err)
	}
	// build and sign the transaction, then broadcast to Tendermint
	msg := bank.NewMsgSend(fromAccAddress, toAccAddress, coins)
	fmt.Println(msg)
	txbytes, err := auth.CompleteAndBroadcastTxCLI(txBldr, s.ctx, []types.Msg{msg}, true)
	fmt.Println(txbytes, err)
}

// query transaction hashHexStr
func (s *Service) QueryTX(hashHexStr string) {
	res, err := auth.QueryTx(s.ctx, hashHexStr)
	if err != nil {
		fmt.Println(err)
	}
	if res.Empty() {
		fmt.Println(fmt.Errorf("No transaction found with hash %s", hashHexStr))
	}
	fmt.Println(res)
}

package revocableTx

import (
	"fmt"
	"testing"

	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

//	gatecli	authcmd.QueryTxCmd(cdc),
// gatecli revocable-tx show [hash] [flags]
/*
 gatecli revocable-tx show
 REVOCABLEPAY-AD60C3155867902FFDB8CEBD7E145ACD82CC6B661C5AE261E01AE6A435C209E0254633DAAA220523BDAD55AC34F7DAA1
*/
// done
func TestQueryTxCmd(t *testing.T) {
	hashHexStr := "IRREVOCABLEPAY-D9451D7488A0D1129654F621D3292D1F8041477D74982D8E9DE0BE67084BC4D3F0D964E91AF319B13F910DCEA26AE167"

	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	QueryTxCmd(ctx, hashHexStr)
}

//	authcmd.GetAccountRevocableTxCmd(cdc),
//
// gatecli revocable-tx list  vault11xcq2vcd2e4vastn9q9vd7kh7l8ksfrpghdkefzvte9a98h9rl8cmasr47gtzdr7ljdg3e6
// done
func TestGetAccountRevocableTxCmd(t *testing.T) {
	vaultAddr := "vault11xcq2vcd2e4vastn9q9vd7kh7l8ksfrpghdkefzvte9a98h9rl8cmasr47gtzdr7ljdg3e6"
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	GetAccountRevocableTxCmd(ctx, vaultAddr)
}

/*
 gatecli revocable-tx send gt11dur74yyhmpytwrx83yerz7hk0ukhrqz2ekvxsje2fl772ceeh8wkkk6dv6gv20elvz0gcv
 100000NANOGT --from vault11f40ewc2ue9466p04gp3xp7h0k49kl7fk0jehdcwajsgz0m092nxqrjyjq6tndv4at5hv5g  --fees
 100000000NANOGT --chain-id gate-66 -y
*/
// revocablecmd.RevocableTxSendCmd
// done
func TestRevocableTxSendCmd(t *testing.T) {

	from_addr := "vault11xcq2vcd2e4vastn9q9vd7kh7l8ksfrpghdkefzvte9a98h9rl8cmasr47gtzdr7ljdg3e6"

	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(from_addr, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	to_addr := "gt11xcq2vcd2e4vastn9q9vd7kh7l8ksfrpghdkefzvte9a98h9rl8cmasr47gtzdr7lzd7aca"
	amount := "100000000000NANOGT"

	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}
	RevocableTxSendCmd(ctx, txBldr, from_addr, to_addr, amount)
}

/*
gatecli revocable-tx revoke REVOCABLEPAY-8E16AF2EEACA40ABB3521E5EA36C777D76181C9DBFB34C2804D08F3E108771F28B3DFAE475BE69F380290EAD4E5AFFD3
--from vault112wldf46teljw9uarndfjqcs9dynrrpkwtmltcwg9yhpun8weh2829783n4ut5hzduw0npt
--fees 100000000NANOGT --chain-id gate-66 -y
*/
// revocablecmd.RevokeTxCmd(cdc),
// done
func TestRevokeTxCmd(t *testing.T) {

	from_addr := "vault11xcq2vcd2e4vastn9q9vd7kh7l8ksfrpghdkefzvte9a98h9rl8cmasr47gtzdr7ljdg3e6"
	hashHexStr := "REVOCABLEPAY-D324BDE93F03D86D9F87CC346825FAC2C7B34FB98AD7BB76E5DCA0D95BFA2FDB881C5D097408F943A7DFE0A7AB5BFAF3"

	cdc := utils.MakeCodec()

	ctx := context.NewCLIContextWithFrom(from_addr, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	txBldr := auth2.NewTxBuilderFromCLI2(common.RootDir, fees, chainID, gas, cdc)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	if err != nil {
		fmt.Println(err)
	}
	RevokeTxCmd(ctx, txBldr, from_addr, hashHexStr)
}

//revocablecmd.QueryTxCmd(cdc)
/* gatecli revocable-tx  status
REVOCABLEPAY-8E16AF2EEACA40ABB3521E5EA36C777D76181C9DBFB34C2804D08F3E108771F28B3DFAE475BE69F380290EAD4E5AFFD3
*/
// done
func TestTxStatusCmd(t *testing.T) {

	hashHexStr := "IRREVOCABLEPAY-D9451D7488A0D1129654F621D3292D1F8041477D74982D8E9DE0BE67084BC4D3F0D964E91AF319B13F910DCEA26AE167"

	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	TxStatusCmd(hashHexStr, ctx)
}

package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestQueryTxCmd(t *testing.T) {
	hashHexStr := "IRREVOCABLEPAY-D9451D7488A0D1129654F621D3292D1F8041477D74982D8E9DE0BE67084BC4D3F0D964E91AF319B13F910DCEA26AE167"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.RevocableTx.QueryTx(hashHexStr)
}

func TestGetAccountRevocableTxCmd(t *testing.T) {
	vaultAddr := "vault114v3kx0330l6pde8qjzgqkm0yqaafz4y3y6vyr90ktqngaccajghklsrefjqutkruqxzkhv"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.RevocableTx.GetAccountRevocableTx(vaultAddr)
}

func TestRevocableTxSendCmd(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	from_addr := "vault11xcq2vcd2e4vastn9q9vd7kh7l8ksfrpghdkefzvte9a98h9rl8cmasr47gtzdr7ljdg3e6"
	to_addr := "gt11xcq2vcd2e4vastn9q9vd7kh7l8ksfrpghdkefzvte9a98h9rl8cmasr47gtzdr7lzd7aca"
	amount := "100000000000NANOGT"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFrom(from_addr, common.EndPoint, common.APIToken, common.RootDir)
	client.RevocableTx.RevocableTxSend(from_addr, to_addr, amount, fees, chainID, gas)

}

func TestRevokeTxCmd(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	from_addr := "vault11xcq2vcd2e4vastn9q9vd7kh7l8ksfrpghdkefzvte9a98h9rl8cmasr47gtzdr7ljdg3e6"
	hashHexStr := "REVOCABLEPAY-D324BDE93F03D86D9F87CC346825FAC2C7B34FB98AD7BB76E5DCA0D95BFA2FDB881C5D097408F943A7DFE0A7AB5BFAF3"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFrom(from_addr, common.EndPoint, common.APIToken, common.RootDir)
	client.RevocableTx.RevokeTx(from_addr, hashHexStr, fees, chainID, gas)
}

func TestTxStatusCmd(t *testing.T) {
	hashHexStr := "IRREVOCABLEPAY-D9451D7488A0D1129654F621D3292D1F8041477D74982D8E9DE0BE67084BC4D3F0D964E91AF319B13F910DCEA26AE167"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.RevocableTx.TxStatus(hashHexStr)
}

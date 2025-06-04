package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestQueryTxCmd(t *testing.T) {
	hashHexStr := "REVOCABLEPAY-2E57289F4C81752E6726B313B57684BB82DD930A15DAECDE8C333E1A21EA04E79DF0A6EAFA2A9D58D6B6FEF3770BB8BF"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.RevocableTx.QueryTx(hashHexStr)
}

func TestGetAccountRevocableTxCmd(t *testing.T) {
	vaultAddr := "vault11k4rhqkllr5fke2hy2pyuj92vhk3jv8u7vzsvg4y8qkxppv8m29xnz2dm9kwp6hhgg5g47t"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.RevocableTx.GetAccountRevocableTx(vaultAddr)
}

func TestTxStatusCmd(t *testing.T) {
	hashHexStr := "REVOCABLEPAY-2E57289F4C81752E6726B313B57684BB82DD930A15DAECDE8C333E1A21EA04E79DF0A6EAFA2A9D58D6B6FEF3770BB8BF"
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.RevocableTx.TxStatus(hashHexStr)
}

package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestQueryVaultAccount(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	vaultAddr := "vault11k4rhqkllr5fke2hy2pyuj92vhk3jv8u7vzsvg4y8qkxppv8m29xnz2dm9kwp6hhgg5g47t"
	client.VaultAccount.QueryVaultAccount(vaultAddr)
}

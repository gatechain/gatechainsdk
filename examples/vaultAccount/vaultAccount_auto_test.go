package main

import (
	"strconv"
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
	"github.com/gatechain/gatechainsdk/examples/testutil"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
)

var kb keys.Keybase

func init() {
	kb = testutil.ImportAllPrivateKeys()
}

func TestBroadcastMsgCreateVaultTx(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	from_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClientWithFromMemKeyBase(from_addr, common.EndPoint, common.APIToken, kb)
	currentHeight, err := client.Ctx.GetChainHeight()
	if err != nil {
		t.Fatal(err)
	}
	clearTimeHeigh := currentHeight + 20000
	clearTimeHeightStr := strconv.FormatUint(uint64(clearTimeHeigh), 10)

	to_addr := "gt11k4rhqkllr5fke2hy2pyuj92vhk3jv8u7vzsvg4y8qkxppv8m29xnz2dm9kwp6hhgc57elv"
	security_addr := "gt116lma26rvst9xhlg6fm2facm02mpgfhtcje3l0zul2ey06kfvfdf03q9u6vyd9fz7n65l7e"
	delayHeightStr := "10"

	coinsStr := "100000000000NANOGT"
	pubkeyStr := ""
	fees := "100000000NANOGT"
	chainID := "gate-66"
	gas := uint64(200000)

	client.VaultAccount.BroadcastMsgCreateVault(from_addr, to_addr, security_addr, delayHeightStr, clearTimeHeightStr,
		coinsStr, pubkeyStr, fees, chainID, gas, kb)

}

func TestBroadcastUpdateClearingHeightTx(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	//to addr
	from_addr := "gt11k4rhqkllr5fke2hy2pyuj92vhk3jv8u7vzsvg4y8qkxppv8m29xnz2dm9kwp6hhgc57elv"
	vaultAddr := "vault11k4rhqkllr5fke2hy2pyuj92vhk3jv8u7vzsvg4y8qkxppv8m29xnz2dm9kwp6hhgg5g47t"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	client := client.NewClientWithFromMemKeyBase(from_addr, common.EndPoint, common.APIToken, kb)
	currentHeight, err := client.Ctx.GetChainHeight()
	if err != nil {
		t.Fatal(err)
	}
	clearTimeHeigh := currentHeight + 30000
	clearTimeHeightStr := strconv.FormatUint(uint64(clearTimeHeigh), 10)
	client.VaultAccount.BroadcastUpdateClearingHeightTx(clearTimeHeightStr, vaultAddr, fees, chainID, gas, kb)
}

func TestClearVaultAccountTx(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	from_addr := "gt116lma26rvst9xhlg6fm2facm02mpgfhtcje3l0zul2ey06kfvfdf03q9u6vyd9fz7n65l7e"
	fees := "100000000NANOGT"
	chainID := "gate-66"
	gas := uint64(200000)
	vaultAddresses := []string{
		"vault11k4rhqkllr5fke2hy2pyuj92vhk3jv8u7vzsvg4y8qkxppv8m29xnz2dm9kwp6hhgg5g47t",
	}
	client := client.NewClientWithFromMemKeyBase(from_addr, common.EndPoint, common.APIToken, kb)

	client.VaultAccount.ClearVaultAccountTx(fees, chainID, vaultAddresses, gas, kb)
}

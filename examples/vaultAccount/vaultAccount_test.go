package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestBroadcastMsgCreateVault(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	from_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClientWithFrom(from_addr, common.EndPoint, common.APIToken, common.RootDir)

	to_addr := "gt116gptv427t2tu447arzqkz2uhmwk8rmzxc09cffcmg9fsh3ntt7cszy4wjevhhvccn37m52"
	security_addr := "gt11469zdcp9ehj4xvv5q2kvs4uhs62uvuul73a364h5dwdtk8mwxfud03vvrzepgm23tccdm6"
	delayHeightStr := "10"
	clearTimeHeightStr := "20000"
	coinsStr := "100000000000NANOGT"
	pubkeyStr := ""
	fees := "100000000NANOGT"
	chainID := "gate-66"
	gas := uint64(200000)

	client.VaultAccount.BroadcastMsgCreateVault(from_addr, to_addr, security_addr, delayHeightStr, clearTimeHeightStr,
		coinsStr, pubkeyStr, fees, chainID, gas)

}

func TestBroadcastUpdateClearingHeightTx(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	from_addr := "gt116gptv427t2tu447arzqkz2uhmwk8rmzxc09cffcmg9fsh3ntt7cszy4wjevhhvccn37m52"

	clearTimeHeight := "20020"
	vaultAddr := "vault116gptv427t2tu447arzqkz2uhmwk8rmzxc09cffcmg9fsh3ntt7cszy4wjevhhvccr3gh4d"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFrom(from_addr, common.EndPoint, common.APIToken, common.RootDir)
	client.VaultAccount.BroadcastUpdateClearingHeightTx(clearTimeHeight, vaultAddr, fees, chainID, gas)
}

func TestQueryVaultAccount(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	vaultAddr := "vault116gptv427t2tu447arzqkz2uhmwk8rmzxc09cffcmg9fsh3ntt7cszy4wjevhhvccr3gh4d"
	client.VaultAccount.QueryVaultAccount(vaultAddr)

}

func TestClearVaultAccountTx(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	from_addr := "gt1196669klaryfct82mtylw7hqzgyx59e08f7f8nen0p56569tkma250s22jhmyfdyhmvda0f"
	fees := "100000000NANOGT"
	chainID := "gate-66"
	gas := uint64(200000)
	vaultAddresses := []string{
		"vault11zt4y23l79rwksl6w6mgrta8sarqwqe9kwyauc6g4gd8at6wuzzhsyqvkrz3fle79f42tmp",
	}
	client := client.NewClientWithFrom(from_addr, common.EndPoint, common.APIToken, common.RootDir)

	client.VaultAccount.ClearVaultAccountTx(fees, chainID, vaultAddresses, gas)
}

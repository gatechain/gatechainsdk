package main

import (
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

func TestRevocableTxSendCmd(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	from_addr := "vault11k4rhqkllr5fke2hy2pyuj92vhk3jv8u7vzsvg4y8qkxppv8m29xnz2dm9kwp6hhgg5g47t"
	to_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	amount := "100000NANOGT"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFromMemKeyBase(from_addr, common.EndPoint, common.APIToken, kb)
	client.RevocableTx.RevocableTxSend(from_addr, to_addr, amount, fees, chainID, gas, kb)

}

func TestRevokeTxCmd(t *testing.T) {
	t.Skip("Skipping only can execute once ")
	from_addr := "vault11k4rhqkllr5fke2hy2pyuj92vhk3jv8u7vzsvg4y8qkxppv8m29xnz2dm9kwp6hhgg5g47t"
	hashHexStr := "REVOCABLEPAY-2E57289F4C81752E6726B313B57684BB82DD930A15DAECDE8C333E1A21EA04E79DF0A6EAFA2A9D58D6B6FEF3770BB8BF"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client := client.NewClientWithFromMemKeyBase(from_addr, common.EndPoint, common.APIToken, kb)
	client.RevocableTx.RevokeTx(from_addr, hashHexStr, fees, chainID, gas, kb)
}

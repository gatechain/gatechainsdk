package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestQueryAccount(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client.Account.QueryAccount(addr)
}

func TestCreateAccount(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	name := ""
	client.Account.CreateAccount(name)
}

func TestGetAccountBlance(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client.Account.GetAccountBlance(addr)
}

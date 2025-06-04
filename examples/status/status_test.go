package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestStatus(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Status.Status()
}

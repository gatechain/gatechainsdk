package block

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

func Block(node *context.NodeVaultQuerierImpl, blockHeight uint64) {
	block, err := node.Client.Block(blockHeight)
	if err != nil {
		fmt.Println(err)
	}
	cdc := utils.MakeCodec()

	utils.JsonOutPut(block, cdc)
}

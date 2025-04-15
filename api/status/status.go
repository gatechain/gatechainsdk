package status

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

func Status(node *context.NodeVaultQuerierImpl) {
	status, err := node.Client.Status()
	if err != nil {
		fmt.Println(err)
	}
	utils.JsonOutPut(status, node.GetCodec())
}

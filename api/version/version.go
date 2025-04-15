package version

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

// done
func Version(node *context.NodeVaultQuerierImpl) {
	version, err := node.Client.Versions()
	if err != nil {
		fmt.Println(err)
	}
	utils.JsonOutPut(version, node.GetCodec())
}

package block

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

func TestBlock(t *testing.T) {
	blockHeight := uint64(1)
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)

	Block(ctx, blockHeight)
}

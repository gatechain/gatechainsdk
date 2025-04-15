package version

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

// done
func TestVersion(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx.WithCodec(cdc)
	Version(ctx)
}

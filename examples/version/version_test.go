package version

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/common"
)

func TestVersion(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Version.Version()
}

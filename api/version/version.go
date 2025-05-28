package version

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

type Service struct {
	ctx *context.NodeVaultQuerierImpl
}

func NewService(ctx *context.NodeVaultQuerierImpl) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) Version() error {
	version, err := s.ctx.Client.Versions()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return utils.JsonOutPut(version, s.ctx.GetCodec())
}

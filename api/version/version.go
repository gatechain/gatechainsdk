package version

import (
	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/rpc/spec/common"
)

type Service struct {
	ctx *context.NodeVaultQuerierImpl
}

func NewService(ctx *context.NodeVaultQuerierImpl) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) Version() (response common.Version, err error) {
	version, err := s.ctx.Client.Versions()
	if err != nil {
		return response, err
	}
	utils.JsonOutPut(version, s.ctx.GetCodec())
	return version, nil
}

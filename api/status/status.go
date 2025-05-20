package status

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

func (s *Service) Status() {
	status, err := s.ctx.Client.Status()
	if err != nil {
		fmt.Println(err)
		return
	}
	utils.JsonOutPut(status, s.ctx.GetCodec())
}

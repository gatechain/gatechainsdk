package status

import (
	"fmt"
	v1 "github.com/gatechain/gatechainsdk/gatechain/rpc/spec/v1"

	"github.com/gatechain/gatechainsdk/api/utils"
	"github.com/gatechain/gatechainsdk/gatechain/context"
)

type Service struct {
	ctx *context.NodeVaultQuerierImpl
}

func NewService(ctx *context.NodeVaultQuerierImpl) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) Status() (response v1.NodeStatus, err error) {
	status, err := s.ctx.Client.Status()
	if err != nil {
		fmt.Println(err)
		return v1.NodeStatus{}, err
	}
	utils.JsonOutPut(status, s.ctx.GetCodec())
	return status, nil
}

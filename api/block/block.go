package block

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

func (s *Service) Block(blockHeight uint64) {
	height := int64(blockHeight)
	if blockHeight == 0 {
		height, _ = s.ctx.GetChainHeight()
	}

	block, err := s.ctx.Client.Block(uint64(height))
	if err != nil {
		fmt.Println(err)
		return
	}
	cdc := utils.MakeCodec()

	utils.JsonOutPut(block, cdc)
}

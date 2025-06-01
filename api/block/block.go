package block

import (
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

func (s *Service) Block(blockHeight uint64) (response v1.Block, err error) {
	height := int64(blockHeight)
	if blockHeight == 0 {
		height, _ = s.ctx.GetChainHeight()
	}

	block, err := s.ctx.Client.Block(uint64(height))
	if err != nil {
		return v1.Block{}, err
	}
	cdc := utils.MakeCodec()
	utils.JsonOutPut(block, cdc)
	return block, nil
}

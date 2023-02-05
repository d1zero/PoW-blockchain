package domain

import (
	"pow-blockchain/internal/domain/dto"
	"pow-blockchain/internal/domain/entity"
)

type BlockService interface {
	CalculateHash(entity.Block) string
	GetBlockchain() ([]entity.Block, error)
	WriteBlock(dto.WriteBlock) error
}

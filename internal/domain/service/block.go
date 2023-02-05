package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"pow-blockchain/internal/domain"
	"pow-blockchain/internal/domain/dto"
	"pow-blockchain/internal/domain/entity"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Block struct {
	Blockchain []entity.Block
	mutex      *sync.Mutex
}

func (s *Block) GetBlockchain() ([]entity.Block, error) {
	return s.Blockchain, nil
}

func (s *Block) isBlockValid(newBlock, oldBlock entity.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if s.CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func (s *Block) generateBlock(oldBlock entity.Block, data int64) entity.Block {
	newBlock := entity.Block{
		Index:      oldBlock.Index + 1,
		TimeStamp:  time.Now().String(),
		Data:       data,
		PrevHash:   oldBlock.Hash,
		Difficulty: entity.Difficulty,
	}

	for i := 0; ; i++ {
		if s.isHashValid(s.CalculateHash(newBlock), entity.Difficulty) {
			newBlock.Nonce = strconv.Itoa(i)
			newBlock.Hash = s.CalculateHash(newBlock)
			break
		}
	}

	return newBlock
}

func (s *Block) CalculateHash(block entity.Block) string {
	record := strconv.FormatInt(block.Index, 10) + block.TimeStamp +
		strconv.FormatInt(block.Data, 10) + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func (s *Block) isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

func (s *Block) WriteBlock(p dto.WriteBlock) error {
	s.mutex.Lock()
	newBlock := s.generateBlock(s.Blockchain[len(s.Blockchain)-1], p.Data)
	s.mutex.Unlock()

	if !s.isBlockValid(newBlock, s.Blockchain[len(s.Blockchain)-1]) {
		return errors.New("invalid block")
	}

	s.Blockchain = append(s.Blockchain, newBlock)

	return nil
}

var _ domain.BlockService = &Block{}

func NewBlockService() *Block {
	blockService := &Block{
		mutex: &sync.Mutex{},
	}
	blockService.Blockchain = append(blockService.Blockchain, entity.Block{
		Index:      0,
		TimeStamp:  time.Now().String(),
		Data:       0,
		Hash:       blockService.CalculateHash(entity.Block{}),
		PrevHash:   "",
		Difficulty: entity.Difficulty,
		Nonce:      "",
	})

	return blockService
}

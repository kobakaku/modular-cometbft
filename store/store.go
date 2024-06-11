package store

import (
	"fmt"
	"sync/atomic"

	"github.com/kobakaku/modular-cometbft/types"
)

type DefaultStore struct {
	blocks []*types.Block
	height atomic.Uint64
}

var _ Store = &DefaultStore{}

func New() *DefaultStore {
	return &DefaultStore{}
}

func (s *DefaultStore) SetHeight(height uint64) {
	s.height.Store(height)
}

func (s *DefaultStore) Height() uint64 {
	return s.height.Load()
}

// SaveBlock adds block to the store.
func (s *DefaultStore) SaveBlock(block *types.Block) error {
	s.blocks = append(s.blocks, block)
	return nil
}

// GetBlock returns block at given height.
// TODO: indexing height→hash, and store blocks by hash.
func (s *DefaultStore) GetBlock(height uint64) (*types.Block, error) {
	index := height - 1
	if uint64(len(s.blocks)) == index {
		return nil, fmt.Errorf("error getting block")
	} else {
		block := s.blocks[index]
		return block, nil
	}
}

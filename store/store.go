package store

import (
	"sync/atomic"

	"github.com/kobakaku/modular-cometbft/types"
)

type DefaultStore struct {
	blocks []types.Block
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

func (s *DefaultStore) GetBlock(height uint64) (*types.Block, error) {
	return &types.Block{
		Txs: []types.Tx{},
	}, nil
}

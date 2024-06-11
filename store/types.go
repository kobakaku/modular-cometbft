package store

import "github.com/kobakaku/modular-cometbft/types"

type Store interface {
	SetHeight(height uint64)

	Height() uint64

	SaveBlock(block *types.Block) error

	GetBlock(height uint64) (*types.Block, error)
}

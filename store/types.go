package store

import (
	"context"

	"github.com/kobakaku/modular-cometbft/types"
)

type Store interface {
	SetHeight(height uint64)

	Height() uint64

	SaveBlock(ctx context.Context, block *types.Block) error

	GetBlock(ctx context.Context, height uint64) (*types.Block, error)
}

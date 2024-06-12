package store

import (
	"context"
	"fmt"
	"strconv"
	"sync/atomic"

	ds "github.com/ipfs/go-datastore"

	"github.com/kobakaku/modular-cometbft/types"
)

type DefaultStore struct {
	db     ds.TxnDatastore
	height atomic.Uint64
}

var _ Store = &DefaultStore{}

func New(ds ds.TxnDatastore) *DefaultStore {
	return &DefaultStore{
		db: ds,
	}
}

func (s *DefaultStore) SetHeight(height uint64) {
	s.height.Store(height)
}

func (s *DefaultStore) Height() uint64 {
	return s.height.Load()
}

// SaveBlock adds block to the store.
func (s *DefaultStore) SaveBlock(ctx context.Context, block *types.Block) error {
	blockBlob, err := block.MarshalBinary(block)
	if err != nil {
		return fmt.Errorf("failed to marshal Block to binary form: %w", err)
	}

	txn, err := s.db.NewTransaction(ctx, false)
	if err != nil {
		return fmt.Errorf("failed to create a new batch for transaction: %w", err)
	}

	txn.Put(ctx, ds.NewKey(getBlockKey(block.Height())), blockBlob)

	return nil
}

// GetBlock returns block at given height.
// TODO: indexing height→hash, and store blocks by hash.
func (s *DefaultStore) GetBlock(ctx context.Context, height uint64) (*types.Block, error) {
	blockBlob, err := s.db.Get(ctx, ds.NewKey(getBlockKey(height)))
	if err != nil {
		return nil, fmt.Errorf("failed to get a block at ginen height: %w", err)
	}

	block := new(types.Block)
	block, err = block.UnmarshalBinary(blockBlob)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal binary format of Block into Block: %w", err)
	}

	return block, nil
}

func getBlockKey(height uint64) string {
	return GenerateKey([]string{"b", strconv.FormatUint(height, 10)})
}

// func getBlockKey(data []byte) string {
// 	return GenerateKey([]string{"b", hex.EncodeToString(data)})
// }

// func getIndexKey(data []byte) string {
// 	return GenerateKey([]string{"i", hex.EncodeToString(data)})
// }

package state

import (
	"github.com/cometbft/cometbft/libs/log"

	"github.com/kobakaku/modular-cometbft/types"
)

// BlockExecutor creates and applies blocks and maintains state
type BlockExecutor struct {
	logger log.Logger
}

// NewBlockExecutor creates new instance of BlockExecutor.
func NewBlockExecutor(logger log.Logger) *BlockExecutor {
	return &BlockExecutor{logger: logger}
}

// CreateBlock gets transactions from mempool and builds a block.
func (be *BlockExecutor) CreateBlock() error {
	return nil
}

// ApplyBlock executes the block
func (be *BlockExecutor) ApplyBlock(block *types.Block) error {
	// This makes calls to the AppClient
	err := be.execute(block)
	if err != nil {
		return err
	}
	return nil
}

func (be *BlockExecutor) execute(block *types.Block) error {
	// TODO: TXを実行する
	be.logger.Debug("executed block", "num_txs", len(block.Txs))
	return nil
}

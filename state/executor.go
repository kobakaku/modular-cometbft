package state

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/proxy"

	"github.com/kobakaku/modular-cometbft/types"
)

// BlockExecutor creates and applies blocks and maintains state
type BlockExecutor struct {
	proxyApp proxy.AppConnConsensus

	logger log.Logger
}

// NewBlockExecutor creates new instance of BlockExecutor.
func NewBlockExecutor(proxyApp proxy.AppConnConsensus, logger log.Logger) *BlockExecutor {
	return &BlockExecutor{proxyApp: proxyApp, logger: logger}
}

// CreateBlock gets transactions from mempool and builds a block.
func (be *BlockExecutor) CreateBlock() error {
	return nil
}

// ApplyBlock executes the block
func (be *BlockExecutor) ApplyBlock(ctx context.Context, block *types.Block) error {
	// This makes calls to the AppClient
	_, err := be.execute(ctx, block)
	if err != nil {
		return err
	}
	return nil
}

func (be *BlockExecutor) Commit(ctx context.Context) (int64, error) {
	respCommit, err := be.proxyApp.Commit(ctx)
	if err != nil {
		return 0, err
	}
	return respCommit.RetainHeight, nil
}

func (be *BlockExecutor) execute(ctx context.Context, block *types.Block) (*abci.ResponseFinalizeBlock, error) {
	resp, err := be.proxyApp.FinalizeBlock(ctx, &abci.RequestFinalizeBlock{})
	if err != nil {
		return nil, err
	}
	be.logger.Debug("executed block", "num_txs", len(block.Txs))
	return resp, nil
}

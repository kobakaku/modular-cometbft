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
func (be *BlockExecutor) CreateBlock(height uint64) (*types.Block, error) {
	// TODO: mempoolからtransactionを取得し、blockを作成する
	return &types.Block{
		Header: types.Header{BaseHeader: types.BaseHeader{Height: height}},
		Txs:    []types.Tx{[]byte("TODO")},
	}, nil
}

// ApplyBlock executes the block
func (be *BlockExecutor) ApplyBlock(ctx context.Context, block *types.Block) (*abci.ResponseFinalizeBlock, error) {
	// This makes calls to the AppClient
	resp, err := be.execute(ctx, block)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Commit commits the block
func (be *BlockExecutor) Commit(ctx context.Context, block *types.Block, resp *abci.ResponseFinalizeBlock) ([]byte, int64, error) {
	commitResp, err := be.proxyApp.Commit(ctx)
	if err != nil {
		return nil, 0, err
	}
	return resp.AppHash, commitResp.RetainHeight, nil
}

func (be *BlockExecutor) execute(ctx context.Context, block *types.Block) (*abci.ResponseFinalizeBlock, error) {
	resp, err := be.proxyApp.FinalizeBlock(ctx, &abci.RequestFinalizeBlock{})
	if err != nil {
		return nil, err
	}
	be.logger.Debug("executed block", "num_txs", len(block.Txs))
	return resp, nil
}

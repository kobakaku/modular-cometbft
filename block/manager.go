package block

import (
	"context"
	"fmt"
	"time"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/kobakaku/modular-cometbft/da"
	"github.com/kobakaku/modular-cometbft/state"
	"github.com/kobakaku/modular-cometbft/store"
	"github.com/kobakaku/modular-cometbft/types"
)

type Manager struct {
	daClient *da.DAClient
	store    store.Store
	executor *state.BlockExecutor

	logger log.Logger
}

func NewManager(daClient *da.DAClient, store store.Store, logger log.Logger) (*Manager, error) {
	exec := state.NewBlockExecutor(logger)

	mgr := &Manager{
		daClient: daClient,
		store:    store,
		executor: exec,
		logger:   logger,
	}
	return mgr, nil
}

// BlockSubmissionLoop is responsible for submitting blocks to the DA layer.
func (m *Manager) BlockSubmissionLoop(ctx context.Context) {
	// TODO: 適切なブロック送信間隔を考える (現状 30s)
	timer := time.NewTicker(30_000000000)
	defer timer.Stop()

	for {
		select {
		// TODO: ここの分岐は必要か確認
		case <-ctx.Done():
			return
		case <-timer.C:
		}

		err := m.submitBlocksToDA(ctx)
		if err != nil {
			m.logger.Error("error while submitting block to DA", "error", err)
		}
	}

}

// AggregationLoop is responsible for aggregating transactions into rollup blocks.
func (m *Manager) AggregationLoop(ctx context.Context) {
	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		select {
		// TODO: ここの分岐は必要か確認
		case <-ctx.Done():
			return
		case <-timer.C:
		}
		m.publishBlock()

		// TODO: 適切なブロック生成間隔を考える (現状 10s)
		timer.Reset(10_000000000)
	}
}

func (m *Manager) publishBlock() error {
	height := m.store.Height()
	newHeight := height + 1

	var block *types.Block

	block, err := m.store.GetBlock(newHeight)
	if block != nil {
		m.logger.Error("error getting wrong block", "height", newHeight)
	} else {
		m.logger.Debug("creating block info", "num_txs", len(block.Txs))
		block, err = m.createBlock(newHeight)
		if err != nil {
			return err
		}

		m.store.SaveBlock(block)
		m.store.SetHeight(newHeight)
	}

	m.applyBlock(block)

	m.logger.Debug("Creating and publishing block", "height", newHeight)

	// TODO: DBにBlockを保存する

	m.logger.Info("successfully proposed block", "height", newHeight)

	return nil
}

func (m *Manager) submitBlocksToDA(ctx context.Context) error {
	// TODO: 正しいblobを取得する
	var (
		blobs [][]byte
	)

	ctx, cancel := context.WithTimeout(ctx, 10000000)
	defer cancel()

	ids, err := m.daClient.DA.Submit(ctx, blobs, m.daClient.GasPrice, m.daClient.Namespace)
	if err != nil {
		return fmt.Errorf("error while submitting blocks to DA layer: %w", err)
	}
	if len(ids) == 0 {
		m.logger.Error("failed to submit blocks: unexpected len(ids): 0")
	}

	m.logger.Info("successfully submitted Rollkit blocks to DA layer", "ids", ids)
	return nil
}

func (m *Manager) createBlock(height uint64) (*types.Block, error) {
	// TODO: mempoolからblockを生成する
	return &types.Block{
		Txs: []types.Tx{},
	}, nil
}

func (m *Manager) applyBlock(block *types.Block) (*types.Block, error) {
	err := m.executor.ApplyBlock(block)
	if err != nil {
		return nil, fmt.Errorf("error while applying blocks: %w", err)
	}
	return block, nil
}

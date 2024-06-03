package block

import (
	"context"
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/kobakaku/modular-cometbft/da"
)

type Manager struct {
	daClient *da.DAClient

	logger log.Logger
}

func NewManager(daClient *da.DAClient, logger log.Logger) (*Manager, error) {
	mgr := &Manager{
		daClient: daClient,
		logger:   logger,
	}
	return mgr, nil
}

// BlockSubmissionLoop is responsible for submitting blocks to the DA layer.
func (m *Manager) BlockSubmissionLoop(ctx context.Context) {
	err := m.submitBlocksToDA(ctx)
	if err != nil {
		m.logger.Error("error while submitting block to DA", "error", err)
	}
}

// AggregationLoop is responsible for aggregating transactions into rollup blocks.
func (m *Manager) AggregationLoop() {

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

	m.logger.Info("successfully submitted Rollkit blocks to DA layer", "ids", ids)
	return nil
}

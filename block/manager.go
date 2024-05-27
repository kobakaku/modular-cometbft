package block

import "github.com/kobakaku/modular-cometbft/da"

type Manager struct {
	daClient da.DAClient
}

func (m *Manager) BlockSubmissionLoop() {}
func (m *Manager) AggregationLoop()     {}

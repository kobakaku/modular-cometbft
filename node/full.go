package node

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/kobakaku/modular-cometbft/block"
	"github.com/kobakaku/modular-cometbft/da"
	"github.com/kobakaku/modular-cometbft/utils"
)

var _ Node = &FullNode{}

type FullNode struct {
	*service.BaseService

	client       rpcclient.Client
	daClient     *da.DAClient
	blockManager *block.Manager

	threadManager utils.ThreadManager
}

func newFullNode(logger log.Logger) (fn *FullNode, err error) {
	node := &FullNode{}
	node.BaseService = service.NewBaseService(logger, "FullNode", node)
	return node, nil
}

func (fn *FullNode) OnStart() error {
	fn.Logger.Info("starting full node...")

	fn.threadManager.Go(func() { fn.blockManager.BlockSubmissionLoop() })
	fn.threadManager.Go(func() { fn.blockManager.AggregationLoop() })

	return nil
}

func (fn *FullNode) GetClient() rpcclient.Client {
	return fn.client
}

func (fn *FullNode) OnStop() {
	fn.Logger.Info("halting full node...")
	fn.Logger.Error("errors while stopping node:", "errors", "context canceled")
}

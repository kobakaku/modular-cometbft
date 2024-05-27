package node

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
)

var _ Node = &FullNode{}

type FullNode struct {
	*service.BaseService

	client rpcclient.Client
}

func newFullNode(logger log.Logger) (fn *FullNode, err error) {
	node := &FullNode{}
	node.BaseService = service.NewBaseService(logger, "FullNode", node)
	return node, nil
}

func (fn *FullNode) OnStart() error {
	fn.Logger.Info("starting full node...")
	return nil
}

func (fn *FullNode) GetClient() rpcclient.Client {
	return fn.client
}

func (fn *FullNode) OnStop() {
	fn.Logger.Info("halting full node...")
	fn.Logger.Error("errors while stopping node:", "errors", "context canceled")
}

package node

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
)

var _ Node = &LightNode{}

type LightNode struct {
	*service.BaseService

	client rpcclient.Client
}

func newLightNode(logger log.Logger) (ln *LightNode, err error) {
	node := &LightNode{}
	node.BaseService = service.NewBaseService(logger, "LightNode", node)
	return node, nil
}

func (ln *LightNode) OnStart() error {
	ln.Logger.Info("starting light node...")
	return nil
}

func (ln *LightNode) GetClient() rpcclient.Client {
	return ln.client
}

func (ln *LightNode) OnStop() {
	ln.Logger.Info("halting light node...")
	ln.Logger.Error("errors while stopping node:", "errors", "context canceled")
}

package node

import (
	"context"
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/service"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/kobakaku/modular-cometbft/block"
	"github.com/kobakaku/modular-cometbft/config"
	"github.com/kobakaku/modular-cometbft/da"
	"github.com/kobakaku/modular-cometbft/utils"
	proxyda "github.com/rollkit/go-da/proxy"
)

var _ Node = &FullNode{}

type FullNode struct {
	*service.BaseService

	client       rpcclient.Client
	daClient     *da.DAClient
	blockManager *block.Manager

	ctx           context.Context
	threadManager utils.ThreadManager
}

func newFullNode(nodeConfig config.NodeConfig, logger log.Logger) (fn *FullNode, err error) {
	daClient, err := initDAClient(nodeConfig)
	if err != nil {
		return nil, err
	}

	node := &FullNode{daClient: daClient}

	node.BaseService = service.NewBaseService(logger, "FullNode", node)

	return node, nil
}

func initDAClient(nodeConfig config.NodeConfig) (*da.DAClient, error) {
	namespace := []byte(nodeConfig.DANamespace)

	client, err := proxyda.NewClient(nodeConfig.DAAddress, nodeConfig.DAAuthToken)
	if err != nil {
		return nil, fmt.Errorf("error while establishing connection to DA layer: %w", err)
	}
	return da.NewDAClient(client, nodeConfig.DAGasPrice, namespace), nil
}

func (fn *FullNode) OnStart() error {
	fn.Logger.Info("starting full node...")

	fn.threadManager.Go(func() { fn.blockManager.BlockSubmissionLoop(fn.ctx) })
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

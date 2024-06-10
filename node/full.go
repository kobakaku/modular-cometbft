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
	"github.com/kobakaku/modular-cometbft/store"
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

func newFullNode(ctx context.Context, nodeConfig config.NodeConfig, logger log.Logger) (fn *FullNode, err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		// If there is an error, cancel the context
		if err != nil {
			cancel()
		}
	}()

	daClient, err := initDAClient(nodeConfig)
	if err != nil {
		return nil, err
	}

	store := store.New()

	blockManager, err := initBlockManager(daClient, store, logger)
	if err != nil {
		return nil, err
	}

	node := &FullNode{daClient: daClient, blockManager: blockManager, ctx: ctx}

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

func initBlockManager(daClient *da.DAClient, store store.Store, logger log.Logger) (*block.Manager, error) {
	blockManager, err := block.NewManager(daClient, store, logger)
	if err != nil {
		return nil, fmt.Errorf("error while initializeing BlockManger: %w", err)
	}
	return blockManager, nil
}

func (fn *FullNode) OnStart() error {
	fn.Logger.Info("starting full node...")

	fn.threadManager.Go(func() { fn.blockManager.AggregationLoop(fn.ctx) })
	fn.threadManager.Go(func() { fn.blockManager.BlockSubmissionLoop(fn.ctx) })

	return nil
}

func (fn *FullNode) GetClient() rpcclient.Client {
	return fn.client
}

func (fn *FullNode) OnStop() {
	fn.Logger.Info("halting full node...")
	fn.Logger.Error("errors while stopping node:", "errors", "context canceled")
}

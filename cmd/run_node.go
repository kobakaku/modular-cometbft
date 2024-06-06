package cmd

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"

	"github.com/cometbft/cometbft/libs/log"
	cometos "github.com/cometbft/cometbft/libs/os"
	"github.com/kobakaku/modular-cometbft/config"
	"github.com/kobakaku/modular-cometbft/node"
	"github.com/kobakaku/modular-cometbft/rpc"

	daproxy "github.com/rollkit/go-da/proxy/jsonrpc"
	goDATest "github.com/rollkit/go-da/test"
)

var (
	nodeConfig = config.DefaultNodeConfig
)

var RunNodeCmd = &cobra.Command{
	Use:   "start",
	Short: "Run node",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Initialize logging
		logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

		// Start Mock DA server
		srv, err := startMockDAServer(context.Background())
		if err != nil {
			return fmt.Errorf("failed to launch mock da server: %w", err)
		}
		defer func() { srv.Stop(cmd.Context()) }()

		// Start p2p node
		node, err := node.NewNode(context.Background(), nodeConfig, logger)
		if err != nil {
			return fmt.Errorf("failed to create new node: %v", err)
		}
		if err := node.Start(); err != nil {
			return fmt.Errorf("failed to start node: %v", err)
		}

		// Start RPC server
		server := rpc.NewServer(logger)
		err = server.Start()
		if err != nil {
			return fmt.Errorf("failed to launch rpc server: %v", err)
		}

		// Stop upon receiving SIGTERM or CTRL-C.
		cometos.TrapSignal(logger, func() {
			if err := node.Stop(); err != nil {
				logger.Error("unable to stop the node", "error", err)
			}
		})

		// Block forever to force user to stop node
		select {}
	},
}

// Start a mock DA server
func startMockDAServer(ctx context.Context) (*daproxy.Server, error) {
	addr, err := url.Parse(nodeConfig.DAAddress)
	if err != nil {
		return nil, err
	}

	srv := daproxy.NewServer(addr.Hostname(), addr.Port(), goDATest.NewDummyDA())
	err = srv.Start(ctx)
	if err != nil {
		return nil, err
	}
	return srv, err
}

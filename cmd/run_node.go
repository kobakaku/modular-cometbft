package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cometbft/cometbft/libs/log"
	cometos "github.com/cometbft/cometbft/libs/os"
	"github.com/kobakaku/modular-cometbft/config"
	"github.com/kobakaku/modular-cometbft/node"
	"github.com/kobakaku/modular-cometbft/rpc"
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

		node, err := node.NewNode(nodeConfig, logger)
		if err != nil {
			return fmt.Errorf("failed to create new node: %v", err)
		}

		server := rpc.NewServer(logger)
		err = server.Start()
		if err != nil {
			return fmt.Errorf("failed to launch rpc server: %v", err)
		}

		if err := node.Start(); err != nil {
			return fmt.Errorf("failed to start node: %v", err)
		}

		// Stop upon receiving SIGTERM or CTRL-C.
		cometos.TrapSignal(logger, func() {
			if node.IsRunning() {
				if err := node.Stop(); err != nil {
					logger.Error("unable to stop the node", "error", err)
				}
			}
		})

		if true {
			// Block forever to force user to stop node
			select {}
		}

		return node.Stop()
	},
}

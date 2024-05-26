package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/kobakaku/modular-cometbft/node"
	"github.com/kobakaku/modular-cometbft/rpc"
)

var RunNodeCmd = &cobra.Command{
	Use:   "start",
	Short: "Run node",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Initialize logging
		logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

		node, err := node.NewNode(logger)
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

		logger.Info("Started node")

		if true {
			// Block forever to force user to stop node
			select {}
		}

		return fmt.Errorf("TODO: should stop node")
	},
}

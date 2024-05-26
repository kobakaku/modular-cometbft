package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/kobakaku/modular-cometbft/rpc"
)

var RunNodeCmd = &cobra.Command{
	Use:   "start",
	Short: "Run node",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running node...")

		// Initialize logging
		logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

		// TODO: create the node

		// TODO: Launch the RPC server
		server := rpc.NewServer(logger)
		err := server.Start()
		if err != nil {
			fmt.Printf("Failed to launch rpc server: %v", err)
		}

		// TODO: Start the node
	},
}

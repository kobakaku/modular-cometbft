package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RunNodeCmd = &cobra.Command{
	Use:   "start",
	Short: "Run node",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runnig node...")
	},
}

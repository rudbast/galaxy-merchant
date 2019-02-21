package cmd

import (
	"fmt"
	"os"

	"github.com/rudbast/galaxy-merchant/core"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "galaxy",
	Short: "Galaxy Merchant Trading is a tool to manage your day-to-day trading conversion.",
	Run: func(cmd *cobra.Command, args []string) {
		core.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

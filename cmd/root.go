package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	nfTraceTable = "nftrace-table"
	nfTraceChain = "nftrace-chain"
)

var rootCmd = &cobra.Command{
	Use:   "nftrace",
	Short: "nftrace",
	Long: `A thin wrapper for nftables rules debugging
                Complete documentation is available at https://wiki.nftables.org/wiki-nftables/index.php/Ruleset_debug/tracing`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

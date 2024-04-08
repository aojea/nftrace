package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"sigs.k8s.io/knftables"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove the trace",
	RunE: func(cmd *cobra.Command, args []string) error {

		nft, err := knftables.New(knftables.InetFamily, nfTraceTable)
		if err != nil {
			return fmt.Errorf("no nftables support: %v", err)
		}

		tx := nft.NewTransaction()
		tx.Delete(&knftables.Table{})
		return nft.Run(context.TODO(), tx)

	},
}

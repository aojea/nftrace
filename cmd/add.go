package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/knftables"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [expression]",
	Short: "add a trace (empty matches all traffic)",
	RunE: func(cmd *cobra.Command, args []string) error {
		trace := []string{"meta", "nftrace", "set", "1"}

		nft, err := knftables.New(knftables.InetFamily, nfTraceTable)
		if err != nil && !knftables.IsNotFound(err) {
			return fmt.Errorf("no nftables support: %v", err)
		}

		tx := nft.NewTransaction()
		tx.Add(&knftables.Table{
			Comment: ptr.To("table for nftrace"),
		})

		tx.Flush(&knftables.Table{})

		tx.Add(&knftables.Chain{
			Name:     nfTraceChain,
			Comment:  knftables.PtrTo("nftrace chain"),
			Type:     ptr.To(knftables.FilterType),
			Hook:     ptr.To(knftables.PreroutingHook),
			Priority: ptr.To(knftables.FilterPriority + "-10"),
		})
		tx.Flush(&knftables.Chain{
			Name: nfTraceChain,
		})

		if len(args) > 0 {
			trace = append(args, trace...)
		}
		tx.Add(&knftables.Rule{
			Chain: nfTraceChain,
			Rule:  knftables.Concat(trace),
		})

		return nft.Run(context.Background(), tx)

	},
}

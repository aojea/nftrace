package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/knftables"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the existing trace",
	RunE: func(cobraCmd *cobra.Command, args []string) error {

		output, err := exec.Command("nft", "list", "table", string(knftables.InetFamily), nfTraceTable).CombinedOutput()
		if err != nil {
			if strings.Contains(string(output), "No such file or directory") {
				fmt.Println("no traces active")
				return nil
			}
			fmt.Println(string(output))
			return err
		}
		if len(output) == 0 {
			fmt.Println("no traces active")
		} else {
			fmt.Println(string(output))
		}
		return nil
	},
}

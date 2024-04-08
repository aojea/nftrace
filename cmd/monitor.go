package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"

	"github.com/spf13/cobra"
	"golang.org/x/sys/unix"
)

func init() {
	rootCmd.AddCommand(monitorCmd)
}

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor the nftables traces",
	RunE: func(c *cobra.Command, args []string) error {
		// trap Ctrl+C and call cancel on the context
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		// Enable signal handler
		signalCh := make(chan os.Signal, 2)
		defer func() {
			close(signalCh)
			cancel()
		}()
		signal.Notify(signalCh, os.Interrupt, unix.SIGINT)

		cmd := exec.CommandContext(ctx, "nft", "monitor", "trace")

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}

		err = cmd.Start()
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
		}
		return cmd.Wait()
	},
}

package app

import (
	"context"

	"github.com/spf13/cobra"
)

var reversalBalanceWorkerCommand = &cobra.Command{
	Use:   "reversal-balance-worker",
	Short: "Worker for reversal balance on inital transfer",
	Run:   runProcessReversalBalanceWorker,
}

func init() {
	rootCmd.AddCommand(reversalBalanceWorkerCommand)
}

func runProcessReversalBalanceWorker(cmd *cobra.Command, args []string) {
	Subscribe.Subscribe(
		awsConfig.ReversalBalanceQueue,
		context.Background(),
		ReversalBalanceWorker.GetHandler(),
	)
}

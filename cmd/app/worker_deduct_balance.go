package app

import (
	"context"

	"github.com/spf13/cobra"
)

var deductBalanceWorkerCommand = &cobra.Command{
	Use:   "deduct-balance-worker",
	Short: "Worker for deduct balance on inital transfer",
	Run:   runProcessDeductBalanceWorker,
}

func init() {
	rootCmd.AddCommand(deductBalanceWorkerCommand)
}

func runProcessDeductBalanceWorker(cmd *cobra.Command, args []string) {
	Subscribe.Subscribe(
		awsConfig.DeductBalanceQueue,
		context.Background(),
		DeductBalanceWorker.GetHandler(),
	)
}

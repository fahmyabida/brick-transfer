package app

import (
	"context"

	"github.com/spf13/cobra"
)

var proceedTransferWorkerCommand = &cobra.Command{
	Use:   "proceed-transfer-worker",
	Short: "Worker for proceed transfer",
	Run:   runProcessProceedTransferWorker,
}

func init() {
	rootCmd.AddCommand(proceedTransferWorkerCommand)
}

func runProcessProceedTransferWorker(cmd *cobra.Command, args []string) {
	Subscribe.Subscribe(
		awsConfig.ProceedTransferQueue,
		context.Background(),
		ProceedTransferWorker.GetHandler(),
	)
}

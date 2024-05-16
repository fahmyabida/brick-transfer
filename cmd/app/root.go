package app

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fahmyabida/brick-transfer/cmd/config"
	"github.com/fahmyabida/brick-transfer/internal/app/domain"
	"github.com/fahmyabida/brick-transfer/internal/app/repository"
	"github.com/fahmyabida/brick-transfer/internal/app/usecase"
	"github.com/fahmyabida/brick-transfer/pkg/external/client/bankserviceclient"
	"github.com/fahmyabida/brick-transfer/pkg/queue/publisher"
	"github.com/fahmyabida/brick-transfer/pkg/queue/worker"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	rootCmd = &cobra.Command{
		Use:   "brick-transfer-service",
		Short: "brick-transfer-service is application for transfer",
	}
)

var (
	// database
	database *gorm.DB

	// aws
	SNS              *sns.SNS
	SQS              *sqs.SQS
	awsConfig        *config.AWS
	bankClientConfig *config.BankClient

	// http client
	httpClient *http.Client

	// queue
	Publisher domain.IPublisher
	Subscribe domain.ISubscribe

	// repository
	TransferRepo    domain.ITransferRepo
	UserBalanceRepo domain.IUserBalanceRepo

	// usecase
	BankAccountUsecase domain.IBankAccountUsecase
	CallbackUsecase    domain.ICallbackUsecase
	TransferUsecase    domain.ITransferUsecase
	UserBalanceUsecase domain.IUserBalanceUsecase

	// worker
	DeductBalanceWorker   domain.IWorker
	ProceedTransferWorker domain.IWorker
	ReversalBalanceWorker domain.IWorker

	// clinet
	BankServiceClient bankserviceclient.IBankServiceClient
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	if err := config.InitEnv(); err != nil {
		log.Fatal(err)
	}

	cobra.OnInitialize(func() {
		initDatabase()
		initAWS()
		initHttpClient()
		initApp()
	})
}

func initDatabase() {
	rw, ro := config.LoadForPostgres()
	database = config.InitDB(rw, ro)
}

func initAWS() {
	awsConfig = config.LoadForAWS()
	SNS, SQS = config.InitAWS_SNS_SQS(awsConfig)
}

func initHttpClient() {
	bankClientConfig = config.LoadForBankClient()
	httpClient = &http.Client{}
}

func initApp() {
	Publisher = publisher.NewPublisher(SNS, *awsConfig)
	Subscribe = worker.NewSubcribe(SQS)

	BankServiceClient = bankserviceclient.NewBankServiceClient(httpClient, *bankClientConfig)

	TransferRepo = repository.NewTransfersRepository(database)
	UserBalanceRepo = repository.NewUserBalancesRepository(database)

	BankAccountUsecase = usecase.NewBankAccountUsecase(BankServiceClient)
	CallbackUsecase = usecase.NewCallbackUsecase(TransferRepo, Publisher)
	TransferUsecase = usecase.NewTransferUsecase(TransferRepo, BankServiceClient, Publisher)
	UserBalanceUsecase = usecase.NewUserBalanceUsecase(TransferRepo, UserBalanceRepo, Publisher)

	DeductBalanceWorker = worker.NewDeductBalanceWorker(UserBalanceUsecase)
	ProceedTransferWorker = worker.NewProceedTransferWorker(TransferUsecase)
	ReversalBalanceWorker = worker.NewReversalBalanceWorker(UserBalanceUsecase)
}

package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	// echoOpenAPI "github.com/alexferl/echo-openapi"
	httpHandler "github.com/fahmyabida/brick-transfer/pkg/http/handler"
	customMiddleware "github.com/fahmyabida/brick-transfer/pkg/http/middleware"
)

var restCommand = &cobra.Command{
	Use:   "rest",
	Short: "Start REST server",
	Run:   RunRestServer,
}

func init() {
	rootCmd.AddCommand(restCommand)
}

func RunRestServer(cmd *cobra.Command, args []string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Use(customMiddleware.ErrorMiddleware())

	// e.Use(echoOpenAPI.OpenAPI("./docs/openapi.yaml"))

	healthcheckGroup := e.Group("/healthcheck")
	httpHandler.InitHealthcheckHandler(healthcheckGroup)

	v1 := e.Group("/api/v1")
	httpHandler.InitBankAccountHandler(v1, BankAccountUsecase)
	httpHandler.InitTransferHandler(v1, TransferUsecase)
	httpHandler.InitCallbackHandler(v1, CallbackUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}

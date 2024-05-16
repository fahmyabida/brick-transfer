package handler

import (
	"net/http"

	"github.com/fahmyabida/brick-transfer/internal/app/domain"

	"github.com/labstack/echo/v4"
)

type TransferHandler struct {
	transferUsecase domain.ITransferUsecase
}

func InitTransferHandler(e *echo.Group, transferUsecase domain.ITransferUsecase) {
	handler := TransferHandler{transferUsecase: transferUsecase}

	e.POST("/transfer", handler.CreateTransferHandler)
}

func (h *TransferHandler) CreateTransferHandler(c echo.Context) error {

	var transfer domain.Transfers

	if err := c.Bind(&transfer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	ctx := c.Request().Context()

	err := h.transferUsecase.CreateTransfers(ctx, &transfer)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, transfer)
}

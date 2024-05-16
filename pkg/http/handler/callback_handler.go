package handler

import (
	"net/http"

	"github.com/fahmyabida/brick-transfer/internal/app/domain"

	"github.com/labstack/echo/v4"
)

type CallbackHandler struct {
	callbackUsecase domain.ICallbackUsecase
}

func InitCallbackHandler(e *echo.Group, callbackUsecase domain.ICallbackUsecase) {
	handler := CallbackHandler{callbackUsecase: callbackUsecase}

	e.POST("/callbacks/transfer", handler.TransferCallbackHandler)
}

func (h *CallbackHandler) TransferCallbackHandler(c echo.Context) error {

	var callback domain.Callbacks

	if err := c.Bind(&callback); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	ctx := c.Request().Context()

	err := h.callbackUsecase.TransferCallback(ctx, &callback)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, callback)
}

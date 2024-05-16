package handler

import (
	"net/http"

	"github.com/fahmyabida/brick-transfer/internal/app/domain"

	"github.com/labstack/echo/v4"
)

type BankAccountHandler struct {
	bankAccountUsecase domain.IBankAccountUsecase
}

func InitBankAccountHandler(e *echo.Group, bankAccountUsecase domain.IBankAccountUsecase) {
	handler := BankAccountHandler{bankAccountUsecase: bankAccountUsecase}

	e.POST("/bank-account/validate", handler.ValidateBankAccontHandler)
}

func (h *BankAccountHandler) ValidateBankAccontHandler(c echo.Context) error {

	var bankAccount domain.BankAccount

	if err := c.Bind(&bankAccount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	ctx := c.Request().Context()

	err := h.bankAccountUsecase.Validate(ctx, &bankAccount)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, bankAccount)
}

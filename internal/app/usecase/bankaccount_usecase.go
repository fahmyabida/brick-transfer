package usecase

import (
	"context"

	"github.com/fahmyabida/brick-transfer/internal/app/domain"
	"github.com/fahmyabida/brick-transfer/pkg/external/client/bankserviceclient"
)

type BankAccountUsecaseImpl struct {
	bankServiceClient bankserviceclient.IBankServiceClient
}

func NewBankAccountUsecase(bankServiceClient bankserviceclient.IBankServiceClient) domain.IBankAccountUsecase {
	return &BankAccountUsecaseImpl{
		bankServiceClient,
	}
}

func (u *BankAccountUsecaseImpl) Validate(ctx context.Context, payload *domain.BankAccount) error {
	clientResponseData, err := u.bankServiceClient.ValidateBankAccount(ctx, bankserviceclient.ValidateBankAccountRequest{
		AccountNumber: payload.AccountNumber,
		BankCode:      payload.BankCode,
	})
	if err != nil {
		return err
	}
	response := payload
	response.BankName = clientResponseData.BankName
	response.Valid = clientResponseData.Valid
	return nil
}

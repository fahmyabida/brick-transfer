package usecase

import (
	"context"
	"fmt"

	"github.com/fahmyabida/brick-transfer/internal/app/domain"
	"github.com/fahmyabida/brick-transfer/pkg/external/client/bankserviceclient"
)

type TransferUsecaseImpl struct {
	transferRepo      domain.ITransferRepo
	bankServiceClient bankserviceclient.IBankServiceClient
	publisher         domain.IPublisher
}

func NewTransferUsecase(
	transferRepo domain.ITransferRepo,
	bankServiceClient bankserviceclient.IBankServiceClient,
	publisher domain.IPublisher,
) domain.ITransferUsecase {
	return &TransferUsecaseImpl{
		transferRepo,
		bankServiceClient,
		publisher,
	}
}

func (u *TransferUsecaseImpl) CreateTransfers(ctx context.Context, data *domain.Transfers) error {
	data.Status = string(domain.Accepted)
	err := u.transferRepo.Create(ctx, data)
	if err != nil {
		return err
	}
	return u.publisher.PublishAcceptedTransfer(*data)
}

func (u *TransferUsecaseImpl) ProceedTransfer(ctx context.Context, data *domain.Transfers) error {
	transferData, err := u.transferRepo.FindByID(ctx, data.ID)
	if err != nil {
		return err
	}

	if transferData.Status != string(domain.Deducted) {
		return fmt.Errorf("status transfer is invalid ('%v') while proceed transfer with id '%v'",
			transferData.Status,
			transferData.ID)
	}

	bankClientResponse, err := u.bankServiceClient.TransferMoney(ctx, bankserviceclient.TransferMoneyRequest{
		Amount:        transferData.Amount,
		AccountNumber: transferData.DestinationAccount,
		RecipientName: transferData.TransferMetadata.Recipient.Name,
	})
	if err != nil {
		return err
	}

	transferData.BankTransferID = bankClientResponse.TransferID
	transferData.Status = bankClientResponse.Status
	return u.transferRepo.UpdateByID(ctx, &transferData)
}

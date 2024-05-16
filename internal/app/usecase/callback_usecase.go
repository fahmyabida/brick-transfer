package usecase

import (
	"context"

	"github.com/fahmyabida/brick-transfer/internal/app/domain"
)

type CallbackUsecaseImpl struct {
	transferRepo domain.ITransferRepo
	publisher    domain.IPublisher
}

func NewCallbackUsecase(transferRepo domain.ITransferRepo, publisher domain.IPublisher) domain.ICallbackUsecase {
	return &CallbackUsecaseImpl{
		transferRepo,
		publisher,
	}
}

/*
CreateCallbacks is to handle callback transfer from the bank
 1. check data availablity on our end
 2. check status (need reversal/not) & update it
*/
func (u *CallbackUsecaseImpl) TransferCallback(ctx context.Context, payload *domain.Callbacks) error {

	transferData, err := u.transferRepo.FindByBankTransferID(ctx, payload.TransferID)
	if err != nil {
		return err
	}

	if transferData.Status != string(domain.Proceed) {
		return nil
	}

	if payload.Status != string(domain.Success) {
		u.publisher.PublishReversalTransfer(transferData)
	}

	transferData.Status = payload.Status
	err = u.transferRepo.UpdateByID(ctx, &transferData)
	if err != nil {
		return err
	}
	return nil
}

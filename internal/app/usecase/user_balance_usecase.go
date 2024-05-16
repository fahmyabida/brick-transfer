package usecase

import (
	"context"
	"fmt"

	"github.com/fahmyabida/brick-transfer/internal/app/domain"
	pkgErrors "github.com/fahmyabida/brick-transfer/pkg/errors"
)

type UserBalanceUsecaseImpl struct {
	transferRepo    domain.ITransferRepo
	userBalanceRepo domain.IUserBalanceRepo
	publisher       domain.IPublisher
}

func NewUserBalanceUsecase(
	transferRepo domain.ITransferRepo,
	userBalanceRepo domain.IUserBalanceRepo,
	publisher domain.IPublisher,
) domain.IUserBalanceUsecase {
	return &UserBalanceUsecaseImpl{
		transferRepo:    transferRepo,
		userBalanceRepo: userBalanceRepo,
		publisher:       publisher,
	}
}

func (u *UserBalanceUsecaseImpl) DeductBalance(ctx context.Context, payload *domain.Transfers) error {

	transferData, err := u.transferRepo.FindByID(ctx, payload.ID)
	if err != nil {
		return err
	}

	if transferData.Status != string(domain.Accepted) {
		return fmt.Errorf("status transfer is invalid ('%v') while proceed transfer with id '%v'",
			transferData.Status,
			transferData.ID)
	}
	_, err = u.userBalanceRepo.DeductBalanceByUserIdAndCurrency(ctx, transferData.UserID, transferData.Currency, transferData.Amount)
	if err != nil {
		if err == pkgErrors.InsufficientBalanceError(pkgErrors.InsufficientBalance) {
			transferData.Status = string(domain.Rejected)
			err = u.transferRepo.UpdateByID(ctx, &transferData)
			if err != nil {
				return err
			}
		}
		return err
	}

	transferData.Status = string(domain.Deducted)
	err = u.transferRepo.UpdateByID(ctx, &transferData)
	if err != nil {
		return err
	}

	return u.publisher.PublishProceedTransfer(transferData)
}

func (u *UserBalanceUsecaseImpl) ReversalBalance(ctx context.Context, payload *domain.Transfers) error {
	transferData, err := u.transferRepo.FindByID(ctx, payload.ID)
	if err != nil {
		return err
	}
	_, err = u.userBalanceRepo.ReversalBalanceByUserIdAndCurrency(ctx, transferData.UserID, transferData.Currency, transferData.Amount)
	if err != nil {
		return err
	}

	return nil
}

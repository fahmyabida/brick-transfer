package repository

import (
	"context"
	"fmt"

	"github.com/fahmyabida/brick-transfer/internal/app/domain"
	pkgErrors "github.com/fahmyabida/brick-transfer/pkg/errors"
	"gorm.io/gorm"
)

// TransfersRepository ...
type UserBalancesRepository struct {
	DB *gorm.DB
}

// NewUserBalancesRepository will return
func NewUserBalancesRepository(db *gorm.DB) *UserBalancesRepository {
	return &UserBalancesRepository{
		DB: db,
	}
}

func (r UserBalancesRepository) DeductBalanceByUserIdAndCurrency(ctx context.Context, userId, currency string, deductedAmount float64) (
	data domain.UserBalances, err error) {

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx).Where("user_id = ? AND currency = ?", userId, currency).First(&data)
		if tx.Error != nil {
			if tx.Error == gorm.ErrRecordNotFound {
				err = pkgErrors.UserBalancesNotFoundError(fmt.Sprintf(pkgErrors.ErrUserBalanceNotFound, data.ID))
				return err
			}
			err = tx.Error
			return err
		}

		if data.Amount < deductedAmount {
			err = pkgErrors.InsufficientBalanceError(pkgErrors.InsufficientBalance)
			return err
		}

		data.Amount -= deductedAmount
		tx = tx.WithContext(ctx).Where("id = ?", data.ID).Updates(data)
		if tx.Error != nil {
			err = pkgErrors.UserBalanceUpdateFailedError(
				fmt.Sprintf(pkgErrors.ErrUserBalanceUpdateFailed, data.ID, tx.Error.Error()))
			return err
		}

		if tx.RowsAffected == 0 {
			err = pkgErrors.UserBalancesNotFoundError(fmt.Sprintf(pkgErrors.ErrUserBalanceNotFound, data.ID))
			return err
		}

		return nil
	})

	return data, err

}

func (r UserBalancesRepository) ReversalBalanceByUserIdAndCurrency(ctx context.Context, userId, currency string, deductedAmount float64) (
	data domain.UserBalances, err error) {

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx).Where("user_id = ? AND currency = ?", userId, currency).First(&data)
		if tx.Error != nil {
			if tx.Error == gorm.ErrRecordNotFound {
				err = pkgErrors.UserBalancesNotFoundError(fmt.Sprintf(pkgErrors.ErrUserBalanceNotFound, data.ID))
				return err
			}
			err = tx.Error
			return err
		}

		data.Amount += deductedAmount
		tx = tx.WithContext(ctx).Where("id = ?", data.ID).Updates(data)
		if tx.Error != nil {
			err = pkgErrors.UserBalanceUpdateFailedError(
				fmt.Sprintf(pkgErrors.ErrUserBalanceUpdateFailed, data.ID, tx.Error.Error()))
			return err
		}

		if tx.RowsAffected == 0 {
			err = pkgErrors.UserBalancesNotFoundError(fmt.Sprintf(pkgErrors.ErrUserBalanceNotFound, data.ID))
			return err
		}

		return nil
	})

	return data, err

}

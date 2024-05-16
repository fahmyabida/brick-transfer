package repository

import (
	"context"
	"fmt"

	"github.com/fahmyabida/brick-transfer/internal/app/domain"
	pkgErrors "github.com/fahmyabida/brick-transfer/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// TransfersRepository ...
type TransfersRepository struct {
	DB *gorm.DB
}

// NewTransfersRepository will return
func NewTransfersRepository(db *gorm.DB) *TransfersRepository {
	return &TransfersRepository{
		DB: db,
	}
}

func (r TransfersRepository) Create(ctx context.Context, data *domain.Transfers) (err error) {
	err = data.Serialize()
	if err != nil {
		return err
	}

	dbResult := r.DB.WithContext(ctx).Create(data)
	if dbResult.Error != nil {
		// https://www.postgresql.org/docs/current/errcodes-appendix.html
		postgresError, ok := dbResult.Error.(*pgconn.PgError)
		if ok && postgresError.Code == "23505" {
			return pkgErrors.DuplicateTransferError(pkgErrors.ErrDuplicateTransfer)
		}
		return dbResult.Error
	}

	return
}

func (r TransfersRepository) FindByID(ctx context.Context, id string) (data domain.Transfers, err error) {
	dbResult := r.DB.Model(&data).Where("id = ?", id).Find(&data)
	if dbResult.RowsAffected == 0 {
		err = pkgErrors.TransferNotFoundError(fmt.Sprintf(pkgErrors.ErrTransferNotFound, data.ID))
		return
	}

	err = data.Deserialize()

	return
}

func (r TransfersRepository) UpdateByID(ctx context.Context, data *domain.Transfers) (err error) {
	dbResult := r.DB.WithContext(ctx).Where("id = ?", data.ID).Updates(data)
	if dbResult.Error != nil {
		err = pkgErrors.UserBalanceUpdateFailedError(
			fmt.Sprintf(pkgErrors.ErrTransferUpdateFailed, data.ID, dbResult.Error.Error()))
		return err
	}

	if dbResult.RowsAffected == 0 {
		err = pkgErrors.TransferNotFoundError(fmt.Sprintf(pkgErrors.ErrTransferNotFound, data.ID))
		return err
	}

	return nil
}

func (r TransfersRepository) FindByBankTransferID(ctx context.Context, bankTransferId string) (data domain.Transfers, err error) {
	dbResult := r.DB.Model(&data).Where("bank_transfer_id = ?", bankTransferId).Find(&data)
	if dbResult.RowsAffected == 0 {
		err = pkgErrors.TransferNotFoundError(fmt.Sprintf(pkgErrors.ErrTransferNotFound, data.ID))
		return
	}

	err = data.Deserialize()

	return
}

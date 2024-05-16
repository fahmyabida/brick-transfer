package domain

import (
	"context"
	"time"
)

type UserBalances struct {
	ID        string     `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4();PRIMARY_KEY"`
	UserID    string     `gorm:"column:user_id;NOT NULL"`
	Currency  string     `gorm:"column:currency;NOT NULL"`
	Amount    float64    `gorm:"column:amount"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
}

type IUserBalanceRepo interface {
	DeductBalanceByUserIdAndCurrency(ctx context.Context, userId, currency string, deductedAmount float64) (UserBalances, error)
	ReversalBalanceByUserIdAndCurrency(ctx context.Context, userId, currency string, amount float64) (UserBalances, error)
}

type IUserBalanceUsecase interface {
	DeductBalance(ctx context.Context, data *Transfers) error
	ReversalBalance(ctx context.Context, data *Transfers) error
}

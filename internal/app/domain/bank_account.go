package domain

import "context"

type BankAccount struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	BankCode      string `json:"bank_code"`
	Valid         bool   `json:"valid"`
}

type IBankAccountUsecase interface {
	Validate(ctx context.Context, payload *BankAccount) error
}

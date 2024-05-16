package domain

import "context"

type Callbacks struct {
	TransferID    string  `json:"transfer_id"`
	Status        string  `json:"status"`
	Message       string  `json:"message"`
	Amount        float64 `json:"amount"`
	AccountNumber string  `json:"account_number"`
	RecipientName string  `json:"recipient_name"`
}

type ICallbackUsecase interface {
	TransferCallback(ctx context.Context, data *Callbacks) error
}

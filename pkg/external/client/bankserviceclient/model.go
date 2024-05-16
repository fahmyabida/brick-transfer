package bankserviceclient

type ValidateBankAccountRequest struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
}

type ValidateBankAccountResponse struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	BankName      string `json:"bank_name"`
	Valid         bool   `json:"valid"`
}

type TransferMoneyRequest struct {
	Amount        float64 `json:"amount"`
	AccountNumber string  `json:"account_number"`
	RecipientName string  `json:"recipient_name"`
}

type TransferMoneyResponse struct {
	TransferID    string `json:"transfer_id"`
	Status        string `json:"status"`
	Message       string `json:"message"`
	Amount        int    `json:"amount"`
	AccountNumber string `json:"account_number"`
	RecipientName string `json:"recipient_name"`
}

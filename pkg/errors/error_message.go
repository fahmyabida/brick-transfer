package errors

const (
	ErrDuplicateTransfer       = "Transfer is duplicated, use different reference_id"
	ErrTransferNotFound        = "User Balance is not found with id '%v'"
	ErrTransferUpdateFailed    = "Error while update transfer with id '%v' & error '%v'"
	ErrUserBalanceNotFound     = "User Balance is not found with id '%v'"
	ErrUserBalanceUpdateFailed = "Error while update user balance with id '%v' & error '%v'"
	InsufficientBalance        = "Your balance is insufficient"
)

package errors

import (
	"net/http"
)

type InsufficientBalanceError string

// Error represents the error message
func (e InsufficientBalanceError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e InsufficientBalanceError) ErrCode() string {
	return InsufficientBalance
}

// StatusCode represents the HTTP status code
func (e InsufficientBalanceError) StatusCode() int {
	return http.StatusConflict
}

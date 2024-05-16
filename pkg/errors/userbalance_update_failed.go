package errors

import (
	"net/http"
)

type UserBalanceUpdateFailedError string

// Error represents the error message
func (e UserBalanceUpdateFailedError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e UserBalanceUpdateFailedError) ErrCode() string {
	return UserBalanceUpdateFailed
}

// StatusCode represents the HTTP status code
func (e UserBalanceUpdateFailedError) StatusCode() int {
	return http.StatusInternalServerError
}

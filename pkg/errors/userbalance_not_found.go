package errors

import (
	"net/http"
)

type UserBalancesNotFoundError string

// Error represents the error message
func (e UserBalancesNotFoundError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e UserBalancesNotFoundError) ErrCode() string {
	return UserBalanceNotFound
}

// StatusCode represents the HTTP status code
func (e UserBalancesNotFoundError) StatusCode() int {
	return http.StatusInternalServerError
}

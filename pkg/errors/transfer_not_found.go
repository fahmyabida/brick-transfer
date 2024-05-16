package errors

import (
	"net/http"
)

type TransferNotFoundError string

// Error represents the error message
func (e TransferNotFoundError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e TransferNotFoundError) ErrCode() string {
	return TransferNotFound
}

// StatusCode represents the HTTP status code
func (e TransferNotFoundError) StatusCode() int {
	return http.StatusInternalServerError
}

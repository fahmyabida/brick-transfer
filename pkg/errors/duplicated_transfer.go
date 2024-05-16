package errors

import (
	"net/http"
)

type DuplicateTransferError string

// Error represents the error message
func (e DuplicateTransferError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e DuplicateTransferError) ErrCode() string {
	return DuplicateTransfer
}

// StatusCode represents the HTTP status code
func (e DuplicateTransferError) StatusCode() int {
	return http.StatusConflict
}

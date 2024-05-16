package errors

import (
	"net/http"
)

type TransferUpdateFailedError string

// Error represents the error message
func (e TransferUpdateFailedError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e TransferUpdateFailedError) ErrCode() string {
	return TransferUpdateFailed
}

// StatusCode represents the HTTP status code
func (e TransferUpdateFailedError) StatusCode() int {
	return http.StatusInternalServerError
}

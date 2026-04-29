package gospotify

import (
	"errors"
	"fmt"
)

var (
	// ErrNoActivePlayback is an error indicating a playback is not available or active
	ErrNoActivePlayback = errors.New("playback is not available or active")
)

// ErrorResponse is the error response returned by the Spotify API.
type ErrorResponse struct {
	// ErrorObject is the error object.
	ErrorObject struct {
		// Status is the HTTP status code also returned in the response header.
		Status int `json:"status"`
		// Message is a short description of the cause of the error.
		Message string `json:"message"`
	} `json:"error"`
}

// Error implements the error interface.
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.ErrorObject.Status, e.ErrorObject.Message)
}

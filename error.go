package gospotify

import "fmt"

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

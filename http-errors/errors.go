// Package httperrors
package httperrors

import (
	"errors"
	"fmt"
)

type HTTPError struct {
	statusCode int
	appErr     error
	err        error
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("<%d> %v - %v", e.statusCode, e.appErr, e.err)
}

func NewHTTPErr(statusCode int, appErr error, err error) *HTTPError {
	return &HTTPError{
		statusCode,
		appErr,
		err,
	}
}

func (e HTTPError) StatusCode() int {
	return e.statusCode
}

func (e HTTPError) Err() string {
	if e.err != nil {
		return e.err.Error()
	}

	return ""
}

func FromErr(e error) (*HTTPError, bool) {
	var httpErr HTTPError
	if errors.As(e, &httpErr) {
		return &httpErr, true
	}

	return nil, false
}

// func (e HTTPError) Err() error {
// 	return e.err
// }

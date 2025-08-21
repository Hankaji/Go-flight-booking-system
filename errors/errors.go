// Package errors
package errors

import (
	"fmt"
)

type HTTPError struct {
	statusCode uint
	err        error
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("<%d> %v", e.statusCode, e.err.Error())
}

func NewHTTPErr(statusCode uint, err error) *HTTPError {
	return &HTTPError{
		statusCode,
		err,
	}
}

func (e HTTPError) StatusCode() uint {
	return e.statusCode
}

// func (e HTTPError) Err() error {
// 	return e.err
// }

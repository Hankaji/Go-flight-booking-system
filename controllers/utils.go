// Package controllers
package controllers

import (
	"errors"
	"flight-booking-server/controllers/queries"
	httperrors "flight-booking-server/http-errors"
	"net/http"
)

func ValidatePagination(q queries.PaginationQuery) *httperrors.HTTPError {
	if q.Index < 1 {
		return httperrors.NewHTTPErr(
			http.StatusBadRequest,
			errors.New("index must be at least 1"),
			nil)
	}

	if q.Limit == 0 || q.Limit > 100 {
		return httperrors.NewHTTPErr(
			http.StatusBadRequest,
			errors.New("limit must be between 1 and 100"),
			nil)
	}

	return nil
}

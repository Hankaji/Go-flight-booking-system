// Package repositories
package repositories

import "flight-booking-server/controllers/queries"

type IRepoRead[T any] interface {
	GetAll() ([]T, error)
	GetByID(string) (*T, error)
}

type IRepoCreate[T any] interface {
	Create(any) (*T, error)
}

type Pagination struct {
	index uint
	limit uint
}

func NewPagination(index, limit uint) *Pagination {
	return &Pagination{
		index,
		limit,
	}
}

func PaginationFromQuery(query queries.PaginationQuery) *Pagination {
	return &Pagination{
		index: query.Index,
		limit: query.Limit,
	}
}

func DefaultPagination() *Pagination {
	return &Pagination{
		index: 0,
		limit: 10,
	}
}

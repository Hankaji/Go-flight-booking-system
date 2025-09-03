// Package repositories
package repositories

import (
	"flight-booking-server/controllers/queries"

	"gorm.io/gorm"
)

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

func (p *Pagination) Apply(db *gorm.DB) *gorm.DB {
	if p == nil {
		p = DefaultPagination()
	}

	db.Offset(int((p.index - 1) * p.limit))
	db.Limit(int(p.limit))

	return db
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

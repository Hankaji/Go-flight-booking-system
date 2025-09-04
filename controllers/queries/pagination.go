// Package queries
package queries

type PaginationQuery struct {
	Index uint `form:"index,default=1"`
	Limit uint `form:"limit,default=10"`
}

// func (pq *PaginationQuery) ToPagination() repositories.Pagination {
// 	return &Pagination{
// 		index: pq.Index,
// 		limit: pq.Limit,
// 	}
// }

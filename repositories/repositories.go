// Package repositories
package repositories

type IRepoRead[T any] interface {
	GetAll() ([]T, error)
	GetByID(id string) (*T, error)
}

// type IRepoCreate[T any] interface {
// 	Create(any) (*T, error)
// }

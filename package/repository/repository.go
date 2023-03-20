package repository

import (
	"github.com/Tasma-110110/mid1-prj"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user mid1.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

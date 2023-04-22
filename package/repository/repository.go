package repository

import (
	"github.com/Tasma-110110/mid1-prj"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user mid1.User) (int, error)
}


type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

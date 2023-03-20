package repository

import (
	"fmt"
	"os/user"

	"github.com/Tasma-110110/mid1-prj"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(mid1.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(name,username,password) values ($1, $2, $3) Returning id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (mid1.User, error) {
	var user mid1.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=%1 AND password_hash=$2", userTable)

	return user
}

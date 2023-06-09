package service

import (
	"github.com/Tasma-110110/mid1-prj"
	"github.com/Tasma-110110/mid1-prj/package/repository"
)

type Authorization interface {
	CreateUser(user mid1.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

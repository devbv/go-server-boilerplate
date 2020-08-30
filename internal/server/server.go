package server

import (
	"github.com/devbv/go-server-boilerplate/internal/repository/fake"
	"github.com/devbv/go-server-boilerplate/internal/service"
)

type TodoSever struct {
	svc *service.TodoService
}

//New returns a new instance of TodoSever
func New() *TodoSever {
	repo := fake.NewFakeTaskRepository()
	return &TodoSever{
		svc: service.NewTodoService(repo),
	}
}

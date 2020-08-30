package service

import (
	"context"

	m "github.com/devbv/go-server-boilerplate/internal/model"
	r "github.com/devbv/go-server-boilerplate/internal/repository"
)

type TodoService struct {
	repo r.ITaskRepository
}

func NewTodoService(repo r.ITaskRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (svc *TodoService) GetTaskList(ctx context.Context) ([]*m.Task, error) {
	return svc.repo.List()
}

func (svc *TodoService) AddTask(ctx context.Context, task *m.Task) error {
	return svc.repo.Add(task)
}

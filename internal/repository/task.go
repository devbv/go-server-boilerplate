package repository

import m "github.com/devbv/go-server-boilerplate/internal/model"

type ITaskRepository interface {
	Add(*m.Task) error
	List() ([]*m.Task, error)
}

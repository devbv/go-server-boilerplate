package server

import (
	"context"

	"github.com/devbv/go-server-boilerplate/internal/model"
	todo "github.com/devbv/go-server-boilerplate/pkg/todo"
)

func (ts *TodoSever) GetTaskList(ctx context.Context, _ *todo.GetTaskListRequest) (*todo.GetTaskListResponse, error) {
	list, _ := ts.svc.GetTaskList(ctx)

	ret := make([]*todo.Task, 0)
	for _, task := range list {
		ret = append(ret, &todo.Task{
			Name:   task.Name,
			Detail: task.Detail,
			Due:    "",
		})
	}

	return &todo.GetTaskListResponse{
		List: ret,
	}, nil
}

func (ts *TodoSever) AddTask(ctx context.Context, req *todo.AddTaskRequest) (*todo.AddTaskResponse, error) {
	err := ts.svc.AddTask(ctx, &model.Task{
		Name:   req.Task.Name,
		Detail: req.Task.Detail,
		Due:    nil,
	})
	return &todo.AddTaskResponse{
		Success: err == nil,
	}, nil
}

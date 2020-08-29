package server

import (
	"context"

	rpc "github.com/devbv/go-server-boilerplate/pkg/rpc"
)

func (ts *TODOServer) GetTaskList(context.Context, *rpc.GetTaskListRequest) (*rpc.GetTaskListResponse, error) {
	return nil, nil
}
func (ts *TODOServer) AddTask(context.Context, *rpc.AddTaskRequest) (*rpc.AddTaskResponse, error) {
	return nil, nil
}

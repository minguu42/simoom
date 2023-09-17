package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

var errUnimplemented = connect.NewError(connect.CodeUnimplemented, errors.New("this RPC method is not yet implemented"))

type taskHandler struct {
	repo repository.Repository
}

func (h taskHandler) CreateTask(_ context.Context, _ *connect.Request[simoompb.CreateTaskRequest]) (*connect.Response[simoompb.TaskResponse], error) {
	return nil, errUnimplemented
}

func (h taskHandler) ListTasks(_ context.Context, _ *connect.Request[simoompb.ListTasksRequest]) (*connect.Response[simoompb.TasksResponse], error) {
	return nil, errUnimplemented
}

func (h taskHandler) UpdateTask(_ context.Context, _ *connect.Request[simoompb.UpdateTaskRequest]) (*connect.Response[simoompb.TaskResponse], error) {
	return nil, errUnimplemented
}

func (h taskHandler) DeleteTask(_ context.Context, _ *connect.Request[simoompb.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, errUnimplemented
}

package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func newTaskResponse(t model.Task) *simoompb.TaskResponse {
	return &simoompb.TaskResponse{
		Id:          t.ID,
		ProjectId:   t.ProjectID,
		Steps:       newStepsResponse(t.Steps),
		Tags:        newTagsResponse(t.Tags),
		Title:       t.Title,
		Content:     t.Content,
		Priority:    uint32(t.Priority),
		DueOn:       newDate(t.DueOn),
		CompletedAt: newTimestamp(t.CompletedAt),
		CreatedAt:   timestamppb.New(t.CreatedAt),
		UpdatedAt:   timestamppb.New(t.UpdatedAt),
	}
}

func newTasksResponse(ts []model.Task) []*simoompb.TaskResponse {
	tasks := make([]*simoompb.TaskResponse, 0, len(ts))
	for _, t := range ts {
		tasks = append(tasks, newTaskResponse(t))
	}
	return tasks
}

func (h handler) CreateTask(ctx context.Context, req *connect.Request[simoompb.CreateTaskRequest]) (*connect.Response[simoompb.TaskResponse], error) {
	if len(req.Msg.ProjectId) != 26 {
		return nil, newErrInvalidArgument("project_id is a 26-character string")
	}
	if req.Msg.Title == "" {
		return nil, newErrInvalidArgument("title cannot be an empty string")
	}
	if req.Msg.Priority > 3 {
		return nil, newErrInvalidArgument("priority is specified by 0 to 3")
	}

	out, err := h.task.CreateTask(ctx, usecase.CreateTaskInput{
		ProjectID: req.Msg.ProjectId,
		Title:     req.Msg.Title,
		Priority:  uint(req.Msg.Priority),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newTaskResponse(out.Task)), nil
}

func (h handler) ListTasksByProjectID(ctx context.Context, req *connect.Request[simoompb.ListTasksByProjectIDRequest]) (*connect.Response[simoompb.TasksResponse], error) {
	if len(req.Msg.ProjectId) != 26 {
		return nil, newErrInvalidArgument("project_id is a 26-character string")
	}
	if req.Msg.Limit < 1 {
		return nil, newErrInvalidArgument("limit is greater than or equal to 1")
	}

	out, err := h.task.ListTasksByProjectID(ctx, usecase.ListTasksByProjectIDInput{
		ProjectID: req.Msg.ProjectId,
		Limit:     uint(req.Msg.Limit),
		Offset:    uint(req.Msg.Offset),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&simoompb.TasksResponse{
		Tasks:   newTasksResponse(out.Tasks),
		HasNext: out.HasNext,
	}), nil
}

func (h handler) ListTasksByTagID(ctx context.Context, req *connect.Request[simoompb.ListTasksByTagIDRequest]) (*connect.Response[simoompb.TasksResponse], error) {
	if len(req.Msg.TagId) != 26 {
		return nil, newErrInvalidArgument("tag_id is a 26-character string")
	}
	if req.Msg.Limit < 1 {
		return nil, newErrInvalidArgument("limit is greater than or equal to 1")
	}

	out, err := h.task.ListTasksByTagID(ctx, usecase.ListTasksByTagIDInput{
		TagID:  req.Msg.TagId,
		Limit:  uint(req.Msg.Limit),
		Offset: uint(req.Msg.Offset),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&simoompb.TasksResponse{
		Tasks:   newTasksResponse(out.Tasks),
		HasNext: false,
	}), nil
}

func (h handler) UpdateTask(ctx context.Context, req *connect.Request[simoompb.UpdateTaskRequest]) (*connect.Response[simoompb.TaskResponse], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}
	if req.Msg.Title == nil && req.Msg.Content == nil && req.Msg.Priority == nil && req.Msg.DueOn == nil && req.Msg.CompletedAt == nil {
		return nil, newErrInvalidArgument("must contain some argument other than id")
	}
	if req.Msg.Title != nil && *req.Msg.Title == "" {
		return nil, newErrInvalidArgument("title cannot be an empty string")
	}
	if req.Msg.Priority != nil && *req.Msg.Priority > 3 {
		return nil, newErrInvalidArgument("priority is specified by 0 to 3")
	}

	out, err := h.task.UpdateTask(ctx, usecase.UpdateTaskInput{
		ID:          req.Msg.Id,
		Title:       req.Msg.Title,
		Content:     req.Msg.Content,
		Priority:    nil,
		DueOn:       nil,
		CompletedAt: nil,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newTaskResponse(out.Task)), nil
}

func (h handler) DeleteTask(ctx context.Context, req *connect.Request[simoompb.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}

	if err := h.task.DeleteTask(ctx, usecase.DeleteTaskInput{
		ID: req.Msg.Id,
	}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

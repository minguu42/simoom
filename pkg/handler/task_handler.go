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

type taskHandler struct {
	uc usecase.TaskUsecase
}

func (h taskHandler) CreateTask(ctx context.Context, req *connect.Request[simoompb.CreateTaskRequest]) (*connect.Response[simoompb.TaskResponse], error) {
	if req.Msg.ProjectId != "" {
		return nil, errInvalidArgument
	}
	if req.Msg.Title != "" {
		return nil, errInvalidArgument
	}

	out, err := h.uc.CreateTask(ctx, usecase.CreateTaskInput{
		ProjectID: req.Msg.ProjectId,
		Title:     req.Msg.Title,
		Priority:  uint(req.Msg.Priority),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newTaskResponse(out.Task)), nil
}

func (h taskHandler) ListTasksByProjectID(ctx context.Context, req *connect.Request[simoompb.ListTasksByProjectIDRequest]) (*connect.Response[simoompb.TasksResponse], error) {
	out, err := h.uc.ListTasksByProjectID(ctx, usecase.ListTasksByProjectIDInput{
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

func (h taskHandler) ListTasksByTagID(ctx context.Context, req *connect.Request[simoompb.ListTasksByTagIDRequest]) (*connect.Response[simoompb.TasksResponse], error) {
	out, err := h.uc.ListTasksByTagID(ctx, usecase.ListTasksByTagIDInput{
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

func (h taskHandler) UpdateTask(ctx context.Context, req *connect.Request[simoompb.UpdateTaskRequest]) (*connect.Response[simoompb.TaskResponse], error) {
	if req.Msg.Id == "" {
		return nil, errInvalidArgument
	}
	if req.Msg.Title == nil && req.Msg.Content == nil && req.Msg.Priority == nil && req.Msg.DueOn == nil && req.Msg.CompletedAt == nil {
		return nil, errInvalidArgument
	}

	out, err := h.uc.UpdateTask(ctx, usecase.UpdateTaskInput{
		ID:          req.Msg.Id,
		ProjectID:   req.Msg.ProjectId,
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

func (h taskHandler) DeleteTask(ctx context.Context, req *connect.Request[simoompb.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error) {
	if err := h.uc.DeleteTask(ctx, usecase.DeleteTaskInput{
		ID: req.Msg.Id,
	}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

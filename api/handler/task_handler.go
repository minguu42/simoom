package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newTask(t model.Task) *simoompb.Task {
	return &simoompb.Task{
		Id:          string(t.ID),
		ProjectId:   string(t.ProjectID),
		Steps:       newSteps(t.Steps),
		Tags:        newTags(t.Tags),
		Name:        t.Name,
		Content:     t.Content,
		Priority:    uint32(t.Priority),
		DueOn:       newDate(t.DueOn),
		CompletedAt: newTimestamp(t.CompletedAt),
	}
}

func newTasksResponse(ts []model.Task) []*simoompb.Task {
	tasks := make([]*simoompb.Task, 0, len(ts))
	for _, t := range ts {
		tasks = append(tasks, newTask(t))
	}
	return tasks
}

func (h handler) CreateTask(ctx context.Context, req *connect.Request[simoompb.CreateTaskRequest]) (*connect.Response[simoompb.Task], error) {
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, apperr.ErrInvalidRequest(err)
	}

	out, err := h.task.CreateTask(ctx, usecase.CreateTaskInput{
		ProjectID: model.ProjectID(req.Msg.ProjectId),
		Name:      req.Msg.Name,
		Priority:  uint(req.Msg.Priority),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newTask(out.Task)), nil
}

func (h handler) ListTasks(ctx context.Context, req *connect.Request[simoompb.ListTasksRequest]) (*connect.Response[simoompb.Tasks], error) {
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, apperr.ErrInvalidRequest(err)
	}

	out, err := h.task.ListTasks(ctx, usecase.ListTasksInput{
		Limit:     uint(req.Msg.Limit),
		Offset:    uint(req.Msg.Offset),
		ProjectID: (*model.ProjectID)(req.Msg.ProjectId),
		TagID:     (*model.TagID)(req.Msg.TagId),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&simoompb.Tasks{
		Tasks:   newTasksResponse(out.Tasks),
		HasNext: out.HasNext,
	}), nil
}

func (h handler) UpdateTask(ctx context.Context, req *connect.Request[simoompb.UpdateTaskRequest]) (*connect.Response[simoompb.Task], error) {
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, apperr.ErrInvalidRequest(err)
	}

	out, err := h.task.UpdateTask(ctx, usecase.UpdateTaskInput{
		ID:          model.TaskID(req.Msg.Id),
		Name:        req.Msg.Name,
		Content:     req.Msg.Content,
		Priority:    nil,
		DueOn:       nil,
		CompletedAt: nil,
		TagIDs:      tagIDsOrNil(req.Msg.TagIds, !req.Msg.ShouldUpdateTag),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newTask(out.Task)), nil
}

func (h handler) DeleteTask(ctx context.Context, req *connect.Request[simoompb.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error) {
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, apperr.ErrInvalidRequest(err)
	}

	if err := h.task.DeleteTask(ctx, usecase.DeleteTaskInput{
		ID: model.TaskID(req.Msg.Id),
	}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

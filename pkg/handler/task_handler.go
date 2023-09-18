package handler

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/pkg/domain/idgen"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func newTaskResponse(t model.Task) *simoompb.TaskResponse {
	return &simoompb.TaskResponse{
		Id:          t.ID,
		ProjectId:   t.ProjectID,
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
	repo repository.Repository
}

func (h taskHandler) CreateTask(ctx context.Context, req *connect.Request[simoompb.CreateTaskRequest]) (*connect.Response[simoompb.TaskResponse], error) {
	p, err := h.repo.GetProjectByID(ctx, req.Msg.ProjectId)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errProjectNotFound
		}
		return nil, errInternal
	}
	if p.UserID != userID {
		return nil, errProjectNotFound
	}

	now := time.Now()
	t := model.Task{
		ID:        idgen.Generate(),
		ProjectID: req.Msg.ProjectId,
		Title:     req.Msg.Title,
		Priority:  uint(req.Msg.Priority),
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := h.repo.CreateTask(ctx, t); err != nil {
		return nil, errInternal
	}
	return connect.NewResponse(newTaskResponse(t)), nil
}

func (h taskHandler) ListTasks(ctx context.Context, req *connect.Request[simoompb.ListTasksRequest]) (*connect.Response[simoompb.TasksResponse], error) {
	p, err := h.repo.GetProjectByID(ctx, req.Msg.ProjectId)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errProjectNotFound
		}
		return nil, errInternal
	}
	if p.UserID != userID {
		return nil, errProjectNotFound
	}

	ts, err := h.repo.ListTasksByProjectID(ctx, req.Msg.ProjectId, uint(req.Msg.Limit), uint(req.Msg.Offset))
	if err != nil {
		return nil, errInternal
	}

	return connect.NewResponse(&simoompb.TasksResponse{
		Tasks:   newTasksResponse(ts),
		HasNext: false,
	}), nil
}

func (h taskHandler) UpdateTask(ctx context.Context, req *connect.Request[simoompb.UpdateTaskRequest]) (*connect.Response[simoompb.TaskResponse], error) {
	p, err := h.repo.GetProjectByID(ctx, req.Msg.ProjectId)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errProjectNotFound
		}
		return nil, errInternal
	}
	if p.UserID != userID {
		return nil, errProjectNotFound
	}

	t, err := h.repo.GetTaskByID(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errTaskNotFound
		}
		return nil, errInternal
	}
	if !p.ContainsTask(t) {
		return nil, errTaskNotFound
	}

	if req.Msg.Title == nil && req.Msg.Content == nil && req.Msg.Priority == nil &&
		req.Msg.DueOn == nil && req.Msg.CompletedAt == nil {
		return nil, errInvalidArgument
	}
	if req.Msg.Title != nil {
		t.Title = *req.Msg.Title
	}
	if req.Msg.Content != nil {
		t.Content = *req.Msg.Content
	}
	if req.Msg.Priority != nil {
		t.Priority = uint(*req.Msg.Priority)
	}
	if req.Msg.DueOn != nil {
		dueOn := time.Date(int(req.Msg.DueOn.Year), time.Month(req.Msg.DueOn.Month), int(req.Msg.DueOn.Day), 0, 0, 0, 0, time.UTC)
		t.DueOn = &dueOn
	}
	if req.Msg.CompletedAt != nil {
		completedAt := req.Msg.CompletedAt.AsTime()
		t.CompletedAt = &completedAt
	}

	if err := h.repo.UpdateTask(ctx, t); err != nil {
		return nil, errInternal
	}

	return connect.NewResponse(newTaskResponse(t)), nil
}

func (h taskHandler) DeleteTask(ctx context.Context, req *connect.Request[simoompb.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error) {
	p, err := h.repo.GetProjectByID(ctx, req.Msg.ProjectId)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errProjectNotFound
		}
		return nil, errInternal
	}
	if p.UserID != userID {
		return nil, errProjectNotFound
	}

	t, err := h.repo.GetTaskByID(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errTaskNotFound
		}
		return nil, errInternal
	}
	if !p.ContainsTask(t) {
		return nil, errTaskNotFound
	}

	if err := h.repo.DeleteTask(ctx, req.Msg.Id); err != nil {
		return nil, errInternal
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

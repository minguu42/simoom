package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/minguu42/simoom/pkg/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func newStep(s model.Step) *simoompb.Step {
	return &simoompb.Step{
		Id:          s.ID,
		TaskId:      s.TaskID,
		Title:       s.Title,
		CompletedAt: newTimestamp(s.CompletedAt),
		CreatedAt:   timestamppb.New(s.CreatedAt),
		UpdatedAt:   timestamppb.New(s.UpdatedAt),
	}
}

func newSteps(ss []model.Step) []*simoompb.Step {
	steps := make([]*simoompb.Step, 0, len(ss))
	for _, s := range ss {
		steps = append(steps, newStep(s))
	}
	return steps
}

func (h handler) CreateStep(ctx context.Context, req *connect.Request[simoompb.CreateStepRequest]) (*connect.Response[simoompb.Step], error) {
	if len(req.Msg.TaskId) != 26 {
		return nil, newErrInvalidArgument("task_id is a 26-character string")
	}
	if req.Msg.Title == "" {
		return nil, newErrInvalidArgument("title cannot be an empty string")
	}

	out, err := h.step.CreateStep(ctx, usecase.CreateStepInput{
		TaskID: req.Msg.TaskId,
		Title:  req.Msg.Title,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newStep(out.Step)), nil
}

func (h handler) UpdateStep(ctx context.Context, req *connect.Request[simoompb.UpdateStepRequest]) (*connect.Response[simoompb.Step], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}
	if req.Msg.Title == nil && req.Msg.CompletedAt == nil {
		return nil, newErrInvalidArgument("must contain some argument other than id")
	}
	if req.Msg.Title != nil && *req.Msg.Title == "" {
		return nil, newErrInvalidArgument("title cannot be an empty string")
	}

	c := req.Msg.CompletedAt.AsTime()
	out, err := h.step.UpdateStep(ctx, usecase.UpdateStepInput{
		ID:          req.Msg.Id,
		Title:       req.Msg.Title,
		CompletedAt: &c,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newStep(out.Step)), nil
}

func (h handler) DeleteStep(ctx context.Context, req *connect.Request[simoompb.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}

	if err := h.step.DeleteStep(ctx, usecase.DeleteStepInput{ID: req.Msg.Id}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

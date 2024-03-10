package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/pointers"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newStep(s model.Step) *simoompb.Step {
	return &simoompb.Step{
		Id:          s.ID,
		TaskId:      s.TaskID,
		Name:        s.Name,
		CompletedAt: newTimestamp(s.CompletedAt),
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
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, apperr.ErrInvalidRequest(err)
	}

	out, err := h.step.CreateStep(ctx, usecase.CreateStepInput{
		TaskID: req.Msg.TaskId,
		Name:   req.Msg.Name,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newStep(out.Step)), nil
}

func (h handler) UpdateStep(ctx context.Context, req *connect.Request[simoompb.UpdateStepRequest]) (*connect.Response[simoompb.Step], error) {
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, apperr.ErrInvalidRequest(err)
	}

	out, err := h.step.UpdateStep(ctx, usecase.UpdateStepInput{
		ID:          req.Msg.Id,
		Name:        req.Msg.Name,
		CompletedAt: pointers.Ref(req.Msg.CompletedAt.AsTime()),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newStep(out.Step)), nil
}

func (h handler) DeleteStep(ctx context.Context, req *connect.Request[simoompb.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error) {
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, apperr.ErrInvalidRequest(err)
	}

	if err := h.step.DeleteStep(ctx, usecase.DeleteStepInput{ID: req.Msg.Id}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

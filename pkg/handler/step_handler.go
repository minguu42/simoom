package handler

import (
	"context"

	"github.com/minguu42/simoom/pkg/usecase"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/pkg/domain/model"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func newStepResponse(s model.Step) *simoompb.StepResponse {
	return &simoompb.StepResponse{
		Id:          s.ID,
		TaskId:      s.TaskID,
		Title:       s.Title,
		CompletedAt: newTimestamp(s.CompletedAt),
		CreatedAt:   timestamppb.New(s.CreatedAt),
		UpdatedAt:   timestamppb.New(s.UpdatedAt),
	}
}

func newStepsResponse(ss []model.Step) []*simoompb.StepResponse {
	steps := make([]*simoompb.StepResponse, 0, len(ss))
	for _, s := range ss {
		steps = append(steps, newStepResponse(s))
	}
	return steps
}

type stepHandler struct {
	uc usecase.StepUsecase
}

func (h stepHandler) CreateStep(ctx context.Context, req *connect.Request[simoompb.CreateStepRequest]) (*connect.Response[simoompb.StepResponse], error) {
	if len(req.Msg.TaskId) != 26 {
		return nil, newErrInvalidArgument("task_id is a 26-character string")
	}
	if req.Msg.Title == "" {
		return nil, newErrInvalidArgument("title cannot be an empty string")
	}

	out, err := h.uc.CreateStep(ctx, usecase.CreateStepInput{
		TaskID: req.Msg.TaskId,
		Title:  req.Msg.Title,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newStepResponse(out.Step)), nil
}

func (h stepHandler) UpdateStep(ctx context.Context, req *connect.Request[simoompb.UpdateStepRequest]) (*connect.Response[simoompb.StepResponse], error) {
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
	out, err := h.uc.UpdateStep(ctx, usecase.UpdateStepInput{
		ID:          req.Msg.Id,
		Title:       req.Msg.Title,
		CompletedAt: &c,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newStepResponse(out.Step)), nil
}

func (h stepHandler) DeleteStep(ctx context.Context, req *connect.Request[simoompb.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}

	if err := h.uc.DeleteStep(ctx, usecase.DeleteStepInput{ID: req.Msg.Id}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

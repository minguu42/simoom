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
	repo repository.Repository
}

func (h stepHandler) CreateStep(ctx context.Context, req *connect.Request[simoompb.CreateStepRequest]) (*connect.Response[simoompb.StepResponse], error) {
	now := time.Now()
	s := model.Step{
		ID:          idgen.Generate(),
		UserID:      userID,
		TaskID:      req.Msg.TaskId,
		Title:       req.Msg.Title,
		CompletedAt: nil,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := h.repo.CreateStep(ctx, s); err != nil {
		return nil, errInternal
	}
	return connect.NewResponse(newStepResponse(s)), nil
}

func (h stepHandler) UpdateStep(ctx context.Context, req *connect.Request[simoompb.UpdateStepRequest]) (*connect.Response[simoompb.StepResponse], error) {
	s, err := h.repo.GetStepByID(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errStepNotFound
		}
		return nil, errInternal
	}
	if userID != s.UserID {
		return nil, errStepNotFound
	}

	if req.Msg.Title == nil && req.Msg.CompletedAt == nil {
		return nil, errInvalidArgument
	}
	if req.Msg.Title != nil {
		s.Title = *req.Msg.Title
	}
	if req.Msg.CompletedAt != nil {
		c := req.Msg.CompletedAt.AsTime()
		s.CompletedAt = &c
	}

	if err := h.repo.UpdateStep(ctx, s); err != nil {
		return nil, errInternal
	}
	return connect.NewResponse(newStepResponse(s)), nil
}

func (h stepHandler) DeleteStep(ctx context.Context, req *connect.Request[simoompb.DeleteStepRequest]) (*connect.Response[emptypb.Empty], error) {
	s, err := h.repo.GetStepByID(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errStepNotFound
		}
		return nil, errInternal
	}
	if userID != s.UserID {
		return nil, errStepNotFound
	}

	if err := h.repo.DeleteStep(ctx, req.Msg.Id); err != nil {
		return nil, errInternal
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

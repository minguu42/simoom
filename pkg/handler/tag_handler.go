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

func newTagResponse(t model.Tag) *simoompb.TagResponse {
	return &simoompb.TagResponse{
		Id:        t.ID,
		Name:      t.Name,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

func newTagsResponse(ts []model.Tag) []*simoompb.TagResponse {
	tags := make([]*simoompb.TagResponse, 0, len(ts))
	for _, t := range ts {
		tags = append(tags, newTagResponse(t))
	}
	return tags
}

type tagHandler struct {
	uc usecase.TagUsecase
}

func (h tagHandler) CreateTag(ctx context.Context, req *connect.Request[simoompb.CreateTagRequest]) (*connect.Response[simoompb.TagResponse], error) {
	if req.Msg.Name == "" {
		return nil, newErrInvalidArgument("name cannot be an empty string")
	}

	out, err := h.uc.CreateTag(ctx, usecase.CreateTagInput{
		Name: req.Msg.Name,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newTagResponse(out.Tag)), nil
}

func (h tagHandler) ListTags(ctx context.Context, req *connect.Request[simoompb.ListTagsRequest]) (*connect.Response[simoompb.TagsResponse], error) {
	if req.Msg.Limit < 1 {
		return nil, newErrInvalidArgument("limit is greater than or equal to 1")
	}

	out, err := h.uc.ListTags(ctx, usecase.ListTagsInput{
		Limit:  uint(req.Msg.Limit),
		Offset: uint(req.Msg.Offset),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&simoompb.TagsResponse{
		Tags: newTagsResponse(out.Tags),
	}), nil
}

func (h tagHandler) UpdateTag(ctx context.Context, req *connect.Request[simoompb.UpdateTagRequest]) (*connect.Response[simoompb.TagResponse], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}
	if req.Msg.Name == nil {
		return nil, newErrInvalidArgument("must contain some argument other than id")
	}
	if req.Msg.Name != nil && *req.Msg.Name == "" {
		return nil, newErrInvalidArgument("name cannot be an empty string")
	}

	out, err := h.uc.UpdateTag(ctx, usecase.UpdateTagInput{
		ID:   req.Msg.Id,
		Name: req.Msg.Name,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newTagResponse(out.Tag)), nil
}

func (h tagHandler) DeleteTag(ctx context.Context, req *connect.Request[simoompb.DeleteTagRequest]) (*connect.Response[emptypb.Empty], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}

	if err := h.uc.DeleteTag(ctx, usecase.DeleteTagInput{ID: req.Msg.Id}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

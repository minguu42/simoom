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
	out, err := h.uc.CreateTag(ctx, usecase.CreateTagInput{
		Name: req.Msg.Name,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newTagResponse(out.Tag)), nil
}

func (h tagHandler) ListTags(ctx context.Context, _ *connect.Request[simoompb.ListTagsRequest]) (*connect.Response[simoompb.TagsResponse], error) {
	out, err := h.uc.ListTags(ctx, usecase.ListTagsInput{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&simoompb.TagsResponse{
		Tags: newTagsResponse(out.Tags),
	}), nil
}

func (h tagHandler) UpdateTag(ctx context.Context, req *connect.Request[simoompb.UpdateTagRequest]) (*connect.Response[simoompb.TagResponse], error) {
	if req.Msg.Name == nil {
		return nil, errInvalidArgument
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
	if err := h.uc.DeleteTag(ctx, usecase.DeleteTagInput{ID: req.Msg.Id}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

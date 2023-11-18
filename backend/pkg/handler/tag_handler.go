package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/backend/pkg/domain/model"
	"github.com/minguu42/simoom/backend/pkg/usecase"
	"github.com/minguu42/simoom/library/simoompb/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func newTag(t model.Tag) *simoompb.Tag {
	return &simoompb.Tag{
		Id:        t.ID,
		Name:      t.Name,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

func newTags(ts []model.Tag) []*simoompb.Tag {
	tags := make([]*simoompb.Tag, 0, len(ts))
	for _, t := range ts {
		tags = append(tags, newTag(t))
	}
	return tags
}

func (h handler) CreateTag(ctx context.Context, req *connect.Request[simoompb.CreateTagRequest]) (*connect.Response[simoompb.Tag], error) {
	if req.Msg.Name == "" {
		return nil, newErrInvalidArgument("name cannot be an empty string")
	}

	out, err := h.tag.CreateTag(ctx, usecase.CreateTagInput{
		Name: req.Msg.Name,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newTag(out.Tag)), nil
}

func (h handler) ListTags(ctx context.Context, req *connect.Request[simoompb.ListTagsRequest]) (*connect.Response[simoompb.Tags], error) {
	if req.Msg.Limit < 1 {
		return nil, newErrInvalidArgument("limit is greater than or equal to 1")
	}

	out, err := h.tag.ListTags(ctx, usecase.ListTagsInput{
		Limit:  uint(req.Msg.Limit),
		Offset: uint(req.Msg.Offset),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&simoompb.Tags{
		Tags:    newTags(out.Tags),
		HasNext: out.HasNext,
	}), nil
}

func (h handler) UpdateTag(ctx context.Context, req *connect.Request[simoompb.UpdateTagRequest]) (*connect.Response[simoompb.Tag], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}
	if req.Msg.Name == nil {
		return nil, newErrInvalidArgument("must contain some argument other than id")
	}
	if req.Msg.Name != nil && *req.Msg.Name == "" {
		return nil, newErrInvalidArgument("name cannot be an empty string")
	}

	out, err := h.tag.UpdateTag(ctx, usecase.UpdateTagInput{
		ID:   req.Msg.Id,
		Name: req.Msg.Name,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newTag(out.Tag)), nil
}

func (h handler) DeleteTag(ctx context.Context, req *connect.Request[simoompb.DeleteTagRequest]) (*connect.Response[emptypb.Empty], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}

	if err := h.tag.DeleteTag(ctx, usecase.DeleteTagInput{ID: req.Msg.Id}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

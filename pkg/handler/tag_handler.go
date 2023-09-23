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
	repo repository.Repository
}

func (h tagHandler) CreateTag(ctx context.Context, req *connect.Request[simoompb.CreateTagRequest]) (*connect.Response[simoompb.TagResponse], error) {
	now := time.Now()
	t := model.Tag{
		ID:        idgen.Generate(),
		UserID:    userID,
		Name:      req.Msg.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := h.repo.CreateTag(ctx, t); err != nil {
		return nil, errInternal
	}
	return nil, errInternal
}

func (h tagHandler) ListTags(ctx context.Context, req *connect.Request[simoompb.ListTagsRequest]) (*connect.Response[simoompb.TagsResponse], error) {
	ts, err := h.repo.ListTagsByUserID(ctx, userID, 100, 0)
	if err != nil {
		return nil, errInternal
	}
	return connect.NewResponse(&simoompb.TagsResponse{
		Tags: newTagsResponse(ts),
	}), nil
}

func (h tagHandler) UpdateTag(ctx context.Context, req *connect.Request[simoompb.UpdateTagRequest]) (*connect.Response[simoompb.TagResponse], error) {
	t, err := h.repo.GetTagByID(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errTagNotFound
		}
		return nil, errInternal
	}
	if t.UserID != userID {
		return nil, errInternal
	}

	if req.Msg.Name == nil {
		return nil, errInvalidArgument
	}
	if req.Msg.Name != nil {
		t.Name = *req.Msg.Name
	}
	if err := h.repo.UpdateTag(ctx, t); err != nil {
		return nil, errInternal
	}

	return connect.NewResponse(newTagResponse(t)), nil
}

func (h tagHandler) DeleteTag(ctx context.Context, req *connect.Request[simoompb.DeleteTagRequest]) (*connect.Response[emptypb.Empty], error) {
	t, err := h.repo.GetTagByID(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errTagNotFound
		}
		return nil, errInternal
	}
	if t.UserID != userID {
		return nil, errInternal
	}

	if err := h.repo.DeleteTag(ctx, req.Msg.Id); err != nil {
		return nil, errInternal
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

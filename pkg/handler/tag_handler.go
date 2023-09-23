package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/gen/simoompb/v1"
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
	return nil, errInternal
}

func (h tagHandler) ListTags(ctx context.Context, req *connect.Request[simoompb.ListTagsRequest]) (*connect.Response[simoompb.TagsResponse], error) {
	return nil, errInternal
}

func (h tagHandler) UpdateTag(ctx context.Context, req *connect.Request[simoompb.UpdateTagRequest]) (*connect.Response[simoompb.TagResponse], error) {
	return nil, errInternal
}

func (h tagHandler) DeleteTag(ctx context.Context, req *connect.Request[simoompb.DeleteTagRequest]) (*connect.Response[emptypb.Empty], error) {
	return connect.NewResponse(&emptypb.Empty{}), nil
}

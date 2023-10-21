package handler

import (
	"context"
	"strings"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type projectHandler struct {
	uc usecase.ProjectUsecase
}

func newProjectResponse(p model.Project) *simoompb.ProjectResponse {
	return &simoompb.ProjectResponse{
		Id:         p.ID,
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
		CreatedAt:  timestamppb.New(p.CreatedAt),
		UpdatedAt:  timestamppb.New(p.UpdatedAt),
	}
}

func newProjectsResponse(ps []model.Project) []*simoompb.ProjectResponse {
	projects := make([]*simoompb.ProjectResponse, 0, len(ps))
	for _, p := range ps {
		projects = append(projects, newProjectResponse(p))
	}
	return projects
}

func (h projectHandler) CreateProject(ctx context.Context, req *connect.Request[simoompb.CreateProjectRequest]) (*connect.Response[simoompb.ProjectResponse], error) {
	if req.Msg.Name == "" {
		return nil, newErrInvalidArgument("name cannot be an empty string")
	}
	if !strings.HasPrefix(req.Msg.Color, "#") || len(req.Msg.Color) != 7 {
		return nil, newErrInvalidArgument("color is specified in the format #000000")
	}

	out, err := h.uc.CreateProject(ctx, usecase.CreateProjectInput{
		Name:  req.Msg.Name,
		Color: req.Msg.Color,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newProjectResponse(out.Project)), nil
}

func (h projectHandler) ListProjects(ctx context.Context, req *connect.Request[simoompb.ListProjectsRequest]) (*connect.Response[simoompb.ProjectsResponse], error) {
	if req.Msg.Limit < 1 {
		return nil, newErrInvalidArgument("limit is greater than or equal to 1")
	}

	out, err := h.uc.ListProjects(ctx, usecase.ListProjectsInput{
		Limit:  uint(req.Msg.Limit),
		Offset: uint(req.Msg.Offset),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&simoompb.ProjectsResponse{
		Projects: newProjectsResponse(out.Projects),
		HasNext:  out.HasNext,
	}), nil
}

func (h projectHandler) UpdateProject(ctx context.Context, req *connect.Request[simoompb.UpdateProjectRequest]) (*connect.Response[simoompb.ProjectResponse], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}
	if req.Msg.Name == nil && req.Msg.Color == nil && req.Msg.IsArchived == nil {
		return nil, newErrInvalidArgument("must contain some argument other than id")
	}
	if req.Msg.Name != nil && *req.Msg.Name == "" {
		return nil, newErrInvalidArgument("name cannot be an empty string")
	}
	if req.Msg.Color != nil && (!strings.HasPrefix(*req.Msg.Color, "#") || len(*req.Msg.Color) != 7) {
		return nil, newErrInvalidArgument("color is specified in the format #000000")
	}

	out, err := h.uc.UpdateProject(ctx, usecase.UpdateProjectInput{
		ID:         req.Msg.Id,
		Name:       req.Msg.Name,
		Color:      req.Msg.Color,
		IsArchived: req.Msg.IsArchived,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newProjectResponse(out.Project)), nil
}

func (h projectHandler) DeleteProject(ctx context.Context, req *connect.Request[simoompb.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}

	if err := h.uc.DeleteProject(ctx, usecase.DeleteProjectInput{
		ID: req.Msg.Id,
	}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

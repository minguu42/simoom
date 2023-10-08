package handler

import (
	"context"
	"strings"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
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
		return nil, errors.New("name is required")
	}
	if !strings.HasPrefix(req.Msg.Color, "#") || len(req.Msg.Color) != 7 {
		return nil, errors.New("color should be specified in the format #000000")
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

func (h projectHandler) ListProjects(ctx context.Context, _ *connect.Request[simoompb.ListProjectsRequest]) (*connect.Response[simoompb.ProjectsResponse], error) {
	out, err := h.uc.ListProjects(ctx, usecase.ListProjectsInput{
		Limit:  10,
		Offset: 0,
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
		return nil, errors.New("id is required")
	}
	if req.Msg.Name == nil && req.Msg.Color == nil && req.Msg.IsArchived == nil {
		return nil, errInvalidArgument
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
		return nil, errors.New("id is required")
	}

	if err := h.uc.DeleteProject(ctx, usecase.DeleteProjectInput{
		ID: req.Msg.Id,
	}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

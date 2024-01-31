package handler

import (
	"context"
	"strings"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/minguu42/simoom/pkg/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newProject(p model.Project) *simoompb.Project {
	return &simoompb.Project{
		Id:         p.ID,
		Name:       p.Name,
		Color:      p.Color,
		IsArchived: p.IsArchived,
	}
}

func newProjects(ps []model.Project) []*simoompb.Project {
	projects := make([]*simoompb.Project, 0, len(ps))
	for _, p := range ps {
		projects = append(projects, newProject(p))
	}
	return projects
}

func (h handler) CreateProject(ctx context.Context, req *connect.Request[simoompb.CreateProjectRequest]) (*connect.Response[simoompb.Project], error) {
	if len(req.Msg.Name) < 1 || 20 < len(req.Msg.Name) {
		return nil, newErrInvalidArgument("name must be at least 1 and no more than 20 characters")
	}
	if len(req.Msg.Color) != 7 || !strings.HasPrefix(req.Msg.Color, "#") {
		return nil, newErrInvalidArgument("color is specified in the format #000000")
	}

	out, err := h.project.CreateProject(ctx, usecase.CreateProjectInput{
		Name:  req.Msg.Name,
		Color: req.Msg.Color,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newProject(out.Project)), nil
}

func (h handler) ListProjects(ctx context.Context, req *connect.Request[simoompb.ListProjectsRequest]) (*connect.Response[simoompb.Projects], error) {
	if req.Msg.Limit < 1 {
		return nil, newErrInvalidArgument("limit is greater than or equal to 1")
	}

	out, err := h.project.ListProjects(ctx, usecase.ListProjectsInput{
		Limit:  uint(req.Msg.Limit),
		Offset: uint(req.Msg.Offset),
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&simoompb.Projects{
		Projects: newProjects(out.Projects),
		HasNext:  out.HasNext,
	}), nil
}

func (h handler) UpdateProject(ctx context.Context, req *connect.Request[simoompb.UpdateProjectRequest]) (*connect.Response[simoompb.Project], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}
	if req.Msg.Name == nil && req.Msg.Color == nil && req.Msg.IsArchived == nil {
		return nil, newErrInvalidArgument("must contain some argument other than id")
	}
	if req.Msg.Name != nil && (len(*req.Msg.Name) < 1 || 20 < len(*req.Msg.Name)) {
		return nil, newErrInvalidArgument("name must be at least 1 and no more than 20 characters")
	}
	if req.Msg.Color != nil && (len(*req.Msg.Color) != 7 || !strings.HasPrefix(*req.Msg.Color, "#")) {
		return nil, newErrInvalidArgument("color is specified in the format #000000")
	}

	out, err := h.project.UpdateProject(ctx, usecase.UpdateProjectInput{
		ID:         req.Msg.Id,
		Name:       req.Msg.Name,
		Color:      req.Msg.Color,
		IsArchived: req.Msg.IsArchived,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(newProject(out.Project)), nil
}

func (h handler) DeleteProject(ctx context.Context, req *connect.Request[simoompb.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error) {
	if len(req.Msg.Id) != 26 {
		return nil, newErrInvalidArgument("id is a 26-character string")
	}

	if err := h.project.DeleteProject(ctx, usecase.DeleteProjectInput{
		ID: req.Msg.Id,
	}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

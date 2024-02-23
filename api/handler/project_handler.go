package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
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
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, ErrInvalidRequest
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
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, ErrInvalidRequest
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
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, ErrInvalidRequest
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
	if err := h.validator.Validate(req.Msg); err != nil {
		return nil, ErrInvalidRequest
	}

	if err := h.project.DeleteProject(ctx, usecase.DeleteProjectInput{
		ID: req.Msg.Id,
	}); err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

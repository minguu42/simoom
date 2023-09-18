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

type projectHandler struct {
	repo repository.Repository
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
	now := time.Now()
	p := model.Project{
		ID:         idgen.Generate(),
		UserID:     userID,
		Name:       req.Msg.Name,
		Color:      req.Msg.Color,
		IsArchived: false,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	if err := h.repo.CreateProject(ctx, p); err != nil {
		return nil, errInternal
	}

	return connect.NewResponse(newProjectResponse(p)), nil
}

func (h projectHandler) ListProjects(ctx context.Context, _ *connect.Request[simoompb.ListProjectsRequest]) (*connect.Response[simoompb.ProjectsResponse], error) {
	ps, err := h.repo.ListProjectsByUserID(ctx, userID, 10, 0)
	if err != nil {
		return nil, errInternal
	}

	return connect.NewResponse(&simoompb.ProjectsResponse{
		Projects: newProjectsResponse(ps),
		HasNext:  false,
	}), nil
}

func (h projectHandler) UpdateProject(ctx context.Context, req *connect.Request[simoompb.UpdateProjectRequest]) (*connect.Response[simoompb.ProjectResponse], error) {
	p, err := h.repo.GetProjectByID(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errProjectNotFound
		}
		return nil, errInternal
	}

	if req.Msg.Name == nil && req.Msg.Color == nil && req.Msg.IsArchived == nil {
		return nil, errInvalidArgument
	}
	if req.Msg.Name != nil {
		p.Name = *req.Msg.Name
	}
	if req.Msg.Color != nil {
		p.Color = *req.Msg.Color
	}
	if req.Msg.IsArchived != nil {
		p.IsArchived = *req.Msg.IsArchived
	}
	if err := h.repo.UpdateProject(ctx, p); err != nil {
		return nil, errInternal
	}

	return connect.NewResponse(newProjectResponse(p)), nil
}

func (h projectHandler) DeleteProject(ctx context.Context, req *connect.Request[simoompb.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error) {
	p, err := h.repo.GetProjectByID(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errProjectNotFound
		}
		return nil, errInternal
	}
	if p.UserID != userID {
		return nil, errProjectNotFound
	}

	if err := h.repo.DeleteProject(ctx, req.Msg.Id); err != nil {
		return nil, errInternal
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

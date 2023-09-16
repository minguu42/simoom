package handler

import (
	"context"
	"time"

	"connectrpc.com/connect"
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
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(newProjectResponse(p)), nil
}

func (h projectHandler) ListProjects(ctx context.Context, _ *connect.Request[simoompb.ListProjectsRequest]) (*connect.Response[simoompb.ProjectsResponse], error) {
	ps, err := h.repo.GetProjectsByUserID(ctx, userID, 10, 0)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&simoompb.ProjectsResponse{
		Projects: newProjectsResponse(ps),
		HasNext:  false,
	}), nil
}

func (h projectHandler) UpdateProject(_ context.Context, _ *connect.Request[simoompb.UpdateProjectRequest]) (*connect.Response[simoompb.ProjectResponse], error) {
	return connect.NewResponse(&simoompb.ProjectResponse{
		Id:         "",
		Name:       "",
		Color:      "",
		IsArchived: false,
		CreatedAt:  nil,
		UpdatedAt:  nil,
	}), nil
}

func (h projectHandler) DeleteProject(_ context.Context, _ *connect.Request[simoompb.DeleteProjectRequest]) (*connect.Response[emptypb.Empty], error) {
	return connect.NewResponse(&emptypb.Empty{}), nil
}

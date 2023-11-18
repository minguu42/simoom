package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/backend/pkg/pointers"
	"github.com/minguu42/simoom/lib/simoompb/v1"
)

func TestProjectHandler_CreateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.CreateProjectRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "nameに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateProjectRequest{
					Name:  "",
					Color: "#000000",
				}),
			},
			hasError: true,
		},
		{
			name: "colorは#000000の形式で指定する",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateProjectRequest{
					Name:  "some name",
					Color: "red",
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := th.CreateProject(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("th.CreateProject should return an error")
			}
		})
	}
}

func TestProjectHandler_ListProjects(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.ListProjectsRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "limitは1以上である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.ListProjectsRequest{
					Limit: 0,
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := th.ListProjects(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("th.ListProjects should return an error")
			}
		})
	}
}

func TestProjectHandler_UpdateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.UpdateProjectRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "idは26文字の文字列である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateProjectRequest{
					Id: "some-id",
				}),
			},
			hasError: true,
		},
		{
			name: "いずれかの引数は必要である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateProjectRequest{
					Id:         "01DXF6DT000000000000000000",
					Name:       nil,
					Color:      nil,
					IsArchived: nil,
				}),
			},
			hasError: true,
		},
		{
			name: "nameに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateProjectRequest{
					Id:   "01DXF6DT000000000000000000",
					Name: pointers.Ref(""),
				}),
			},
			hasError: true,
		},
		{
			name: "colorは#000000の形式で指定する",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateProjectRequest{
					Id:    "01DXF6DT000000000000000000",
					Color: pointers.Ref("red"),
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := th.UpdateProject(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("th.UpdateProject should return an error")
			}
		})
	}
}

func TestProjectHandler_DeleteProject(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.DeleteProjectRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "idは26文字の文字列である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.DeleteProjectRequest{
					Id: "some-id",
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := th.DeleteProject(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("th.DeleteProject should return an error")
			}
		})
	}
}

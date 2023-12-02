package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/pointers"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/stretchr/testify/assert"
)

func TestProjectHandler_CreateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.CreateProjectRequest]
	}
	tests := []struct {
		name string
		args args
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.CreateProject(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

func TestProjectHandler_ListProjects(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.ListProjectsRequest]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "limitは1以上である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.ListProjectsRequest{
					Limit: 0,
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.ListProjects(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

func TestProjectHandler_UpdateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.UpdateProjectRequest]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "idは26文字の文字列である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateProjectRequest{
					Id: "some-id",
				}),
			},
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.UpdateProject(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

func TestProjectHandler_DeleteProject(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.DeleteProjectRequest]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "idは26文字の文字列である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.DeleteProjectRequest{
					Id: "some-id",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.DeleteProject(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

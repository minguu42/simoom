package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/gen/simoompb/v1"
)

var ph = projectHandler{}

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
			name: "Nameは必須である",
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
			name: "Colorは#000000の形式で指定する",
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
			if !tt.hasError {
				return
			}
			if _, err := ph.CreateProject(tt.args.ctx, tt.args.req); err == nil {
				t.Errorf("ph.CreateProject should return an error")
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
			name: "IDは26文字である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateProjectRequest{
					Id: "some-id",
				}),
			},
			hasError: true,
		},
		{
			name: "いずれかのパラメータは必要である",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.hasError {
				return
			}
			if _, err := ph.UpdateProject(tt.args.ctx, tt.args.req); err == nil {
				t.Errorf("ph.UpdateProject should return an error")
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
			name: "IDは26文字である",
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
			if !tt.hasError {
				return
			}
			if _, err := ph.DeleteProject(tt.args.ctx, tt.args.req); err == nil {
				t.Errorf("ph.UpdateProject should return an error")
			}
		})
	}
}

package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/lib/go/simoompb/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestHandler_CreateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.CreateProjectRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[simoompb.Project]
		wantErr apperr.Error
	}{
		{
			name: "不正なリクエストはバリデーションではじく",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateProjectRequest{}),
			},
			want:    nil,
			wantErr: apperr.ErrInvalidRequest(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testHandler.CreateProject(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}

func TestHandler_ListProjects(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.ListProjectsRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[simoompb.Projects]
		wantErr apperr.Error
	}{
		{
			name: "不正なリクエストはバリデーションではじく",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.ListProjectsRequest{}),
			},
			want:    nil,
			wantErr: apperr.ErrInvalidRequest(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testHandler.ListProjects(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}

func TestHandler_UpdateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.UpdateProjectRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[simoompb.Project]
		wantErr apperr.Error
	}{
		{
			name: "不正なリクエストはバリデーションではじく",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateProjectRequest{}),
			},
			want:    nil,
			wantErr: apperr.ErrInvalidRequest(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testHandler.UpdateProject(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}

func TestHandler_DeleteProject(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.DeleteProjectRequest]
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[emptypb.Empty]
		wantErr apperr.Error
	}{
		{
			name: "不正なリクエストはバリデーションではじく",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.DeleteProjectRequest{}),
			},
			want:    nil,
			wantErr: apperr.ErrInvalidRequest(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testHandler.DeleteProject(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, resp)

			var appErr apperr.Error
			if tt.wantErr.Error() != "" && assert.ErrorAs(t, err, &appErr) {
				assert.Equal(t, tt.wantErr.ConnectError().Code(), appErr.ConnectError().Code())
			}
		})
	}
}

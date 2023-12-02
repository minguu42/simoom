package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/pointers"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/stretchr/testify/assert"
)

func TestTagHandler_CreateTag(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.CreateTagRequest]
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
				req: connect.NewRequest(&simoompb.CreateTagRequest{
					Name: "",
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.CreateTag(tt.args.ctx, tt.args.req)
			if tt.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestTagHandler_ListTags(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.ListTagsRequest]
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
				req: connect.NewRequest(&simoompb.ListTagsRequest{
					Limit: 0,
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.ListTags(tt.args.ctx, tt.args.req)
			if tt.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestTagHandler_UpdateTag(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.UpdateTagRequest]
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
				req: connect.NewRequest(&simoompb.UpdateTagRequest{
					Id: "some-id",
				}),
			},
			hasError: true,
		},
		{
			name: "いずれかの引数は必要である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateTagRequest{
					Id:   "01DXF6DT000000000000000000",
					Name: nil,
				}),
			},
			hasError: true,
		},
		{
			name: "nameに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateTagRequest{
					Id:   "01DXF6DT000000000000000000",
					Name: pointers.Ref(""),
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.UpdateTag(tt.args.ctx, tt.args.req)
			if tt.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestTagHandler_DeleteTag(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.DeleteTagRequest]
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
				req: connect.NewRequest(&simoompb.DeleteTagRequest{
					Id: "some-id",
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.DeleteTag(tt.args.ctx, tt.args.req)
			if tt.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

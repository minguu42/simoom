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
		name string
		args args
	}{
		{
			name: "nameに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateTagRequest{
					Name: "",
				}),
			},
		},
		{
			name: "nameに21文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateTagRequest{
					Name: "very-long-long-name01",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.CreateTag(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

func TestTagHandler_ListTags(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.ListTagsRequest]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "limitに0は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.ListTagsRequest{
					Limit:  0,
					Offset: 0,
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.ListTags(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

func TestTagHandler_UpdateTag(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.UpdateTagRequest]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "idに25文字以下の文字列を指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateTagRequest{
					Id:   "xxxx-xxxx-xxxx-xxxx-id345",
					Name: pointers.Ref("some-tag"),
				}),
			},
		},
		{
			name: "idに27文字以上の文字列を指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateTagRequest{
					Id:   "xxxx-xxxx-xxxx-xxxx-xxxx-id",
					Name: pointers.Ref("some-tag"),
				}),
			},
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
		},
		{
			name: "nameに21文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateTagRequest{
					Id:   "01DXF6DT000000000000000000",
					Name: pointers.Ref("very-long-long-name01"),
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.UpdateTag(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

func TestTagHandler_DeleteTag(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.DeleteTagRequest]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "idに25文字以下の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.DeleteTagRequest{
					Id: "xxxx-xxxx-xxxx-xxxx-id345",
				}),
			},
		},
		{
			name: "idに27文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.DeleteTagRequest{
					Id: "xxxx-xxxx-xxxx-xxxx-xxxx-id",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.DeleteTag(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

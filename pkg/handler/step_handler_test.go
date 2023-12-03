package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/pointers"
	"github.com/minguu42/simoom/pkg/simoompb/v1"
	"github.com/stretchr/testify/assert"
)

func TestStepHandler_CreateStep(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.CreateStepRequest]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "task_idに25文字以下の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateStepRequest{
					TaskId: "xxxx-xxxx-xxxx-xxxx-id345",
					Title:  "some-step",
				}),
			},
		},
		{
			name: "task_idに27文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateStepRequest{
					TaskId: "xxxx-xxxx-xxxx-xxxx-xxxx-id",
					Title:  "some-step",
				}),
			},
		},
		{
			name: "titleに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateStepRequest{
					TaskId: "01DXF6DT000000000000000000",
					Title:  "",
				}),
			},
		},
		{
			name: "titleに81文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateStepRequest{
					TaskId: "01DXF6DT000000000000000000",
					Title:  "very-long-long-long-long-long-long-long-long-long-long-long-long-long-long-step01",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.CreateStep(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

func TestStepHandler_UpdateStep(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.UpdateStepRequest]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "idに25文字以下の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateStepRequest{
					Id:    "xxxx-xxxx-xxxx-xxxx-id345",
					Title: pointers.Ref("some-step"),
				}),
			},
		},
		{
			name: "idに27文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateStepRequest{
					Id:    "xxxx-xxxx-xxxx-xxxx-xxxx-id",
					Title: pointers.Ref("some-step"),
				}),
			},
		},
		{
			name: "いずれかの引数は必要である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateStepRequest{
					Id:          "01DXF6DT000000000000000000",
					Title:       nil,
					CompletedAt: nil,
				}),
			},
		},
		{
			name: "titleに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateStepRequest{
					Id:    "01DXF6DT000000000000000000",
					Title: pointers.Ref(""),
				}),
			},
		},
		{
			name: "titleに81文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateStepRequest{
					Id:    "01DXF6DT000000000000000000",
					Title: pointers.Ref("very-long-long-long-long-long-long-long-long-long-long-long-long-long-long-step01"),
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.UpdateStep(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

func TestStepHandler_DeleteStep(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.DeleteStepRequest]
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "idに25文字以下の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.DeleteStepRequest{
					Id: "xxxx-xxxx-xxxx-xxxx-id345",
				}),
			},
		},
		{
			name: "idに27文字以上の文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.DeleteStepRequest{
					Id: "xxxx-xxxx-xxxx-xxxx-xxxx-id",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := th.DeleteStep(tt.args.ctx, tt.args.req)
			assert.Error(t, err)
		})
	}
}

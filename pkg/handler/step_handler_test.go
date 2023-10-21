package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/gen/simoompb/v1"
	"github.com/minguu42/simoom/pkg/pointers"
)

var tsh = stepHandler{}

func TestStepHandler_CreateStep(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.CreateStepRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "task_idは26文字の文字列である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateStepRequest{
					TaskId: "some-id",
				}),
			},
			hasError: true,
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
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tsh.CreateStep(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("tsh.CreateStep should return an error")
			}
		})
	}
}

func TestStepHandler_UpdateStep(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.UpdateStepRequest]
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
				req: connect.NewRequest(&simoompb.UpdateStepRequest{
					Id: "some-id",
				}),
			},
			hasError: true,
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
			hasError: true,
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
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tsh.UpdateStep(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("tsh.UpdateStep should return an error")
			}
		})
	}
}

func TestStepHandler_DeleteStep(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.DeleteStepRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "idは26文字である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.DeleteStepRequest{
					Id: "some-id",
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tsh.DeleteStep(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("tsh.DeleteStep should return an error")
			}
		})
	}
}

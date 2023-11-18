package handler

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/backend/pkg/pointers"
	"github.com/minguu42/simoom/lib/simoompb/v1"
)

func TestTaskHandler_CreateTask(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.CreateTaskRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "project_idは26文字の文字列である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateTaskRequest{
					ProjectId: "some-id",
				}),
			},
			hasError: true,
		},
		{

			name: "titleに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateTaskRequest{
					ProjectId: "01DXF6DT000000000000000000",
					Title:     "",
				}),
			},
			hasError: true,
		},
		{
			name: "priorityは0から3の整数で指定する",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.CreateTaskRequest{
					ProjectId: "01DXF6DT000000000000000000",
					Title:     "some-task",
					Priority:  4,
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := th.CreateTask(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("th.CreateTask should return an error")
			}
		})
	}
}

func TestTaskHandler_ListTasksByProjectID(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.ListTasksByProjectIDRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "project_idは26文字の文字列である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.ListTasksByProjectIDRequest{
					ProjectId: "some-id",
				}),
			},
			hasError: true,
		},
		{
			name: "limitは1以上である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.ListTasksByProjectIDRequest{
					Limit: 0,
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := th.ListTasksByProjectID(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("th.ListTasksByProjectID should return an error")
			}
		})
	}
}

func TestTaskHandler_ListTasksByTagID(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.ListTasksByTagIDRequest]
	}
	tests := []struct {
		name     string
		args     args
		hasError bool
	}{
		{
			name: "tag_idは26文字の文字列である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.ListTasksByTagIDRequest{
					TagId: "some-id",
				}),
			},
			hasError: true,
		},
		{
			name: "limitは1以上である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.ListTasksByTagIDRequest{
					Limit: 0,
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := th.ListTasksByTagID(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("th.ListTasksByTagID should return an error")
			}
		})
	}
}

func TestTaskHandler_UpdateTask(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.UpdateTaskRequest]
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
				req: connect.NewRequest(&simoompb.UpdateTaskRequest{
					Id: "some-id",
				}),
			},
			hasError: true,
		},
		{
			name: "いずれかの引数は必要である",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateTaskRequest{
					Id:          "01DXF6DT000000000000000000",
					Title:       nil,
					Content:     nil,
					Priority:    nil,
					DueOn:       nil,
					CompletedAt: nil,
				}),
			},
			hasError: true,
		},
		{
			name: "titleに空文字列は指定できない",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateTaskRequest{
					Id:    "01DXF6DT000000000000000000",
					Title: pointers.Ref(""),
				}),
			},
			hasError: true,
		},
		{
			name: "priorityは0から3の整数で指定する",
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&simoompb.UpdateTaskRequest{
					Id:       "01DXF6DT000000000000000000",
					Priority: pointers.Ref(uint32(4)),
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := th.UpdateTask(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("th.UpdateTask should return an error")
			}
		})
	}
}

func TestTaskHandler_DeleteTask(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[simoompb.DeleteTaskRequest]
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
				req: connect.NewRequest(&simoompb.DeleteTaskRequest{
					Id: "some-id",
				}),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := th.DeleteTask(tt.args.ctx, tt.args.req); tt.hasError != (err != nil) {
				t.Errorf("th.DeleteTask should return an error")
			}
		})
	}
}

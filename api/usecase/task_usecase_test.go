package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/pointers"
)

var createTaskOption = cmpopts.IgnoreFields(usecase.TaskOutput{}, "Task.ID")

func TestTaskUsecase_CreateTask(t *testing.T) {
	type args struct {
		ctx context.Context
		in  usecase.CreateTaskInput
	}
	tests := []struct {
		name string
		args args
		want usecase.TaskOutput
	}{
		{
			name: "新タスクを作成する",
			args: args{
				ctx: tctx,
				in: usecase.CreateTaskInput{
					ProjectID: "project_01",
					Name:      "新タスク",
					Priority:  3,
				},
			},
			want: usecase.TaskOutput{Task: model.Task{
				UserID:    "user_01",
				ProjectID: "project_01",
				Name:      "新タスク",
				Content:   "",
				Priority:  3,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := task.CreateTask(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got, createTaskOption); diff != "" {
				t.Errorf("task.CreateTask mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTaskUsecase_UpdateTask(t *testing.T) {
	type args struct {
		ctx context.Context
		in  usecase.UpdateTaskInput
	}
	tests := []struct {
		name string
		args args
		want usecase.TaskOutput
	}{
		{
			name: "改タスク1に更新する",
			args: args{
				ctx: tctx,
				in: usecase.UpdateTaskInput{
					ID:          "task_01",
					Name:        pointers.Ref("改タスク1"),
					Content:     pointers.Ref("テストコンテンツ1"),
					Priority:    pointers.Ref(uint(3)),
					DueOn:       pointers.Ref(time.Date(2020, 1, 10, 0, 0, 1, 0, time.UTC)),
					CompletedAt: pointers.Ref(time.Date(2020, 1, 2, 0, 0, 1, 0, time.UTC)),
				},
			},
			want: usecase.TaskOutput{Task: model.Task{
				ID: "task_01",
				Steps: []model.Step{
					{
						ID:          "step_01",
						UserID:      "user_01",
						TaskID:      "task_01",
						Name:        "ステップ1",
						CompletedAt: nil,
					},
					{
						ID:          "step_02",
						UserID:      "user_01",
						TaskID:      "task_01",
						Name:        "ステップ2",
						CompletedAt: nil,
					},
				},
				Tags: []model.Tag{
					{
						ID:     "tag_01",
						UserID: "user_01",
						Name:   "タグ1",
					},
					{
						ID:     "tag_02",
						UserID: "user_01",
						Name:   "タグ2",
					},
				},
				UserID:      "user_01",
				ProjectID:   "project_01",
				Name:        "改タスク1",
				Content:     "テストコンテンツ1",
				Priority:    3,
				DueOn:       pointers.Ref(time.Date(2020, 1, 10, 0, 0, 1, 0, time.UTC)),
				CompletedAt: pointers.Ref(time.Date(2020, 1, 2, 0, 0, 1, 0, time.UTC)),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := task.UpdateTask(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("task.UpdateTask mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTaskUsecase_DeleteTask(t *testing.T) {
	type args struct {
		ctx context.Context
		in  usecase.DeleteTaskInput
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "タスク1を削除する",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteTaskInput{ID: "task_01"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			if err := task.DeleteTask(tt.args.ctx, tt.args.in); err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

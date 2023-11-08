package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/infra/mysql"
	"github.com/minguu42/simoom/pkg/pointers"
	"github.com/minguu42/simoom/pkg/usecase"
)

var (
	createTaskOption = cmpopts.IgnoreFields(usecase.TaskOutput{},
		"Task.ID",
		"Task.CreatedAt",
		"Task.UpdatedAt",
	)
	updateTaskOption = cmpopts.IgnoreFields(usecase.TaskOutput{}, "Task.UpdatedAt")
)

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
				ctx: ctx,
				in: usecase.CreateTaskInput{
					ProjectID: "project_01",
					Title:     "新タスク",
					Priority:  3,
				},
			},
			want: usecase.TaskOutput{Task: model.Task{
				UserID:    "user_01",
				ProjectID: "project_01",
				Title:     "新タスク",
				Content:   "",
				Priority:  3,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetTask(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
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
				ctx: ctx,
				in: usecase.UpdateTaskInput{
					ID:          "task_01",
					Title:       pointers.Ref("改タスク1"),
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
						Title:       "ステップ1",
						CompletedAt: nil,
						CreatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
						UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
					},
					{
						ID:          "step_02",
						UserID:      "user_01",
						TaskID:      "task_01",
						Title:       "ステップ2",
						CompletedAt: nil,
						CreatedAt:   time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
						UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
					},
				},
				Tags: []model.Tag{
					{
						ID:        "tag_01",
						UserID:    "user_01",
						Name:      "タグ1",
						CreatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
						UpdatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
					},
					{
						ID:        "tag_02",
						UserID:    "user_01",
						Name:      "タグ2",
						CreatedAt: time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
						UpdatedAt: time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
					},
				},
				UserID:      "user_01",
				ProjectID:   "project_01",
				Title:       "改タスク1",
				Content:     "テストコンテンツ1",
				Priority:    3,
				DueOn:       pointers.Ref(time.Date(2020, 1, 10, 0, 0, 1, 0, time.UTC)),
				CompletedAt: pointers.Ref(time.Date(2020, 1, 2, 0, 0, 1, 0, time.UTC)),
				CreatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetTask(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})

			got, err := task.UpdateTask(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got, updateTaskOption); diff != "" {
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
				ctx: ctx,
				in:  usecase.DeleteTaskInput{ID: "task_01"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetTask(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})

			if err := task.DeleteTask(tt.args.ctx, tt.args.in); err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

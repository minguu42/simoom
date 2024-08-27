package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/pointers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTaskUsecase_CreateTask(t *testing.T) {
	task := usecase.NewTask(tc, &domain.IDGeneratorMock{GenerateFunc: func() string {
		return "task_99"
	}})
	type args struct {
		ctx context.Context
		in  usecase.CreateTaskInput
	}
	tests := []struct {
		name    string
		args    args
		want    usecase.TaskOutput
		wantErr apperr.Error
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
				ID:        "task_99",
				UserID:    "user_01",
				ProjectID: "project_01",
				Name:      "新タスク",
				Priority:  3,
			}},
		},
		{
			name: "指定したプロジェクトが存在しない場合はエラーを返す",
			args: args{
				ctx: tctx,
				in: usecase.CreateTaskInput{
					ProjectID: "project_99",
					Name:      "新タスク",
				},
			},
			wantErr: apperr.ErrProjectNotFound(nil),
		},
		{
			name: "指定したプロジェクトを所有していない場合はエラーを返す",
			args: args{
				ctx: tctxUser2,
				in: usecase.CreateTaskInput{
					ProjectID: "project_01",
					Name:      "新タスク",
				},
			},
			wantErr: apperr.ErrProjectNotFound(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := task.CreateTask(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.want, got)
			if tt.wantErr.IsZero() {
				assert.NoError(t, err)
			} else {
				var appErr apperr.Error
				require.ErrorAs(t, err, &appErr)
				assert.Equal(t, tt.wantErr.ID(), appErr.ID())
			}
		})
	}
}

func TestTaskUsecase_ListTasks(t *testing.T) {
	t.Parallel()

	task := usecase.NewTask(tc, &domain.IDGeneratorMock{})
	type args struct {
		ctx context.Context
		in  usecase.ListTasksInput
	}
	tests := []struct {
		name string
		args args
		want usecase.TasksOutput
	}{
		{
			name: "タスク一覧を取得する",
			args: args{
				ctx: tctx,
				in: usecase.ListTasksInput{
					Limit:  10,
					Offset: 0,
				},
			},
			want: usecase.TasksOutput{
				Tasks: []model.Task{
					{
						ID: "task_01",
						Steps: []model.Step{
							{ID: "step_01", UserID: "user_01", TaskID: "task_01", Name: "ステップ1"},
							{ID: "step_02", UserID: "user_01", TaskID: "task_01", Name: "ステップ2"},
						},
						Tags: []model.Tag{
							{ID: "tag_01", UserID: "user_01", Name: "タグ1"},
							{ID: "tag_02", UserID: "user_01", Name: "タグ2"},
						},
						UserID:    "user_01",
						ProjectID: "project_01",
						Name:      "タスク1",
						Content:   "コンテンツ1",
						Priority:  3,
						DueOn:     pointers.Ref(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)),
					},
					{
						ID:        "task_02",
						Steps:     []model.Step{},
						Tags:      []model.Tag{},
						UserID:    "user_01",
						ProjectID: "project_01",
						Name:      "タスク2",
					},
				},
				HasNext: false,
			},
		},
		{
			name: "limitで制限をかける",
			args: args{
				ctx: tctx,
				in: usecase.ListTasksInput{
					Limit:  1,
					Offset: 0,
				},
			},
			want: usecase.TasksOutput{
				Tasks: []model.Task{
					{
						ID: "task_01",
						Steps: []model.Step{
							{ID: "step_01", UserID: "user_01", TaskID: "task_01", Name: "ステップ1"},
							{ID: "step_02", UserID: "user_01", TaskID: "task_01", Name: "ステップ2"},
						},
						Tags: []model.Tag{
							{ID: "tag_01", UserID: "user_01", Name: "タグ1"},
							{ID: "tag_02", UserID: "user_01", Name: "タグ2"},
						},
						UserID:    "user_01",
						ProjectID: "project_01",
						Name:      "タスク1",
						Content:   "コンテンツ1",
						Priority:  3,
						DueOn:     pointers.Ref(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)),
					},
				},
				HasNext: true,
			},
		},
		{
			name: "limitとoffsetでページングを行う",
			args: args{
				ctx: tctx,
				in: usecase.ListTasksInput{
					Limit:  1,
					Offset: 1,
				},
			},
			want: usecase.TasksOutput{
				Tasks: []model.Task{
					{
						ID:        "task_02",
						Steps:     []model.Step{},
						Tags:      []model.Tag{},
						UserID:    "user_01",
						ProjectID: "project_01",
						Name:      "タスク2",
					},
				},
				HasNext: false,
			},
		},
		{
			name: "projectIDで絞り込む",
			args: args{
				ctx: tctx,
				in: usecase.ListTasksInput{
					Limit:     1,
					ProjectID: pointers.Ref(model.ProjectID("project_01")),
				},
			},
			want: usecase.TasksOutput{
				Tasks: []model.Task{
					{
						ID: "task_01",
						Steps: []model.Step{
							{ID: "step_01", UserID: "user_01", TaskID: "task_01", Name: "ステップ1"},
							{ID: "step_02", UserID: "user_01", TaskID: "task_01", Name: "ステップ2"},
						},
						Tags: []model.Tag{
							{ID: "tag_01", UserID: "user_01", Name: "タグ1"},
							{ID: "tag_02", UserID: "user_01", Name: "タグ2"},
						},
						UserID:    "user_01",
						ProjectID: "project_01",
						Name:      "タスク1",
						Content:   "コンテンツ1",
						Priority:  3,
						DueOn:     pointers.Ref(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)),
					},
				},
				HasNext: true,
			},
		},
		{
			name: "tagIDで絞り込む",
			args: args{
				ctx: tctx,
				in: usecase.ListTasksInput{
					Limit: 1,
					TagID: pointers.Ref(model.TagID("tag_01")),
				},
			},
			want: usecase.TasksOutput{
				Tasks: []model.Task{
					{
						ID: "task_01",
						Steps: []model.Step{
							{ID: "step_01", UserID: "user_01", TaskID: "task_01", Name: "ステップ1"},
							{ID: "step_02", UserID: "user_01", TaskID: "task_01", Name: "ステップ2"},
						},
						Tags: []model.Tag{
							{ID: "tag_01", UserID: "user_01", Name: "タグ1"},
							{ID: "tag_02", UserID: "user_01", Name: "タグ2"},
						},
						UserID:    "user_01",
						ProjectID: "project_01",
						Name:      "タスク1",
						Content:   "コンテンツ1",
						Priority:  3,
						DueOn:     pointers.Ref(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)),
					},
				},
				HasNext: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := task.ListTasks(tt.args.ctx, tt.args.in)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTaskUsecase_UpdateTask(t *testing.T) {
	task := usecase.NewTask(tc, &domain.IDGeneratorMock{})
	type args struct {
		ctx context.Context
		in  usecase.UpdateTaskInput
	}
	tests := []struct {
		name    string
		args    args
		want    usecase.TaskOutput
		wantErr apperr.Error
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
					DueOn:       pointers.Ref(time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)),
					CompletedAt: pointers.Ref(time.Date(2020, 1, 2, 0, 0, 1, 0, time.UTC)),
					TagIDs:      []model.TagID{"tag_02"},
				},
			},
			want: usecase.TaskOutput{Task: model.Task{
				ID: "task_01",
				Steps: []model.Step{
					{ID: "step_01", UserID: "user_01", TaskID: "task_01", Name: "ステップ1"},
					{ID: "step_02", UserID: "user_01", TaskID: "task_01", Name: "ステップ2"},
				},
				Tags: []model.Tag{
					{ID: "tag_02", UserID: "user_01", Name: "タグ2"},
				},
				UserID:      "user_01",
				ProjectID:   "project_01",
				Name:        "改タスク1",
				Content:     "テストコンテンツ1",
				Priority:    3,
				DueOn:       pointers.Ref(time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)),
				CompletedAt: pointers.Ref(time.Date(2020, 1, 2, 0, 0, 1, 0, time.UTC)),
			}},
		},
		{
			name: "指定したタスクが存在しない場合はエラーを返す",
			args: args{
				ctx: tctx,
				in: usecase.UpdateTaskInput{
					ID:   "task_99",
					Name: pointers.Ref("改タスク99"),
				},
			},
			wantErr: apperr.ErrTaskNotFound(nil),
		},
		{
			name: "指定したタスクを所有していない場合はエラーを返す",
			args: args{
				ctx: tctxUser2,
				in: usecase.UpdateTaskInput{
					ID:   "task_01",
					Name: pointers.Ref("改タスク1"),
				},
			},
			wantErr: apperr.ErrTaskNotFound(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := task.UpdateTask(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.want, got)
			if tt.wantErr.IsZero() {
				assert.NoError(t, err)
			} else {
				var appErr apperr.Error
				require.ErrorAs(t, err, &appErr)
				assert.Equal(t, tt.wantErr.ID(), appErr.ID())
			}
		})
	}
}

func TestTaskUsecase_DeleteTask(t *testing.T) {
	task := usecase.NewTask(tc, &domain.IDGeneratorMock{})
	type args struct {
		ctx context.Context
		in  usecase.DeleteTaskInput
	}
	tests := []struct {
		name string
		args args
		want apperr.Error
	}{
		{
			name: "タスク1を削除する",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteTaskInput{ID: "task_01"},
			},
		},
		{
			name: "指定したタスクが存在しない場合はエラーを返す",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteTaskInput{ID: "task_99"},
			},
			want: apperr.ErrTaskNotFound(nil),
		},
		{
			name: "指定したタスクを所有していない場合はエラーを返す",
			args: args{
				ctx: tctxUser2,
				in:  usecase.DeleteTaskInput{ID: "task_01"},
			},
			want: apperr.ErrTaskNotFound(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			err := task.DeleteTask(tt.args.ctx, tt.args.in)
			if tt.want.IsZero() {
				assert.NoError(t, err)
			} else {
				var appErr apperr.Error
				require.ErrorAs(t, err, &appErr)
				assert.Equal(t, tt.want.ID(), appErr.ID())
			}
		})
	}
}

package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_CreateTask(t *testing.T) {
	type args struct {
		ctx context.Context
		t   model.Task
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "新タスクを作成する",
			args: args{
				ctx: context.Background(),
				t: model.Task{
					ID:          "task_99",
					Steps:       []model.Step{},
					Tags:        []model.Tag{},
					UserID:      "user_01",
					ProjectID:   "project_01",
					Title:       "新タスク",
					Content:     "",
					Priority:    1,
					DueOn:       nil,
					CompletedAt: nil,
					CreatedAt:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				ResetTask(t, tc)
			})
			err := tc.CreateTask(tt.args.ctx, tt.args.t)
			require.NoError(t, err)

			if got, err := tc.GetTaskByID(context.Background(), tt.args.t.ID); assert.NoError(t, err) {
				assert.Equal(t, tt.args.t, got)
			}
		})
	}
}

func TestClient_ListTasksByProjectID(t *testing.T) {
	type args struct {
		ctx       context.Context
		projectID string
		limit     uint
		offset    uint
	}
	dueOn := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name string
		args args
		want []model.Task
	}{
		{
			name: "プロジェクトIDからタスク一覧を取得する",
			args: args{
				ctx:       context.Background(),
				projectID: "project_01",
				limit:     1,
				offset:    0,
			},
			want: []model.Task{
				{
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
					ProjectID:   "project_01",
					UserID:      "user_01",
					Title:       "タスク1",
					Content:     "コンテンツ1",
					Priority:    3,
					DueOn:       &dueOn,
					CompletedAt: nil,
					CreatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
					UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
				},
			},
		},
		{
			name: "offsetが大きいので、空のスライスを取得する",
			args: args{
				ctx:       context.Background(),
				projectID: "project_01",
				limit:     10,
				offset:    1000,
			},
			want: []model.Task{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tc.ListTasksByProjectID(tt.args.ctx, tt.args.projectID, tt.args.limit, tt.args.offset); assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestClient_ListTasksByTagID(t *testing.T) {
	type args struct {
		ctx    context.Context
		tagID  string
		limit  uint
		offset uint
	}
	dueOn := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name string
		args args
		want []model.Task
	}{
		{
			name: "タグIDからタスク一覧を取得する",
			args: args{
				ctx:    context.Background(),
				tagID:  "tag_01",
				limit:  1,
				offset: 0,
			},
			want: []model.Task{
				{
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
					Title:       "タスク1",
					Content:     "コンテンツ1",
					Priority:    3,
					DueOn:       &dueOn,
					CompletedAt: nil,
					CreatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
					UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
				},
			},
		},
		{
			name: "offsetが大きいので、空のスライスを取得する",
			args: args{
				ctx:    context.Background(),
				tagID:  "tag_01",
				limit:  10,
				offset: 1000,
			},
			want: []model.Task{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tc.ListTasksByTagID(tt.args.ctx, tt.args.tagID, tt.args.limit, tt.args.offset); assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestClient_GetTaskByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	dueOn := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		args    args
		want    model.Task
		wantErr error
	}{
		{
			name: "タスク1を取得する",
			args: args{
				ctx: context.Background(),
				id:  "task_01",
			},
			want: model.Task{
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
				Title:       "タスク1",
				Content:     "コンテンツ1",
				Priority:    3,
				DueOn:       &dueOn,
				CompletedAt: nil,
				CreatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
				UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			},
		},
		{
			name: "存在しないタスクを指定する",
			args: args{
				ctx: context.Background(),
				id:  "task_xx",
			},
			wantErr: repository.ErrModelNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tc.GetTaskByID(tt.args.ctx, tt.args.id); assert.Equal(t, tt.wantErr, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestClient_UpdateTask(t *testing.T) {
	type args struct {
		ctx context.Context
		t   model.Task
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "改タスク1に更新する",
			args: args{
				ctx: context.Background(),
				t: model.Task{
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
					ProjectID: "project_01",
					UserID:    "user_01",
					Title:     "改タスク1",
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
					UpdatedAt: time.Date(2020, 1, 2, 0, 0, 1, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				ResetTask(t, tc)
			})
			err := tc.UpdateTask(tt.args.ctx, tt.args.t)
			require.NoError(t, err)

			if got, err := tc.GetTaskByID(context.Background(), tt.args.t.ID); assert.NoError(t, err) {
				assert.Equal(t, tt.args.t, got)
			}
		})
	}
}

func TestClient_DeleteTask(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "タスク1を削除する",
			args: args{
				ctx: context.Background(),
				id:  "task_01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				ResetTask(t, tc)
			})
			err := tc.DeleteTask(tt.args.ctx, tt.args.id)
			require.NoError(t, err)

			_, err = tc.GetTaskByID(context.Background(), tt.args.id)
			assert.ErrorIs(t, err, repository.ErrModelNotFound)
		})
	}
}

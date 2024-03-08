package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
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
					Name:        "新タスク",
					Content:     "",
					Priority:    1,
					DueOn:       nil,
					CompletedAt: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})
			err := tc.CreateTask(tt.args.ctx, tt.args.t)
			require.NoError(t, err)

			if got, err := tc.GetTaskByID(context.Background(), tt.args.t.ID); assert.NoError(t, err) {
				assert.Equal(t, tt.args.t, got)
			}
		})
	}
}

func TestClient_ListTasksByUserID(t *testing.T) {
	type args struct {
		ctx       context.Context
		userID    string
		limit     uint
		offset    uint
		projectID *string
		tagID     *string
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
				userID: "user_01",
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
					Name:        "タスク1",
					Content:     "コンテンツ1",
					Priority:    3,
					DueOn:       &dueOn,
					CompletedAt: nil,
				},
			},
		},
		{
			name: "offsetが大きいので、空のスライスを取得する",
			args: args{
				ctx:    context.Background(),
				userID: "user_01",
				limit:  10,
				offset: 1000,
			},
			want: []model.Task{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tc.ListTasksByUserID(tt.args.ctx, tt.args.userID, tt.args.limit, tt.args.offset, tt.args.projectID, tt.args.tagID); assert.NoError(t, err) {
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
				Name:        "タスク1",
				Content:     "コンテンツ1",
				Priority:    3,
				DueOn:       &dueOn,
				CompletedAt: nil,
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
					ProjectID: "project_01",
					UserID:    "user_01",
					Name:      "改タスク1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
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
				_ = fixtures.Load()
			})
			err := tc.DeleteTask(tt.args.ctx, tt.args.id)
			require.NoError(t, err)

			_, err = tc.GetTaskByID(context.Background(), tt.args.id)
			assert.ErrorIs(t, err, repository.ErrModelNotFound)
		})
	}
}

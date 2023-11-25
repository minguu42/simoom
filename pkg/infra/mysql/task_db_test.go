package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
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
			ctx := context.Background()
			t.Cleanup(func() {
				if err := ResetTask(ctx, tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.CreateTask(tt.args.ctx, tt.args.t); err != nil {
				t.Fatalf("%+v", err)
			}

			got, err := tc.GetTaskByID(ctx, tt.args.t.ID)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.args.t, got); diff != "" {
				t.Errorf("created project mismatch (-want +got):\n%s", diff)
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
			got, err := tc.ListTasksByProjectID(tt.args.ctx, tt.args.projectID, tt.args.limit, tt.args.offset)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tc.ListTasksByProjectID mismatch (-want +got):\n%s", diff)
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
			got, err := tc.ListTasksByTagID(tt.args.ctx, tt.args.tagID, tt.args.limit, tt.args.offset)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tc.ListTasksByTagID mismatch (-want +got):\n%s", diff)
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
			got, err := tc.GetTaskByID(tt.args.ctx, tt.args.id)
			if err != nil {
				if errors.Is(err, tt.wantErr) {
					return
				}
				if tt.wantErr != nil {
					t.Fatalf("tc.GetTaskByID error want %s, but got %s", tt.wantErr, err)
				}
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tc.GetTaskByID mismatch (-want +got):\n%s", diff)
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
			ctx := context.Background()
			t.Cleanup(func() {
				if err := ResetTask(ctx, tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.UpdateTask(tt.args.ctx, tt.args.t); err != nil {
				t.Fatalf("%+v", err)
			}

			got, err := tc.GetTaskByID(ctx, tt.args.t.ID)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.args.t, got); diff != "" {
				t.Errorf("updated task mismatch (-want +got):\n%s", diff)
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
			ctx := context.Background()
			t.Cleanup(func() {
				if err := ResetTask(ctx, tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.DeleteTask(tt.args.ctx, tt.args.id); err != nil {
				t.Fatalf("%+v", err)
			}

			if _, err := tc.GetTaskByID(ctx, tt.args.id); !errors.Is(err, repository.ErrModelNotFound) {
				t.Errorf("deleted task exists")
			}
		})
	}
}

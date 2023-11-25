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

func TestClient_CreateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		p   model.Project
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "新プロジェクトを作成する",
			args: args{
				ctx: context.Background(),
				p: model.Project{
					ID:         "project_99",
					UserID:     "user_01",
					Name:       "新プロジェクト",
					Color:      "#000000",
					IsArchived: false,
					CreatedAt:  time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
					UpdatedAt:  time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			t.Cleanup(func() {
				if err := ResetProject(ctx, tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.CreateProject(tt.args.ctx, tt.args.p); err != nil {
				t.Fatalf("%+v", err)
			}

			got, err := tc.GetProjectByID(ctx, tt.args.p.ID)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.args.p, got); diff != "" {
				t.Errorf("created project mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClient_ListProjectsByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
		limit  uint
		offset uint
	}
	tests := []struct {
		name string
		args args
		want []model.Project
	}{
		{
			name: "ユーザIDからプロジェクト一覧を取得する",
			args: args{
				ctx:    context.Background(),
				userID: "user_01",
				limit:  1,
				offset: 0,
			},
			want: []model.Project{
				{
					ID:         "project_02",
					UserID:     "user_01",
					Name:       "プロジェクト2",
					Color:      "#ffffff",
					IsArchived: false,
					CreatedAt:  time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
					UpdatedAt:  time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
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
			want: []model.Project{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tc.ListProjectsByUserID(tt.args.ctx, tt.args.userID, tt.args.limit, tt.args.offset)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tc.ListProjectsByUserID mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClient_GetProjectByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Project
		wantErr error
	}{
		{
			name: "プロジェクト1を取得する",
			args: args{
				ctx: context.Background(),
				id:  "project_01",
			},
			want: model.Project{
				ID:         "project_01",
				UserID:     "user_01",
				Name:       "プロジェクト1",
				Color:      "#1a2b3c",
				IsArchived: false,
				CreatedAt:  time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
				UpdatedAt:  time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			},
		},
		{
			name: "存在しないプロジェクトを指定する",
			args: args{
				ctx: context.Background(),
				id:  "project_xx",
			},
			wantErr: repository.ErrModelNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tc.GetProjectByID(tt.args.ctx, tt.args.id)
			if err != nil {
				if errors.Is(err, tt.wantErr) {
					return
				}
				if tt.wantErr != nil {
					t.Fatalf("tc.GetProjectByID error want %s, but got %s", tt.wantErr, err)
				}
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tc.GetProjectByID mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClient_UpdateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		p   model.Project
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "改プロジェクト1に更新する",
			args: args{
				ctx: context.Background(),
				p: model.Project{
					ID:         "project_01",
					UserID:     "user_01",
					Name:       "改プロジェクト1",
					Color:      "#0f1e2d",
					IsArchived: false,
					CreatedAt:  time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
					UpdatedAt:  time.Date(2020, 1, 2, 0, 0, 1, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			t.Cleanup(func() {
				if err := ResetProject(ctx, tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.UpdateProject(tt.args.ctx, tt.args.p); err != nil {
				t.Fatalf("%+v", err)
			}

			got, err := tc.GetProjectByID(ctx, tt.args.p.ID)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.args.p, got); diff != "" {
				t.Errorf("updated project mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClient_DeleteProject(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "プロジェクト1を削除する",
			args: args{
				ctx: context.Background(),
				id:  "project_01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			t.Cleanup(func() {
				if err := ResetProject(ctx, tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.DeleteProject(tt.args.ctx, tt.args.id); err != nil {
				t.Fatalf("%+v", err)
			}

			if _, err := tc.GetProjectByID(ctx, tt.args.id); !errors.Is(err, repository.ErrModelNotFound) {
				t.Errorf("deleted project exists")
			}
		})
	}
}

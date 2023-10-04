package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
)

var tagCmpOption = cmpopts.IgnoreFields(model.Tag{}, "UpdatedAt")

func TestClient_CreateTag(t *testing.T) {
	type args struct {
		ctx context.Context
		t   model.Tag
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "新タグを作成する",
			args: args{
				ctx: context.Background(),
				t: model.Tag{
					ID:        "tag_99",
					UserID:    "user_01",
					Name:      "新タグ",
					CreatedAt: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			t.Cleanup(func() {
				if err := resetTag(ctx, tc.db); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.CreateTag(tt.args.ctx, tt.args.t); err != nil {
				t.Fatalf("%+v", err)
			}

			got, err := tc.GetTagByID(ctx, tt.args.t.ID)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.args.t, got); diff != "" {
				t.Errorf("created tag mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClient_ListTagsByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
		limit  uint
		offset uint
	}
	tests := []struct {
		name string
		args args
		want []model.Tag
	}{
		{
			name: "ユーザIDからタグ一覧を取得する",
			args: args{
				ctx:    context.Background(),
				userID: "user_01",
				limit:  1,
				offset: 0,
			},
			want: []model.Tag{
				{
					ID:        "tag_01",
					UserID:    "user_01",
					Name:      "タグ1",
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
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
			want: []model.Tag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tc.ListTagsByUserID(tt.args.ctx, tt.args.userID, tt.args.limit, tt.args.offset)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tc.ListTagsByUserID mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClient_GetTagByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Tag
		wantErr error
	}{
		{
			name: "タグ1を取得する",
			args: args{
				ctx: context.Background(),
				id:  "tag_01",
			},
			want: model.Tag{
				ID:        "tag_01",
				UserID:    "user_01",
				Name:      "タグ1",
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
				UpdatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			},
		},
		{
			name: "存在しないタグを指定する",
			args: args{
				ctx: context.Background(),
				id:  "tag_xx",
			},
			wantErr: repository.ErrModelNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tc.GetTagByID(tt.args.ctx, tt.args.id)
			if err != nil {
				if errors.Is(err, tt.wantErr) {
					return
				}
				if tt.wantErr != nil {
					t.Errorf("tc.GetTagByID error want %s, but got %s", tt.wantErr, err)
					return
				}
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tc.GetTagByID mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClient_UpdateTag(t *testing.T) {
	type args struct {
		ctx context.Context
		t   model.Tag
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "改タグ1に更新する",
			args: args{
				ctx: context.Background(),
				t: model.Tag{
					ID:        "tag_01",
					UserID:    "user_01",
					Name:      "改タグ1",
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			t.Cleanup(func() {
				if err := resetTag(ctx, tc.db); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.UpdateTag(tt.args.ctx, tt.args.t); err != nil {
				t.Fatalf("%+v", err)
			}

			got, err := tc.GetTagByID(ctx, tt.args.t.ID)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.args.t, got, tagCmpOption); diff != "" {
				t.Errorf("updated tag mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClient_DeleteTag(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "タグ1を削除する",
			args: args{
				ctx: context.Background(),
				id:  "tag_01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			t.Cleanup(func() {
				if err := resetTag(ctx, tc.db); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.DeleteTag(tt.args.ctx, tt.args.id); err != nil {
				t.Fatalf("%+v", err)
			}

			if _, err := tc.GetTagByID(ctx, tt.args.id); !errors.Is(err, repository.ErrModelNotFound) {
				t.Errorf("deleted tag exists")
			}
		})
	}
}

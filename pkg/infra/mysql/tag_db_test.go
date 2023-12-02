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
			t.Cleanup(func() {
				ResetTag(tc)
			})
			err := tc.CreateTag(tt.args.ctx, tt.args.t)
			require.NoError(t, err)

			if got, err := tc.GetTagByID(context.Background(), tt.args.t.ID); assert.NoError(t, err) {
				assert.Equal(t, tt.args.t, got)
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
			if got, err := tc.ListTagsByUserID(tt.args.ctx, tt.args.userID, tt.args.limit, tt.args.offset); assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
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
			if got, err := tc.GetTagByID(tt.args.ctx, tt.args.id); assert.Equal(t, tt.wantErr, err) {
				assert.Equal(t, tt.want, got)
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
					UpdatedAt: time.Date(2020, 1, 2, 0, 0, 1, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				ResetTag(tc)
			})
			err := tc.UpdateTag(tt.args.ctx, tt.args.t)
			require.NoError(t, err)

			if got, err := tc.GetTagByID(context.Background(), tt.args.t.ID); assert.NoError(t, err) {
				assert.Equal(t, tt.args.t, got)
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
			t.Cleanup(func() {
				ResetTag(tc)
			})
			err := tc.DeleteTag(tt.args.ctx, tt.args.id)
			require.NoError(t, err)

			_, err = tc.GetTagByID(context.Background(), tt.args.id)
			assert.ErrorIs(t, err, repository.ErrModelNotFound)
		})
	}
}

package usecase_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/pointers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTagUsecase_CreateTag(t *testing.T) {
	tag := usecase.NewTag(tc, &model.IDGeneratorMock{GenerateFunc: func() string {
		return "tag_99"
	}})
	type args struct {
		ctx context.Context
		in  usecase.CreateTagInput
	}
	tests := []struct {
		name string
		args args
		want usecase.TagOutput
	}{
		{
			name: "新タグを作成する",
			args: args{
				ctx: tctx,
				in:  usecase.CreateTagInput{Name: "新タグ"},
			},
			want: usecase.TagOutput{Tag: model.Tag{
				ID:     "tag_99",
				UserID: "user_01",
				Name:   "新タグ",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := tag.CreateTag(tt.args.ctx, tt.args.in)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTagUsecase_ListTags(t *testing.T) {
	t.Parallel()

	tag := usecase.NewTag(tc, &model.IDGeneratorMock{})
	type args struct {
		ctx context.Context
		in  usecase.ListTagsInput
	}
	tests := []struct {
		name string
		args args
		want usecase.TagsOutput
	}{
		{
			name: "タグ一覧を取得する",
			args: args{
				ctx: tctx,
				in: usecase.ListTagsInput{
					Limit:  10,
					Offset: 0,
				},
			},
			want: usecase.TagsOutput{
				Tags: []model.Tag{
					{ID: "tag_01", UserID: "user_01", Name: "タグ1"},
					{ID: "tag_02", UserID: "user_01", Name: "タグ2"},
				},
				HasNext: false,
			},
		},
		{
			name: "limitで制限をかける",
			args: args{
				ctx: tctx,
				in: usecase.ListTagsInput{
					Limit:  1,
					Offset: 0,
				},
			},
			want: usecase.TagsOutput{
				Tags: []model.Tag{
					{ID: "tag_01", UserID: "user_01", Name: "タグ1"},
				},
				HasNext: true,
			},
		},
		{
			name: "limitとoffsetでページングを行う",
			args: args{
				ctx: tctx,
				in: usecase.ListTagsInput{
					Limit:  1,
					Offset: 1,
				},
			},
			want: usecase.TagsOutput{
				Tags: []model.Tag{
					{ID: "tag_02", UserID: "user_01", Name: "タグ2"},
				},
				HasNext: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tag.ListTags(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tag.ListTags mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTagUsecase_UpdateTag(t *testing.T) {
	tag := usecase.NewTag(tc, &model.IDGeneratorMock{})
	type args struct {
		ctx context.Context
		in  usecase.UpdateTagInput
	}
	tests := []struct {
		name    string
		args    args
		want    usecase.TagOutput
		wantErr apperr.Error
	}{
		{
			name: "改タグ1に更新する",
			args: args{
				ctx: tctx,
				in: usecase.UpdateTagInput{
					ID:   "tag_01",
					Name: pointers.Ref("改タグ1"),
				},
			},
			want: usecase.TagOutput{Tag: model.Tag{
				ID:     "tag_01",
				UserID: "user_01",
				Name:   "改タグ1",
			}},
		},
		{
			name: "指定したタグが存在しない場合はエラーを返す",
			args: args{
				ctx: tctx,
				in: usecase.UpdateTagInput{
					ID:   "tag_99",
					Name: pointers.Ref("改タグ99"),
				},
			},
			wantErr: apperr.ErrTagNotFound(nil),
		},
		{
			name: "指定したタグを所有していない場合はエラーを返す",
			args: args{
				ctx: tctxUser2,
				in: usecase.UpdateTagInput{
					ID:   "tag_01",
					Name: pointers.Ref("改タグ1"),
				},
			},
			wantErr: apperr.ErrTagNotFound(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := tag.UpdateTag(tt.args.ctx, tt.args.in)
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

func TestTagUsecase_DeleteTag(t *testing.T) {
	tag := usecase.NewTag(tc, &model.IDGeneratorMock{})
	type args struct {
		ctx context.Context
		in  usecase.DeleteTagInput
	}
	tests := []struct {
		name string
		args args
		want apperr.Error
	}{
		{
			name: "タグ1を削除する",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteTagInput{ID: "tag_01"},
			},
		},
		{
			name: "指定したタグが存在しない場合はエラーを返す",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteTagInput{ID: "tag_99"},
			},
			want: apperr.ErrTagNotFound(nil),
		},
		{
			name: "指定したタグを所有していない場合はエラーを返す",
			args: args{
				ctx: tctxUser2,
				in:  usecase.DeleteTagInput{ID: "tag_01"},
			},
			want: apperr.ErrTagNotFound(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			err := tag.DeleteTag(tt.args.ctx, tt.args.in)
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

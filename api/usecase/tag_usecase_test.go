package usecase_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/pointers"
)

var createTagOption = cmpopts.IgnoreFields(usecase.TagOutput{}, "Tag.ID")

func TestTagUsecase_CreateTag(t *testing.T) {
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
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got, createTagOption); diff != "" {
				t.Errorf("tag.CreateTag mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTagUsecase_ListTags(t *testing.T) {
	t.Parallel()
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
			name: "タスク一覧を表示する",
			args: args{
				ctx: tctx,
				in: usecase.ListTagsInput{
					Limit:  1,
					Offset: 0,
				},
			},
			want: usecase.TagsOutput{
				Tags: []model.Tag{
					{
						ID:     "tag_01",
						UserID: "user_01",
						Name:   "タグ1",
					},
				},
				HasNext: true,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
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
	type args struct {
		ctx context.Context
		in  usecase.UpdateTagInput
	}
	tests := []struct {
		name string
		args args
		want usecase.TagOutput
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := tag.UpdateTag(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tag.UpdateTag mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTagUsecase_DeleteTag(t *testing.T) {
	type args struct {
		ctx context.Context
		in  usecase.DeleteTagInput
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "タグ1を削除する",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteTagInput{ID: "tag_01"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			if err := tag.DeleteTag(tt.args.ctx, tt.args.in); err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

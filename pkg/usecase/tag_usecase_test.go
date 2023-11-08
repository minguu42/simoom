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
	createTagOption = cmpopts.IgnoreFields(usecase.TagOutput{},
		"Tag.ID",
		"Tag.CreatedAt",
		"Tag.UpdatedAt",
	)
	updateTagOption = cmpopts.IgnoreFields(usecase.TagOutput{}, "Tag.UpdatedAt")
)

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
				ctx: ctx,
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
				if err := mysql.ResetTag(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
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
				ctx: ctx,
				in: usecase.ListTagsInput{
					Limit:  1,
					Offset: 0,
				},
			},
			want: usecase.TagsOutput{
				Tags: []model.Tag{
					{
						ID:        "tag_01",
						UserID:    "user_01",
						Name:      "タグ1",
						CreatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
						UpdatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
					},
				},
				HasNext: false, // TODO: ページングの実装後に対応する
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
				ctx: ctx,
				in: usecase.UpdateTagInput{
					ID:   "tag_01",
					Name: pointers.Ref("改タグ1"),
				},
			},
			want: usecase.TagOutput{Tag: model.Tag{
				ID:        "tag_01",
				UserID:    "user_01",
				Name:      "改タグ1",
				CreatedAt: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetTag(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})

			got, err := tag.UpdateTag(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got, updateTagOption); diff != "" {
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
				ctx: ctx,
				in:  usecase.DeleteTagInput{ID: "tag_01"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetTag(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})

			if err := tag.DeleteTag(tt.args.ctx, tt.args.in); err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

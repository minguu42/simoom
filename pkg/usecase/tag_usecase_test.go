package usecase_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/infra/mysql"
	"github.com/minguu42/simoom/pkg/pointers"
	"github.com/minguu42/simoom/pkg/usecase"
	"github.com/stretchr/testify/assert"
)

var createTagOption = cmpopts.IgnoreFields(usecase.TagOutput{}, "Tag.ID")

func TestCreateTagInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.CreateTagInput
		hasError bool
	}{
		{
			name: "nameに空の文字列は指定できない",
			in: usecase.CreateTagInput{
				Name: "",
			},
			hasError: true,
		},
		{
			name: "nameに1文字の文字列は指定できる",
			in: usecase.CreateTagInput{
				Name: "A",
			},
			hasError: false,
		},
		{
			name: "nameに19文字の文字列は指定できる",
			in: usecase.CreateTagInput{
				Name: "やったね。この文字列の長さは19です。",
			},
			hasError: false,
		},
		{
			name: "nameに20文字以上の文字列は指定できない",
			in: usecase.CreateTagInput{
				Name: "やったね。この文字列の長さは20です！。",
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()
			assert.Equal(t, tt.hasError, err != nil)
		})
	}
}

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
				mysql.ResetTag(t, tc)
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

func TestListTagsInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.ListTagsInput
		hasError bool
	}{
		{
			name: "limitに0は指定できない",
			in: usecase.ListTagsInput{
				Limit:  0,
				Offset: 0,
			},
			hasError: true,
		},
		{
			name: "limitに1は指定できる",
			in: usecase.ListTagsInput{
				Limit:  1,
				Offset: 0,
			},
			hasError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()
			assert.Equal(t, tt.hasError, err != nil)
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

func TestUpdateTagInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.UpdateTagInput
		hasError bool
	}{
		{
			name: "idに25文字以下の文字列は指定できない",
			in: usecase.UpdateTagInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxxx",
				Name: pointers.Ref("テストタグ"),
			},
			hasError: true,
		},
		{
			name: "idは26文字の文字列である",
			in: usecase.UpdateTagInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("テストタグ"),
			},
			hasError: false,
		},
		{
			name: "idに27文字以上の文字列は指定できない",
			in: usecase.UpdateTagInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-xx",
				Name: pointers.Ref("テストタグ"),
			},
			hasError: true,
		},
		{
			name: "いずれかの引数は必要である",
			in: usecase.UpdateTagInput{
				ID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
			},
			hasError: true,
		},
		{
			name: "nameに空の文字列は指定できない",
			in: usecase.UpdateTagInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref(""),
			},
			hasError: true,
		},
		{
			name: "nameに1文字の文字列は指定できる",
			in: usecase.UpdateTagInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("A"),
			},
			hasError: false,
		},
		{
			name: "nameに19文字の文字列は指定できる",
			in: usecase.UpdateTagInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("やったね。この文字列の長さは19です。"),
			},
			hasError: false,
		},
		{
			name: "nameに20文字以上の文字列は指定できない",
			in: usecase.UpdateTagInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("やったね。この文字列の長さは20です！。"),
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()
			assert.Equal(t, tt.hasError, err != nil)
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
				mysql.ResetTag(t, tc)
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

func TestDeleteTagInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.DeleteTagInput
		hasError bool
	}{
		{
			name: "idに25文字以下の文字列は指定できない",
			in: usecase.DeleteTagInput{
				ID: "xxxx-xxxx-xxxx-xxxx-xxxxx",
			},
			hasError: true,
		},
		{
			name: "idは26文字の文字列である",
			in: usecase.DeleteTagInput{
				ID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
			},
			hasError: false,
		},
		{
			name: "idに27文字以上の文字列は指定できない",
			in: usecase.DeleteTagInput{
				ID: "xxxx-xxxx-xxxx-xxxx-xxxx-xx",
			},
			hasError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()
			assert.Equal(t, tt.hasError, err != nil)
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
				mysql.ResetTag(t, tc)
			})

			if err := tag.DeleteTag(tt.args.ctx, tt.args.in); err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

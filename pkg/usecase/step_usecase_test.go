package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/minguu42/simoom/lib/go/pointers"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/usecase"
	"github.com/stretchr/testify/assert"
)

var createStepOption = cmpopts.IgnoreFields(usecase.StepOutput{}, "Step.ID")

func TestCreateStepInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.CreateStepInput
		hasError bool
	}{
		{
			name: "task_idに25文字以下の文字列は指定できない",
			in: usecase.CreateStepInput{
				TaskID: "xxxx-xxxx-xxxx-xxxx-xxxxx",
				Name:   "テストステップ",
			},
			hasError: true,
		},
		{
			name: "task_idに26文字の文字列である",
			in: usecase.CreateStepInput{
				TaskID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name:   "テストステップ",
			},
			hasError: false,
		},
		{
			name: "task_idに27文字以上の文字列は指定できない",
			in: usecase.CreateStepInput{
				TaskID: "xxxx-xxxx-xxxx-xxxx-xxxx-xx",
				Name:   "テストステップ",
			},
			hasError: true,
		},
		{
			name: "nameに空の文字列は指定できない",
			in: usecase.CreateStepInput{
				TaskID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name:   "",
			},
			hasError: true,
		},
		{
			name: "nameに1文字の文字列は指定できる",
			in: usecase.CreateStepInput{
				TaskID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name:   "A",
			},
			hasError: false,
		},
		{
			name: "nameに80文字の文字列は指定できる",
			in: usecase.CreateStepInput{
				TaskID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name:   "これはとても長い文字列です。なんとその長さ80文字、さぁここからどんどん伸ばして行きますよ。あいうえおーかきくけこーさしすせそ。さぁあとちょっとです。もう少しだ",
			},
			hasError: false,
		},
		{
			name: "nameに81文字以上の文字列は指定できない",
			in: usecase.CreateStepInput{
				TaskID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name:   "これはとても長い文字列です。なんとその長さ81文字、さぁここからどんどん伸ばして行きますよ。あいうえおーかきくけこーさしすせそ。さぁあとちょっとです。もう少しだよ",
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

func TestStepUsecase_CreateStep(t *testing.T) {
	type args struct {
		ctx context.Context
		in  usecase.CreateStepInput
	}
	tests := []struct {
		name string
		args args
		want usecase.StepOutput
	}{
		{
			name: "新ステップを作成する",
			args: args{
				ctx: tctx,
				in: usecase.CreateStepInput{
					TaskID: "task_01",
					Name:   "新ステップ",
				},
			},
			want: usecase.StepOutput{Step: model.Step{
				UserID:      "user_01",
				TaskID:      "task_01",
				Name:        "新ステップ",
				CompletedAt: nil,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := step.CreateStep(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got, createStepOption); diff != "" {
				t.Errorf("step.CreateStep mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUpdateStepInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.UpdateStepInput
		hasError bool
	}{
		{
			name: "idに25文字以下の文字列は指定できない",
			in: usecase.UpdateStepInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxxx",
				Name: pointers.Ref("テストステップ"),
			},
			hasError: true,
		},
		{
			name: "idは26文字の文字列である",
			in: usecase.UpdateStepInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("テストステップ"),
			},
			hasError: false,
		},
		{
			name: "idに27文字以上の文字列は指定できない",
			in: usecase.UpdateStepInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-xx",
				Name: pointers.Ref("テストステップ"),
			},
			hasError: true,
		},
		{
			name: "いずれかの引数は必要である",
			in: usecase.UpdateStepInput{
				ID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
			},
			hasError: true,
		},
		{
			name: "nameに空の文字列は指定できない",
			in: usecase.UpdateStepInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref(""),
			},
			hasError: true,
		},
		{
			name: "nameに1文字の文字列は指定できる",
			in: usecase.UpdateStepInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("A"),
			},
			hasError: false,
		},
		{
			name: "nameに80文字の文字列は指定できる",
			in: usecase.UpdateStepInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("これはとても長い文字列です。なんとその長さ80文字、さぁここからどんどん伸ばして行きますよ。あいうえおーかきくけこーさしすせそ。さぁあとちょっとです。もう少しだ"),
			},
			hasError: false,
		},
		{
			name: "nameに81文字以上の文字列は指定できない",
			in: usecase.UpdateStepInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("これはとても長い文字列です。なんとその長さ81文字、さぁここからどんどん伸ばして行きますよ。あいうえおーかきくけこーさしすせそ。さぁあとちょっとです。もう少しだよ"),
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

func TestStepUsecase_UpdateStep(t *testing.T) {
	type args struct {
		ctx context.Context
		in  usecase.UpdateStepInput
	}
	tests := []struct {
		name string
		args args
		want usecase.StepOutput
	}{
		{
			name: "改ステップ1に更新する",
			args: args{
				ctx: tctx,
				in: usecase.UpdateStepInput{
					ID:          "step_01",
					Name:        pointers.Ref("改ステップ1"),
					CompletedAt: pointers.Ref(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
			want: usecase.StepOutput{Step: model.Step{
				ID:          "step_01",
				UserID:      "user_01",
				TaskID:      "task_01",
				Name:        "改ステップ1",
				CompletedAt: pointers.Ref(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := step.UpdateStep(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("step.UpdateStep mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDeleteStepInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.DeleteStepInput
		hasError bool
	}{
		{
			name: "idに25文字以下の文字列は指定できない",
			in: usecase.DeleteStepInput{
				ID: "xxxx-xxxx-xxxx-xxxx-xxxxx",
			},
			hasError: true,
		},
		{
			name: "idは26文字の文字列である",
			in: usecase.DeleteStepInput{
				ID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
			},
			hasError: false,
		},
		{
			name: "idに27文字以上の文字列は指定できない",
			in: usecase.DeleteStepInput{
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

func TestStepUsecase_DeleteStep(t *testing.T) {
	type args struct {
		ctx context.Context
		in  usecase.DeleteStepInput
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ステップ1を削除する",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteStepInput{ID: "step_01"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			if err := step.DeleteStep(tt.args.ctx, tt.args.in); err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

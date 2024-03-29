package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/pointers"
)

var createStepOption = cmpopts.IgnoreFields(usecase.StepOutput{}, "Step.ID")

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

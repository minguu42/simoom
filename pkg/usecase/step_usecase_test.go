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
	createStepOption = cmpopts.IgnoreFields(usecase.StepOutput{},
		"Step.ID",
		"Step.CreatedAt",
		"Step.UpdatedAt",
	)
	updateStepOption = cmpopts.IgnoreFields(usecase.StepOutput{},
		"Step.CreatedAt",
		"Step.UpdatedAt",
	)
)

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
				ctx: ctx,
				in: usecase.CreateStepInput{
					TaskID: "task_01",
					Title:  "新ステップ",
				},
			},
			want: usecase.StepOutput{Step: model.Step{
				UserID:      "user_01",
				TaskID:      "task_01",
				Title:       "新ステップ",
				CompletedAt: nil,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetStep(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
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
				ctx: ctx,
				in: usecase.UpdateStepInput{
					ID:          "step_01",
					Title:       pointers.Ref("改ステップ1"),
					CompletedAt: pointers.Ref(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
			want: usecase.StepOutput{Step: model.Step{
				ID:          "step_01",
				UserID:      "user_01",
				TaskID:      "task_01",
				Title:       "改ステップ1",
				CompletedAt: pointers.Ref(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetStep(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})

			got, err := step.UpdateStep(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got, updateStepOption); diff != "" {
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
				ctx: ctx,
				in:  usecase.DeleteStepInput{ID: "step_01"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetStep(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})

			if err := step.DeleteStep(tt.args.ctx, tt.args.in); err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

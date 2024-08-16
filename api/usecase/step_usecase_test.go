package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/pointers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStepUsecase_CreateStep(t *testing.T) {
	step := usecase.NewStep(tc, &model.IDGeneratorMock{GenerateFunc: func() string {
		return "step_99"
	}})
	type args struct {
		ctx context.Context
		in  usecase.CreateStepInput
	}
	tests := []struct {
		name    string
		args    args
		want    usecase.StepOutput
		wantErr apperr.Error
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
				ID:          "step_99",
				UserID:      "user_01",
				TaskID:      "task_01",
				Name:        "新ステップ",
				CompletedAt: nil,
			}},
		},
		{
			name: "指定したタスクが存在しない場合はエラーを返す",
			args: args{
				ctx: tctx,
				in: usecase.CreateStepInput{
					TaskID: "task_99",
					Name:   "新ステップ",
				},
			},
			wantErr: apperr.ErrTaskNotFound(nil),
		},
		{
			name: "指定したタスクを所有していない場合はエラーを返す",
			args: args{
				ctx: tctxUser2,
				in: usecase.CreateStepInput{
					TaskID: "task_01",
					Name:   "新ステップ",
				},
			},
			wantErr: apperr.ErrTaskNotFound(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := step.CreateStep(tt.args.ctx, tt.args.in)
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

func TestStepUsecase_UpdateStep(t *testing.T) {
	step := usecase.NewStep(tc, &model.IDGeneratorMock{})
	type args struct {
		ctx context.Context
		in  usecase.UpdateStepInput
	}
	tests := []struct {
		name    string
		args    args
		want    usecase.StepOutput
		wantErr apperr.Error
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
		{
			name: "指定したステップが存在しない場合はエラーを返す",
			args: args{
				ctx: tctx,
				in: usecase.UpdateStepInput{
					ID:   "step_99",
					Name: pointers.Ref("改ステップ99"),
				},
			},
			wantErr: apperr.ErrStepNotFound(nil),
		},
		{
			name: "指定したステップを所有していない場合はエラーを返す",
			args: args{
				ctx: tctxUser2,
				in: usecase.UpdateStepInput{
					ID:   "step_01",
					Name: pointers.Ref("改ステップ1"),
				},
			},
			wantErr: apperr.ErrStepNotFound(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := step.UpdateStep(tt.args.ctx, tt.args.in)
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

func TestStepUsecase_DeleteStep(t *testing.T) {
	step := usecase.NewStep(tc, &model.IDGeneratorMock{})
	type args struct {
		ctx context.Context
		in  usecase.DeleteStepInput
	}
	tests := []struct {
		name string
		args args
		want apperr.Error
	}{
		{
			name: "ステップ1を削除する",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteStepInput{ID: "step_01"},
			},
		},
		{
			name: "指定したステップが存在しない場合はエラーを返す",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteStepInput{ID: "step_99"},
			},
			want: apperr.ErrStepNotFound(nil),
		},
		{
			name: "指定したステップを所有していない場合はエラーを返す",
			args: args{
				ctx: tctxUser2,
				in:  usecase.DeleteStepInput{ID: "step_01"},
			},
			want: apperr.ErrStepNotFound(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			err := step.DeleteStep(tt.args.ctx, tt.args.in)
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

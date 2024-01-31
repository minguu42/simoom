package mysql

import (
	"context"
	"testing"

	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_CreateStep(t *testing.T) {
	type args struct {
		ctx context.Context
		s   model.Step
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "新ステップを作成する",
			args: args{
				ctx: context.Background(),
				s: model.Step{
					ID:          "step_99",
					UserID:      "user_01",
					TaskID:      "task_01",
					Name:        "新ステップ",
					CompletedAt: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			t.Cleanup(func() {
				ResetStep(t, tc)
			})
			err := tc.CreateStep(tt.args.ctx, tt.args.s)
			require.NoError(t, err)

			if got, err := tc.GetStepByID(context.Background(), tt.args.s.ID); assert.NoError(t, err) {
				assert.Equal(t, tt.args.s, got)
			}
		})
	}
}

func TestClient_GetStepByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Step
		wantErr error
	}{
		{
			name: "ステップ1を取得する",
			args: args{
				ctx: context.Background(),
				id:  "step_01",
			},
			want: model.Step{
				ID:          "step_01",
				UserID:      "user_01",
				TaskID:      "task_01",
				Name:        "ステップ1",
				CompletedAt: nil,
			},
		},
		{
			name: "存在しないステップを指定する",
			args: args{
				ctx: context.Background(),
				id:  "step_xx",
			},
			wantErr: repository.ErrModelNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tc.GetStepByID(tt.args.ctx, tt.args.id); assert.Equal(t, tt.wantErr, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestClient_UpdateStep(t *testing.T) {
	type args struct {
		ctx context.Context
		s   model.Step
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "改ステップ1に更新する",
			args: args{
				ctx: context.Background(),
				s: model.Step{
					ID:          "step_01",
					UserID:      "user_01",
					TaskID:      "task_01",
					Name:        "改ステップ1",
					CompletedAt: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				ResetStep(t, tc)
			})
			err := tc.UpdateStep(tt.args.ctx, tt.args.s)
			require.NoError(t, err)

			if got, err := tc.GetStepByID(context.Background(), tt.args.s.ID); assert.NoError(t, err) {
				assert.Equal(t, tt.args.s, got)
			}
		})
	}
}

func TestClient_DeleteStep(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ステップ1を削除する",
			args: args{
				ctx: context.Background(),
				id:  "step_01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				ResetStep(t, tc)
			})
			err := tc.DeleteStep(tt.args.ctx, tt.args.id)
			require.NoError(t, err)

			_, err = tc.GetStepByID(context.Background(), tt.args.id)
			assert.ErrorIs(t, err, repository.ErrModelNotFound)
		})
	}
}

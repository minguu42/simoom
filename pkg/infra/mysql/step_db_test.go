package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/minguu42/simoom/pkg/domain/model"
	"github.com/minguu42/simoom/pkg/domain/repository"
)

var stepCmpOption = cmpopts.IgnoreFields(model.Step{}, "UpdatedAt")

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
					Title:       "新ステップ",
					CompletedAt: nil,
					CreatedAt:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			ctx := context.Background()
			t.Cleanup(func() {
				if err := resetStep(ctx, tc.db); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.CreateStep(tt.args.ctx, tt.args.s); err != nil {
				t.Fatalf("%+v", err)
			}

			got, err := tc.GetStepByID(ctx, tt.args.s.ID)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.args.s, got); diff != "" {
				t.Errorf("created step mismatch (-want +got):\n%s", diff)
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
				Title:       "ステップ1",
				CompletedAt: nil,
				CreatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
				UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
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
			got, err := tc.GetStepByID(tt.args.ctx, tt.args.id)
			if err != nil {
				if errors.Is(err, tt.wantErr) {
					return
				}
				if tt.wantErr != nil {
					t.Errorf("tc.GetStepByID error want %s, but got %s", tt.wantErr, err)
				}
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("tc.GetStepByID mismatch (-want +got):\n%s", diff)
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
					Title:       "改ステップ1",
					CompletedAt: nil,
					CreatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			t.Cleanup(func() {
				if err := resetStep(ctx, tc.db); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.UpdateStep(tt.args.ctx, tt.args.s); err != nil {
				t.Fatalf("%+v", err)
			}

			got, err := tc.GetStepByID(ctx, tt.args.s.ID)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.args.s, got, stepCmpOption); diff != "" {
				t.Errorf("updated step mismatch (-want +got):\n%s", diff)
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
			ctx := context.Background()
			t.Cleanup(func() {
				if err := resetStep(ctx, tc.db); err != nil {
					t.Fatalf("%+v", err)
				}
			})
			if err := tc.DeleteStep(tt.args.ctx, tt.args.id); err != nil {
				t.Fatalf("%+v", err)
			}

			if _, err := tc.GetStepByID(ctx, tt.args.id); !errors.Is(err, repository.ErrModelNotFound) {
				t.Errorf("deleted project exists")
			}
		})
	}
}

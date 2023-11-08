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
	createProjectOption = cmpopts.IgnoreFields(usecase.ProjectOutput{},
		"Project.ID",
		"Project.CreatedAt",
		"Project.UpdatedAt",
	)
	updateProjectOption = cmpopts.IgnoreFields(usecase.ProjectOutput{},
		"Project.CreatedAt",
		"Project.UpdatedAt",
	)
)

func TestProjectUsecase_CreateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		in  usecase.CreateProjectInput
	}
	tests := []struct {
		name string
		args args
		want usecase.ProjectOutput
	}{
		{
			name: "新プロジェクトを作成する",
			args: args{
				ctx: ctx,
				in: usecase.CreateProjectInput{
					Name:  "新プロジェクト",
					Color: "#f8b500",
				},
			},
			want: usecase.ProjectOutput{Project: model.Project{
				UserID:     "user_01",
				Name:       "新プロジェクト",
				Color:      "#f8b500",
				IsArchived: false,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetProject(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})

			got, err := project.CreateProject(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got, createProjectOption); diff != "" {
				t.Errorf("project.CreateProject mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestProjectUsecase_ListProjects(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
		in  usecase.ListProjectsInput
	}
	tests := []struct {
		name string
		args args
		want usecase.ProjectsOutput
	}{
		{
			name: "タスク一覧を表示する",
			args: args{
				ctx: ctx,
				in: usecase.ListProjectsInput{
					Limit:  1,
					Offset: 0,
				},
			},
			want: usecase.ProjectsOutput{
				Projects: []model.Project{
					{
						ID:         "project_02",
						UserID:     "user_01",
						Name:       "プロジェクト2",
						Color:      "#ffffff",
						IsArchived: false,
						CreatedAt:  time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
						UpdatedAt:  time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
					},
				},
				HasNext: false, // TODO: ページングの実装後に対応する
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := project.ListProjects(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("project.ListProjects mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestProjectUsecase_UpdateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		in  usecase.UpdateProjectInput
	}
	tests := []struct {
		name string
		args args
		want usecase.ProjectOutput
	}{
		{
			name: "",
			args: args{
				ctx: ctx,
				in: usecase.UpdateProjectInput{
					ID:         "project_01",
					Name:       pointers.Ref("改プロジェクト1"),
					Color:      pointers.Ref("#0f1e2d"),
					IsArchived: pointers.Ref(true),
				},
			},
			want: usecase.ProjectOutput{Project: model.Project{
				ID:         "project_01",
				UserID:     "user_01",
				Name:       "改プロジェクト1",
				Color:      "#0f1e2d",
				IsArchived: true,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetProject(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})

			got, err := project.UpdateProject(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got, updateProjectOption); diff != "" {
				t.Errorf("project.UpdateProject mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestProjectUsecase_DeleteProject(t *testing.T) {
	type args struct {
		ctx context.Context
		in  usecase.DeleteProjectInput
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "プロジェクト1を削除する",
			args: args{
				ctx: ctx,
				in:  usecase.DeleteProjectInput{ID: "project_01"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysql.ResetProject(context.Background(), tc); err != nil {
					t.Fatalf("%+v", err)
				}
			})

			if err := project.DeleteProject(tt.args.ctx, tt.args.in); err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

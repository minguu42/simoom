package usecase_test

import (
	"context"
	"testing"

	"github.com/minguu42/simoom/api/apperr"
	"github.com/minguu42/simoom/api/domain"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/usecase"
	"github.com/minguu42/simoom/lib/go/pointers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectUsecase_CreateProject(t *testing.T) {
	project := usecase.NewProject(tc, &domain.IDGeneratorMock{GenerateFunc: func() string {
		return "project_99"
	}})
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
				ctx: tctx,
				in: usecase.CreateProjectInput{
					Name:  "新プロジェクト",
					Color: "#f8b500",
				},
			},
			want: usecase.ProjectOutput{Project: model.Project{
				ID:     "project_99",
				UserID: "user_01",
				Name:   "新プロジェクト",
				Color:  "#f8b500",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := project.CreateProject(tt.args.ctx, tt.args.in)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProjectUsecase_ListProjects(t *testing.T) {
	t.Parallel()

	project := usecase.NewProject(tc, &domain.IDGeneratorMock{})
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
			name: "プロジェクト一覧を取得する",
			args: args{
				ctx: tctx,
				in: usecase.ListProjectsInput{
					Limit:  10,
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
					},
					{
						ID:         "project_01",
						UserID:     "user_01",
						Name:       "プロジェクト1",
						Color:      "#1a2b3c",
						IsArchived: false,
					},
				},
				HasNext: false,
			},
		},
		{
			name: "limitで制限をかける",
			args: args{
				ctx: tctx,
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
					},
				},
				HasNext: true,
			},
		},
		{
			name: "limitとoffsetでページングを行う",
			args: args{
				ctx: tctx,
				in: usecase.ListProjectsInput{
					Limit:  1,
					Offset: 1,
				},
			},
			want: usecase.ProjectsOutput{
				Projects: []model.Project{
					{
						ID:         "project_01",
						UserID:     "user_01",
						Name:       "プロジェクト1",
						Color:      "#1a2b3c",
						IsArchived: false,
					},
				},
				HasNext: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := project.ListProjects(tt.args.ctx, tt.args.in)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProjectUsecase_UpdateProject(t *testing.T) {
	project := usecase.NewProject(tc, &domain.IDGeneratorMock{})
	type args struct {
		ctx context.Context
		in  usecase.UpdateProjectInput
	}
	tests := []struct {
		name    string
		args    args
		want    usecase.ProjectOutput
		wantErr apperr.Error
	}{
		{
			name: "改プロジェクト1に更新する",
			args: args{
				ctx: tctx,
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
		{
			name: "指定したプロジェクトが存在しない場合はエラーを返す",
			args: args{
				ctx: tctx,
				in: usecase.UpdateProjectInput{
					ID:   "project_99",
					Name: pointers.Ref("改プロジェクト99"),
				},
			},
			wantErr: apperr.ErrProjectNotFound(nil),
		},
		{
			name: "指定したプロジェクトを所有していない場合はエラーを返す",
			args: args{
				ctx: tctxUser2,
				in: usecase.UpdateProjectInput{
					ID:   "project_01",
					Name: pointers.Ref("改プロジェクト1"),
				},
			},
			wantErr: apperr.ErrProjectNotFound(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			got, err := project.UpdateProject(tt.args.ctx, tt.args.in)
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

func TestProjectUsecase_DeleteProject(t *testing.T) {
	project := usecase.NewProject(tc, &domain.IDGeneratorMock{})
	type args struct {
		ctx context.Context
		in  usecase.DeleteProjectInput
	}
	tests := []struct {
		name string
		args args
		want apperr.Error
	}{
		{
			name: "プロジェクト1を削除する",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteProjectInput{ID: "project_01"},
			},
		},
		{
			name: "指定したプロジェクトが存在しない場合はエラーを返す",
			args: args{
				ctx: tctx,
				in:  usecase.DeleteProjectInput{ID: "project_99"},
			},
			want: apperr.ErrProjectNotFound(nil),
		},
		{
			name: "指定したプロジェクトを所有していない場合はエラーを返す",
			args: args{
				ctx: tctxUser2,
				in:  usecase.DeleteProjectInput{ID: "project_01"},
			},
			want: apperr.ErrProjectNotFound(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				_ = fixtures.Load()
			})

			err := project.DeleteProject(tt.args.ctx, tt.args.in)
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

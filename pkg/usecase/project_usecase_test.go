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

var createProjectOption = cmpopts.IgnoreFields(usecase.ProjectOutput{}, "Project.ID")

func TestCreateProjectInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.CreateProjectInput
		hasError bool
	}{
		{
			name: "nameに空の文字列は指定できない",
			in: usecase.CreateProjectInput{
				Name:  "",
				Color: "#000000",
			},
			hasError: true,
		},
		{
			name: "nameに1文字の文字列は指定できる",
			in: usecase.CreateProjectInput{
				Name:  "A",
				Color: "#000000",
			},
			hasError: false,
		},
		{
			name: "nameに20文字の文字列は指定できる",
			in: usecase.CreateProjectInput{
				Name:  "あいうえおかきくけこさしすせそたちつてと",
				Color: "#000000",
			},
			hasError: false,
		},
		{
			name: "nameに21文字以上の文字列は指定できない",
			in: usecase.CreateProjectInput{
				Name:  "あいうえおかきくけこさしすせそたちつてとな",
				Color: "#000000",
			},
			hasError: true,
		},
		{
			name: "colorは#000000の形式で指定できる",
			in: usecase.CreateProjectInput{
				Name:  "テストプロジェクト",
				Color: "#a1b2c3",
			},
			hasError: false,
		},
		{
			name: "colorに色の名前は指定できない",
			in: usecase.CreateProjectInput{
				Name:  "テストプロジェクト",
				Color: "red",
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
				ctx: tctx,
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
				mysql.ResetProject(t, tc)
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

func TestListProjectsInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.ListProjectsInput
		hasError bool
	}{
		{
			name: "limitに0は指定できない",
			in: usecase.ListProjectsInput{
				Limit:  0,
				Offset: 0,
			},
			hasError: true,
		},
		{
			name: "limitに1は指定できる",
			in: usecase.ListProjectsInput{
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
	}
	for _, tt := range tests {
		tt := tt
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

func TestUpdateProjectInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.UpdateProjectInput
		hasError bool
	}{
		{
			name: "idに25文字以下の文字列は指定できない",
			in: usecase.UpdateProjectInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxxx",
				Name: pointers.Ref("テストプロジェクト"),
			},
			hasError: true,
		},
		{
			name: "idは26文字の文字列である",
			in: usecase.UpdateProjectInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("テストプロジェクト"),
			},
			hasError: false,
		},
		{
			name: "idに27文字以上の文字列は指定できない",
			in: usecase.UpdateProjectInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-xx",
				Name: pointers.Ref("テストプロジェクト"),
			},
			hasError: true,
		},
		{
			name: "いずれかの引数は必要である",
			in: usecase.UpdateProjectInput{
				ID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
			},
			hasError: true,
		},
		{
			name: "nameに空の文字列は指定できない",
			in: usecase.UpdateProjectInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref(""),
			},
			hasError: true,
		},
		{
			name: "nameに1文字の文字列は指定できる",
			in: usecase.UpdateProjectInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("A"),
			},
			hasError: false,
		},
		{
			name: "nameに20文字の文字列は指定できる",
			in: usecase.UpdateProjectInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("あいうえおかきくけこさしすせそたちつてと"),
			},
			hasError: false,
		},
		{
			name: "nameに21文字以上の文字列は指定できない",
			in: usecase.UpdateProjectInput{
				ID:   "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Name: pointers.Ref("あいうえおかきくけこさしすせそたちつてとな"),
			},
			hasError: true,
		},
		{
			name: "colorは#000000の形式で指定できる",
			in: usecase.UpdateProjectInput{
				ID:    "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Color: pointers.Ref("#a1b2c3"),
			},
			hasError: false,
		},
		{
			name: "colorに色の名前は指定できない",
			in: usecase.UpdateProjectInput{
				ID:    "xxxx-xxxx-xxxx-xxxx-xxxx-x",
				Color: pointers.Ref("red"),
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mysql.ResetProject(t, tc)
			})

			got, err := project.UpdateProject(tt.args.ctx, tt.args.in)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("project.UpdateProject mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDeleteProjectInput_Validate(t *testing.T) {
	tests := []struct {
		name     string
		in       usecase.DeleteProjectInput
		hasError bool
	}{
		{
			name: "idに25文字以下の文字列は指定できない",
			in: usecase.DeleteProjectInput{
				ID: "xxxx-xxxx-xxxx-xxxx-xxxxx",
			},
			hasError: true,
		},
		{
			name: "idは26文字の文字列である",
			in: usecase.DeleteProjectInput{
				ID: "xxxx-xxxx-xxxx-xxxx-xxxx-x",
			},
			hasError: false,
		},
		{
			name: "idに27文字以上の文字列は指定できない",
			in: usecase.DeleteProjectInput{
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
				ctx: tctx,
				in:  usecase.DeleteProjectInput{ID: "project_01"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mysql.ResetProject(t, tc)
			})

			if err := project.DeleteProject(tt.args.ctx, tt.args.in); err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

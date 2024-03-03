package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_HasProject(t *testing.T) {
	type args struct {
		p Project
	}
	tests := []struct {
		name string
		u    User
		args args
		want bool
	}{
		{
			name: "ユーザがプロジェクトを所有している",
			u:    User{ID: "01DXF6DT000000000000000000"},
			args: args{
				p: Project{UserID: "01DXF6DT000000000000000000"},
			},
			want: true,
		},
		{
			name: "ユーザがプロジェクトを所有していない",
			u:    User{ID: "01DXF6DT000000000000000000"},
			args: args{
				p: Project{UserID: "01DXF6DT000000000000000001"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.u.HasProject(tt.args.p)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUser_HasStep(t *testing.T) {
	type args struct {
		s Step
	}
	tests := []struct {
		name string
		u    User
		args args
		want bool
	}{
		{
			name: "ユーザがステップを所有している",
			u:    User{ID: "01DXF6DT000000000000000000"},
			args: args{
				s: Step{UserID: "01DXF6DT000000000000000000"},
			},
			want: true,
		},
		{
			name: "ユーザがステップを所有していない",
			u:    User{ID: "01DXF6DT000000000000000000"},
			args: args{
				s: Step{UserID: "01DXF6DT000000000000000001"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.u.HasStep(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUser_HasTag(t *testing.T) {
	type args struct {
		t Tag
	}
	tests := []struct {
		name string
		u    User
		args args
		want bool
	}{
		{
			name: "ユーザがタグを所有している",
			u:    User{ID: "01DXF6DT000000000000000000"},
			args: args{
				t: Tag{UserID: "01DXF6DT000000000000000000"},
			},
			want: true,
		},
		{
			name: "ユーザがタグを所有していない",
			u:    User{ID: "01DXF6DT000000000000000000"},
			args: args{
				t: Tag{UserID: "01DXF6DT000000000000000001"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.u.HasTag(tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUser_HasTask(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name string
		u    User
		args args
		want bool
	}{
		{
			name: "ユーザがタスクを所有している",
			u:    User{ID: "01DXF6DT000000000000000000"},
			args: args{
				t: Task{UserID: "01DXF6DT000000000000000000"},
			},
			want: true,
		},
		{
			name: "ユーザがタスクを所有していない",
			u:    User{ID: "01DXF6DT000000000000000000"},
			args: args{
				t: Task{UserID: "01DXF6DT000000000000000001"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.u.HasTask(tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

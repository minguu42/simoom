package model

import "testing"

func TestProject_ContainsTask(t *testing.T) {
	type args struct {
		task Task
	}
	tests := []struct {
		name string
		p    Project
		args args
		want bool
	}{
		{
			name: "タスクがプロジェクトに含まれている",
			p:    Project{ID: "01DXF6DT000000000000000000"},
			args: args{
				task: Task{ProjectID: "01DXF6DT000000000000000000"},
			},
			want: true,
		},
		{
			name: "タスクがプロジェクトに含まれていない",
			p:    Project{ID: "01DXF6DT000000000000000000"},
			args: args{
				task: Task{ProjectID: "01DXF6DT000000000000000001"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ContainsTask(tt.args.task); tt.want != got {
				t.Errorf("p.ContainsTask want %t, but got %t", tt.want, got)
			}
		})
	}
}

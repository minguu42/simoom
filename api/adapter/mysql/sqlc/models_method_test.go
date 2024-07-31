package sqlc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSteps_StepsByTaskID(t *testing.T) {
	tests := []struct {
		name string
		ss   Steps
		want map[string]Steps
	}{
		{
			name: "TaskIDをキーとしたマップを生成する",
			ss: []Step{
				{
					ID:     "step01",
					TaskID: "task01",
					Name:   "テストステップ1",
				},
				{
					ID:     "step02",
					TaskID: "task02",
					Name:   "テストステップ2",
				},
				{
					ID:     "step03",
					TaskID: "task01",
					Name:   "テストステップ3",
				},
			},
			want: map[string]Steps{
				"task01": []Step{
					{
						ID:     "step01",
						TaskID: "task01",
						Name:   "テストステップ1",
					},
					{
						ID:     "step03",
						TaskID: "task01",
						Name:   "テストステップ3",
					},
				},
				"task02": []Step{
					{
						ID:     "step02",
						TaskID: "task02",
						Name:   "テストステップ2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ss.StepsByTaskID()
			assert.Equal(t, tt.want, got)
		})
	}
}

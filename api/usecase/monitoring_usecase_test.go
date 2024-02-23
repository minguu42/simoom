package usecase_test

import (
	"testing"

	"github.com/minguu42/simoom/api/usecase"
	"github.com/stretchr/testify/assert"
)

var monitoring = usecase.Monitoring{}

func TestMonitoringUsecase_CheckHealth(t *testing.T) {
	tests := []struct {
		name string
		want usecase.CheckHealthOutput
	}{
		{
			name: "アプリケーションのrevisionを返す",
			want: usecase.CheckHealthOutput{
				Revision: "xxxxxxx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := monitoring.CheckHealth()
			assert.Equal(t, tt.want, got)
		})
	}
}

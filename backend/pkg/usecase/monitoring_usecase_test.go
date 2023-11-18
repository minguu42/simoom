package usecase_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/minguu42/simoom/backend/pkg/usecase"
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
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("monitoring.CheckHealth mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

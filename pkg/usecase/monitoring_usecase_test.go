package usecase

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var testMonitoring = MonitoringUsecase{}

func TestMonitoringUsecase_CheckHealth(t *testing.T) {
	tests := []struct {
		name string
		want CheckHealthOutput
	}{
		{
			name: "アプリケーションのrevisionを返す",
			want: CheckHealthOutput{
				Revision: "xxxxxxx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testMonitoring.CheckHealth()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("testMonitoring.CheckHealth mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

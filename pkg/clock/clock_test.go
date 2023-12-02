package clock_test

import (
	"context"
	"testing"
	"time"

	"github.com/minguu42/simoom/pkg/clock"
	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2020-01-01 00:00:00Z",
			args: args{ctx: clock.NowWithValue(context.Background(),
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clock.Now(tt.args.ctx)
			assert.Equal(t, tt.want, got)
		})
	}
}

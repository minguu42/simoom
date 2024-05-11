package clock

import (
	"context"
	"time"
)

type timeKey struct{}

// Now は ctx に時刻を含む場合はその値を返し、含まない場合は現在時刻を返す
func Now(ctx context.Context) time.Time {
	if t, ok := ctx.Value(timeKey{}).(time.Time); ok {
		return t
	}
	return time.Now()
}

// NowWithValue は ctx に時刻 t を含めた新しいコンテキストを返す
func NowWithValue(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, timeKey{}, t)
}

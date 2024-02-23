package pointers_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/minguu42/simoom/lib/go/pointers"
	"github.com/stretchr/testify/assert"
)

func TestRef(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name         string
		args         args
		wantTypeKind reflect.Kind
	}{
		{
			name:         "string型の値のポインタを返す",
			args:         args{v: "Hello, 世界!"},
			wantTypeKind: reflect.Ptr,
		},
		{
			name:         "int型の値のポインタを返す",
			args:         args{v: 10},
			wantTypeKind: reflect.Ptr,
		},
		{
			name:         "time.Time型の値のポインタを返す",
			args:         args{v: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
			wantTypeKind: reflect.Ptr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTypeKind := reflect.TypeOf(pointers.Ref(tt.args.v)).Kind()
			assert.Equal(t, tt.wantTypeKind, gotTypeKind)
		})
	}
}

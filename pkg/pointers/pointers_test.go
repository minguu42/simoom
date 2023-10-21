package pointers

import (
	"reflect"
	"testing"
	"time"
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
			args:         args{v: time.Now()},
			wantTypeKind: reflect.Ptr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType := reflect.TypeOf(Ref(tt.args.v))
			if tt.wantTypeKind != gotType.Kind() {
				t.Errorf("return type of the ref function want %s, but %s", tt.wantTypeKind, gotType.Kind())
			}
		})
	}
}

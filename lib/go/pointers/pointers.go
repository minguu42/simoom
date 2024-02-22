// Package pointers はポインタに関するユーティリティ関数を含む
package pointers

// Ref 関数は v のポインタを返す
func Ref[T any](v T) *T {
	return &v
}

package pointers

// Ref 関数は v のポインタを返す
func Ref[T any](v T) *T {
	return &v
}

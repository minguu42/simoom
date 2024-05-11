package pointers

// Ref は v のポインタを返す
func Ref[T any](v T) *T {
	return &v
}

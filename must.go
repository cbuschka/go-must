package must

// Resolve returns t and panics if err != nil.
func Resolve[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

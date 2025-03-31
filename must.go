package must

// Be gives t and panics if err != nil.
func Be[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

// Resolve executes f, returns t and panics if err != nil.
func Resolve[T any](f func() (T, error)) T {
	return Be(f())
}

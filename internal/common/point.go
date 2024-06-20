package common

func Point[T any](t T) *T {
	return &t
}

package gopromax

func IndexOf[T comparable](s []T, e T) int {
	for i := range s {
		if s[i] == e {
			return i
		}
	}
	return -1
}

func LastIndexOf[T comparable](s []T, e T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == e {
			return i
		}
	}
	return -1
}

func Find[T any](s []T, f func(T) bool) *T {
	for i := range s {
		if f(s[i]) {
			return &s[i]
		}
	}
	return nil
}
func FindIndexOf[T any](s []T, f func(T) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

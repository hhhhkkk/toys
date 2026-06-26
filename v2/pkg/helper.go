package pkg

func MapSlice[T any, R any](data []T, fn func(T) R) []R {
	ns := make([]R, len(data))
	for k, v := range data {
		ns[k] = fn(v)
	}
	return ns
}

func FilterSlice[T any](data []T, fn func(T) bool) []T {
	ns := make([]T, len(data))
	for _, v := range data {
		if fn(v) {
			ns = append(ns, v)
		}
	}
	return ns
}

func ReduceSlice[T any, R any](data []T, init R, fn func(d R, c T) R) R {
	d := init
	for _, v := range data {
		d = fn(d, v)
	}
	return d
}

func If[T any](condition bool, ifT T, ifF T) T {
	if condition {
		return ifT
	}
	return ifF
}

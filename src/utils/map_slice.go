package utils

func MapSlice[T any, R any](value []T, cb func(T) R) []R {
	arr := []R{}

	for _, v := range value {
		arr = append(arr, cb(v))
	}

	return arr
}

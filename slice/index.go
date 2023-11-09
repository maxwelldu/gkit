package slice

func IndexFunc[T any](src []T, match matchFunc[T]) int {
	for i, v := range src {
		if match(v) {
			return i
		}
	}
	return -1
}

func Index[T comparable](src []T, dst T) int {
	return IndexFunc[T](src, func(src T) bool {
		return src == dst
	})
}

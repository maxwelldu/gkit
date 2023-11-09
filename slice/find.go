package slice

func Find[T any](src []T, match matchFunc[T]) (T, bool) {
	for _, v := range src {
		if match(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

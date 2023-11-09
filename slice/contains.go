package slice

func ContainsFunc[T any](src []T, equal func(src T) bool) bool {
	for _, v := range src {
		if equal(v) {
			return true
		}
	}
	return false
}

func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, func(src T) bool {
		return src == dst
	})
}

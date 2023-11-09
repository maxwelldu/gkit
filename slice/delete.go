package slice

import "github.com/maxwelldu/gkit/internal/slice"

func Delete[T any](src []T, index int) ([]T, error) {
	res, _, err := slice.Delete(src, index)
	return res, err
}

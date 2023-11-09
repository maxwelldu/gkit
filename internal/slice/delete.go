package slice

import "github.com/maxwelldu/gkit/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}

	res := src[index]
	dst := make([]T, len(src)-1, cap(src))
	copy(dst, src[:index])
	copy(dst[index:], src[index+1:])
	return dst, res, nil
}

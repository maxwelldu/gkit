package errs

import (
	"fmt"
)

var ErrEmptySlice = fmt.Errorf("gkit: empty slice")

func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("gkit: index out of range, length %d, index %d", length, index)
}

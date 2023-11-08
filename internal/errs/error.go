package errs

import "fmt"

func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("gkit: index out of range, length %d, index %d", length, index)
}

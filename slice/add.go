package slice

import "github.com/maxwelldu/gkit/internal/slice"

// Add element to slice at index, index should be in [0, len(src))
func Add[T any](src []T, element T, index int) ([]T, error) {
	res, err := slice.Add[T](src, element, index)
	return res, err
}

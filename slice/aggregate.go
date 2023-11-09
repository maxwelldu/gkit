package slice

import (
	"github.com/maxwelldu/gkit"
	"github.com/maxwelldu/gkit/internal/errs"
)

// Max returns the maximum value. Be careful about accuracy issues when using float32 or float64.
func Max[T gkit.RealNumber](elements []T) (maxValue T, err error) {
	if len(elements) == 0 {
		return maxValue, errs.ErrEmptySlice
	}

	maxValue = elements[0]
	for i := 1; i < len(elements); i++ {
		if elements[i] > maxValue {
			maxValue = elements[i]
		}
	}
	return maxValue, nil
}

// Min returns the minimum value. Be careful about accuracy issues when using float32 or float64.
func Min[T gkit.RealNumber](elements []T) (minValue T, err error) {
	if len(elements) == 0 {
		return minValue, errs.ErrEmptySlice
	}

	minValue = elements[0]
	for i := 1; i < len(elements); i++ {
		if elements[i] < minValue {
			minValue = elements[i]
		}
	}
	return minValue, nil
}

// Sum returns the sum of all elements. Be careful about accuracy issues when using float32 or float64.
func Sum[T gkit.Number](elements ...T) T {
	var sum T
	for _, e := range elements {
		sum += e
	}
	return sum
}

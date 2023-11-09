package slice

import (
	"testing"

	"github.com/maxwelldu/gkit/internal/errs"

	"github.com/maxwelldu/gkit"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		wantMaxValue int
		wantErr      error
	}{
		{
			name:    "nil",
			wantErr: errs.ErrEmptySlice,
		},
		{
			name:    "empty",
			input:   []int{},
			wantErr: errs.ErrEmptySlice,
		},
		{
			name:         "value",
			input:        []int{1},
			wantMaxValue: 1,
		},
		{
			name:         "values",
			input:        []int{1, 2, 3},
			wantMaxValue: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMaxValue, err := Max(tt.input)
			assert.Equalf(t, tt.wantErr, err, "Max(%v)", tt.input)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantMaxValue, gotMaxValue)
		})
	}

	testMaxTypes[int](t)
	testMaxTypes[int8](t)
	testMaxTypes[int16](t)
	testMaxTypes[int32](t)
	testMaxTypes[int64](t)
	testMaxTypes[uint](t)
	testMaxTypes[uint8](t)
	testMaxTypes[uint16](t)
	testMaxTypes[uint32](t)
	testMaxTypes[uint64](t)
	testMaxTypes[float32](t)
	testMaxTypes[float64](t)
}

func TestMin(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		wantMinValue int
		wantErr      error
	}{
		{
			name:    "nil",
			wantErr: errs.ErrEmptySlice,
		},
		{
			name:    "empty",
			input:   []int{},
			wantErr: errs.ErrEmptySlice,
		},
		{
			name:         "value",
			input:        []int{3},
			wantMinValue: 3,
		},
		{
			name:         "values",
			input:        []int{3, 1, 2},
			wantMinValue: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMinValue, err := Min(tt.input)
			assert.Equalf(t, tt.wantErr, err, "Min(%v)", tt.input)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantMinValue, gotMinValue)
		})
	}

	testMinTypes[int](t)
	testMinTypes[int8](t)
	testMinTypes[int16](t)
	testMinTypes[int32](t)
	testMinTypes[int64](t)
	testMinTypes[uint](t)
	testMinTypes[uint8](t)
	testMinTypes[uint16](t)
	testMinTypes[uint32](t)
	testMinTypes[uint64](t)
	testMinTypes[float32](t)
	testMinTypes[float64](t)
}

func TestSum(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		wantSumValue int
		wantErr      error
	}{
		{
			name: "nil",
		},
		{
			name:  "empty",
			input: []int{},
		},
		{
			name:         "value",
			input:        []int{1},
			wantSumValue: 1,
		},
		{
			name:         "values",
			input:        []int{1, 2, 3},
			wantSumValue: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSumValue := Sum(tt.input...)
			assert.Equal(t, tt.wantSumValue, gotSumValue)
		})
	}

	testSumTypes[int](t)
	testSumTypes[int8](t)
	testSumTypes[int16](t)
	testSumTypes[int32](t)
	testSumTypes[int64](t)
	testSumTypes[uint](t)
	testSumTypes[uint8](t)
	testSumTypes[uint16](t)
	testSumTypes[uint32](t)
	testSumTypes[uint64](t)
	testSumTypes[float32](t)
	testSumTypes[float64](t)
}

// testMaxTypes[int](t) just for testing all types that meet the constraints of the Max method
func testMaxTypes[T gkit.RealNumber](t *testing.T) {
	res, _ := Max[T]([]T{1, 2, 3})
	assert.Equal(t, T(3), res)
}

// testMinTypes[int](t) just for testing all types that meet the constraints of the Min method
func testMinTypes[T gkit.RealNumber](t *testing.T) {
	res, _ := Min[T]([]T{1, 2, 3})
	assert.Equal(t, T(1), res)
}

// testSumTypes[int](t) just for testing all types that meet the constraints of the Sum method
func testSumTypes[T gkit.RealNumber](t *testing.T) {
	res := Sum[T]([]T{1, 2, 3}...)
	assert.Equal(t, T(6), res)
}

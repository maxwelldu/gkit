package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  int
		want bool
	}{
		{
			name: "contains",
			src:  []int{1, 2, 3},
			dst:  2,
			want: true,
		},
		{
			name: "not contains",
			src:  []int{1, 2, 3},
			dst:  4,
			want: false,
		},
		{
			name: "empty",
			src:  []int{},
			dst:  4,
			want: false,
		},
		{
			name: "nil",
			dst:  4,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Contains[int](tt.src, tt.dst), "Contains(%v, %v)", tt.src, tt.dst)
		})
	}
}

func TestContainsFunc(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  int
		want bool
	}{
		{
			name: "contains",
			src:  []int{1, 2, 3},
			dst:  2,
			want: true,
		},
		{
			name: "not contains",
			src:  []int{1, 2, 3},
			dst:  4,
			want: false,
		},
		{
			name: "empty",
			src:  []int{},
			dst:  4,
			want: false,
		},
		{
			name: "nil",
			dst:  4,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ContainsFunc[int](tt.src, func(src int) bool {
				return src == tt.dst
			}))
		})
	}
}

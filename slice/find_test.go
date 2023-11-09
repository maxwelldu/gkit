package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		match   matchFunc[int]
		want    int
		founded bool
	}{
		{
			name:  "found",
			input: []int{1, 2, 3},
			match: func(v int) bool {
				return v == 2
			},
			want:    2,
			founded: true,
		},
		{
			name:  "not found",
			input: []int{1, 2, 3},
			match: func(v int) bool {
				return v == 4
			},
			want:    0,
			founded: false,
		},
		{
			name:  "empty",
			input: []int{},
			match: func(v int) bool {
				return v == 4
			},
			want:    0,
			founded: false,
		},
		{
			name: "nil",
			match: func(v int) bool {
				return v == 4
			},
			want:    0,
			founded: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, founded := Find(tt.input, tt.match)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.founded, founded)
		})
	}
}

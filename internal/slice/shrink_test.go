package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShrink(t *testing.T) {
	tests := []struct {
		name        string
		originCap   int
		enqueueLoop int
		expectCap   int
	}{
		{
			name:        "less than 64",
			originCap:   32,
			enqueueLoop: 6,
			expectCap:   32,
		},
		{
			name:        "less than 2048, less than 1/4",
			originCap:   1000,
			enqueueLoop: 20,
			expectCap:   500,
		},
		{
			name:        "less than 2048, more than 1/4",
			originCap:   1000,
			enqueueLoop: 400,
			expectCap:   1000,
		},
		{
			name:        "more than 2048, less than 1/2",
			originCap:   3000,
			enqueueLoop: 60,
			expectCap:   1875,
		},
		{
			name:        "more than 2048, more than 1/2",
			originCap:   3000,
			enqueueLoop: 2000,
			expectCap:   3000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := make([]int, 0, tt.originCap)
			for i := 0; i < tt.enqueueLoop; i++ {
				l = append(l, i)
			}
			l = Shrink(l)
			assert.Equal(t, tt.expectCap, cap(l))
		})
	}
}

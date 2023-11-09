package slice

import (
	"testing"

	"github.com/maxwelldu/gkit/internal/errs"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		index   int
		want    []int
		wantErr error
	}{
		{
			name:  "index 0",
			slice: []int{1, 2, 3},
			index: 0,
			want:  []int{2, 3},
		},
		{
			name:    "index -1",
			slice:   []int{1, 2, 3},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(3, -1),
		},
		{
			name:    "index 3",
			slice:   []int{1, 2, 3},
			index:   3,
			wantErr: errs.NewErrIndexOutOfRange(3, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Delete(tt.slice, tt.index)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

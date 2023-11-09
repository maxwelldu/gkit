package slice

import (
	"reflect"
	"testing"

	"github.com/maxwelldu/gkit/internal/errs"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		src     []int
		element int
		index   int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{
			name: "index 0",
			args: args{
				src:     []int{123, 100},
				element: 233,
				index:   0,
			},
			want: []int{233, 123, 100},
		},
		{
			name: "index -1",
			args: args{
				src:   []int{123, 100},
				index: -1,
			},
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.src, tt.args.element, tt.args.index)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

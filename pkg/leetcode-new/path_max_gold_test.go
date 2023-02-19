package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getMaximumGold(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Head test - 3*5 matrix",
			args: args{
				grid: [][]int{
					{1, 0, 7},
					{2, 0, 6},
					{3, 4, 5},
					{0, 3, 0},
					{9, 0, 20},
				},
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getMaximumGold(tt.args.grid), "getMaximumGold(%v)", tt.args.grid)
		})
	}
}

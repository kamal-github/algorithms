package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_spiralMatrixIII(t *testing.T) {
	type args struct {
		rows   int
		cols   int
		rStart int
		cStart int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "first test case",
			args: args{
				rows:   1,
				cols:   4,
				rStart: 0,
				cStart: 0,
			},
			want: [][]int{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := spiralMatrixIII(tt.args.rows, tt.args.cols, tt.args.rStart, tt.args.cStart)

			assert.EqualValues(t, tt.want, actual)
		})
	}
}

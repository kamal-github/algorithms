package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{2, 3, 4, 4, 3, 10, 2},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, singleNumber(tt.args.nums), "singleNumber(%v)", tt.args.nums)
		})
	}
}

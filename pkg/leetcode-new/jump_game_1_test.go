package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: true,
		},
		{
			name: "with 0 at the end",
			args: args{nums: []int{2, 0, 0}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canJump(tt.args.nums), "canJump(%v)", tt.args.nums)
		})
	}
}

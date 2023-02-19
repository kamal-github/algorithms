package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 1},
			want: 1,
		},
		{
			args: args{n: 2},
			want: 2,
		},
		{
			args: args{n: 3},
			want: 3,
		},
		{
			args: args{n: 5},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, climbStairs(tt.args.n), "climbStairs(%v)", tt.args.n)
		})
	}
}

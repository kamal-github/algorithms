package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{n: 27},
			want: true,
		},
		{
			args: args{n: -1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPowerOfThree(tt.args.n), "isPowerOfThree(%v)", tt.args.n)
		})
	}
}

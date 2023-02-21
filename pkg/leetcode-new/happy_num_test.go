package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "for single digit",
			args: args{
				n: 2,
			},
			want: false,
		},
		{
			name: "with 1111111",
			args: args{n: 1111111},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isHappy(tt.args.n), "isHappy(%v)", tt.args.n)
		})
	}
}

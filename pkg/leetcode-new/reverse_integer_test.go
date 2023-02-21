package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseInteger(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "123",
			args: args{
				x: 123,
			},
			want: 321,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reverseInteger(tt.args.x), "reverseInteger(%v)", tt.args.x)
		})
	}
}

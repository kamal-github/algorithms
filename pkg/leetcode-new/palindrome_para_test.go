package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{s: "race a car"},
			want: false,
		},
		{
			name: "",
			args: args{s: "0P"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPalindrome(tt.args.s), "isPalindrome(%v)", tt.args.s)
		})
	}
}

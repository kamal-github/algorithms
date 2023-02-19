package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with small string",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "with all same chars",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "au",
			args: args{
				s: "au",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLongestSubstring(tt.args.s), "lengthOfLongestSubstring(%v)", tt.args.s)
		})
	}
}

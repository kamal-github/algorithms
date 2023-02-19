package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "for 2 digits",
			args: args{digits: "23"},
			want: []string{
				"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
			},
		},
		{
			name: "for 0 digits",
			args: args{digits: ""},
			want: []string{},
		},
		{
			name: "for 1 digit",
			args: args{digits: "3"},
			want: []string{"d", "e", "f"},
		},
		{
			name: "for 3 digits",
			args: args{digits: "235"},
			want: []string{"adj", "adk", "adl", "aej", "aek", "ael", "afj", "afk", "afl", "bdj", "bdk", "bdl", "bej", "bek", "bel", "bfj", "bfk", "bfl", "cdj", "cdk", "cdl", "cej", "cek", "cel", "cfj", "cfk", "cfl"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, letterCombinations(tt.args.digits), "letterCombinations(%v)", tt.args.digits)
		})
	}
}

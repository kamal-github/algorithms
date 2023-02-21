package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with double letters",
			args: args{columnTitle: "ZY"},
			want: 701,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, titleToNumber(tt.args.columnTitle), "titleToNumber(%v)", tt.args.columnTitle)
		})
	}
}

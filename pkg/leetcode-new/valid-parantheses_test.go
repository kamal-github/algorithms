package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			want: false,
			args: args{s: "]"},
		},
		{
			name: "",
			want: false,
			args: args{s: "]]]]]"},
		},
		{
			name: "",
			want: false,
			args: args{s: "("},
		},
		{
			name: "",
			want: false,
			args: args{s: "((("},
		},
		{
			name: "",
			want: false,
			args: args{
				s: "(){}}{",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isValid(tt.args.s), "isValid(%v)", tt.args.s)
		})
	}
}

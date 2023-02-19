package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{path: "/home/..//sub/./"},
			want: "/sub",
		},
		{
			name: "2",
			args: args{path: "/../"},
			want: "/",
		},
		{
			name: "3",
			args: args{
				path: "/a/./b/../../c/",
			},
			want: "/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, simplifyPath(tt.args.path), "simplifyPath(%v)", tt.args.path)
		})
	}
}

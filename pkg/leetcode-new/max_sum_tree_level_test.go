package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxLevelSum(tt.args.root), "maxLevelSum(%v)", tt.args.root)
		})
	}
}

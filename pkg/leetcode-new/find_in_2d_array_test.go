package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindIn2DMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	assert.True(t, searchMatrix(matrix, 3))
}

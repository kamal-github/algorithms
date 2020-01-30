package leetcode_test

import (
	"testing"

	"github.com/kamal-github/algorithms/pkg/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestNumberOfBulbShines(t *testing.T) {
	assert.Equal(t, 3, leetcode.NumberOfBulbShines([]int{2, 1, 3, 5, 4}))
	assert.Equal(t, 2, leetcode.NumberOfBulbShines([]int{2, 3, 4, 1, 5}))
	assert.Equal(t, 3, leetcode.NumberOfBulbShines([]int{1, 3, 4, 2, 5}))
}

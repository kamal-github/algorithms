package leetcode_test

import (
	"testing"

	"github.com/kamal-github/algorithms/pkg/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.EqualValues(
		t,
		[]int{2, 3},
		leetcode.TwoSum([]int{3, 4, 11, 15}, 26),
	)

	assert.EqualValues(
		t,
		[]int{1, 2},
		leetcode.TwoSum([]int{3, 2, 4}, 6),
	)

	assert.EqualValues(
		t,
		[]int{0, 1},
		leetcode.TwoSum([]int{3, 3}, 6),
	)
	assert.EqualValues(
		t,
		[]int{1, 3},
		leetcode.TwoSum([]int{7, 2, 4, 1}, 3),
	)
}

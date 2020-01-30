package leetcode_test

import (
	"testing"

	"github.com/kamal-github/algorithms/pkg/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedInPlace(t *testing.T) {
	a := []int{1, 2, 0, 0}
	leetcode.MergeTwoSortedInPlace(a, 2, []int{1, 3}, 2)
	assert.EqualValues(
		t,
		[]int{1, 1, 2, 3},
		a,
	)

	b := []int{0, 1, 2, 4, 6, 7, 8, 0, 0, 0, 0}
	leetcode.MergeTwoSortedInPlace(b, 7, []int{1, 3, 3, 6}, 4)
	assert.EqualValues(
		t,
		[]int{0, 1, 1, 2, 3, 3, 4, 6, 6, 7, 8},
		b,
	)
}

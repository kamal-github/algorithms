package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArraySumDP(t *testing.T) {
	assert.Equal(t, 0, MaxSubArraySumDP([]int{0}))
	assert.Equal(t, 4, MaxSubArraySumDP([]int{4}))
	assert.Equal(t, 14, MaxSubArraySumDP([]int{4, 5, -1, 6}))
	assert.Equal(t, 18, MaxSubArraySumDP([]int{4, -5, -1, 6, 0, 2,10}))
	assert.Equal(t, 10, MaxSubArraySumDP([]int{1, 2, -3, -4, 2, 7, -2, 3}))
}

func TestMaxSubArraySum(t *testing.T) {
	assert.Equal(t, 0, MaxSubArraySum([]int{0}))
	assert.Equal(t, 4, MaxSubArraySum([]int{4}))
	assert.Equal(t, 14, MaxSubArraySum([]int{4, 5, -1, 6}))
	assert.Equal(t, 18, MaxSubArraySum([]int{4, -5, -1, 6, 0, 2,10}))
}


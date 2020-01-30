package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", LongestCommonPrefix([]string{"flower","flow","flight"}))
	assert.Equal(t, "flow", LongestCommonPrefix([]string{"flower","flow"}))
	assert.Equal(t, "", LongestCommonPrefix([]string{"flower","flow","flight", "kamal"}))
	assert.Equal(t, "", LongestCommonPrefix([]string{"dog","racecar","car"}))
	assert.Equal(t, "a", LongestCommonPrefix([]string{"aa", "a"}))
	assert.Equal(t, "abc", LongestCommonPrefix([]string{"abc"}))
}

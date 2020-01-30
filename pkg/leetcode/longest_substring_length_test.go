package leetcode_test

import (
	"testing"

	"github.com/kamal-github/algorithms/pkg/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestLongestSubstringLength(t *testing.T) {
	assert.Equal(t, 3, leetcode.LongestSubstringLength("pwwkew"))
	assert.Equal(t, 1, leetcode.LongestSubstringLength("bbbb"))
	assert.Equal(t, 0, leetcode.LongestSubstringLength(""))
	assert.Equal(t, 3, leetcode.LongestSubstringLength("abccabcb"))
	assert.Equal(t, 16, leetcode.LongestSubstringLength("abcaddhewrewolafjsdkfjewoifjdskjfsdlkfjeifjosdjflsdkjflsdkfjlsdkjfdlskfjlsdkfdskmclsnvdbjgurewyrewyrieuwweww"))
}

func BenchmarkLongestSubstringLength(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		leetcode.LongestSubstringLength("abcaddhewrewolafjsdkfjewoifjdskjfsdlkfjeifjosdjflsdkjflsdkfjlsdkjfdlskfjlsdkfdskmclsnvdbjgurewyrewyrieuwweww")
	}
}

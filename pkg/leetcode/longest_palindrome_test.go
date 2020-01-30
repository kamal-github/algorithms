package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "bbadabb", LongestPalindrome("abbadabbbadpwpw"))
	assert.Equal(t, "b", LongestPalindrome("b"))
	assert.Equal(t, "bb", LongestPalindrome("bb"))
	assert.Equal(t, "", LongestPalindrome(""))
}

func TestLongestPalindromeDP(t *testing.T) {
	assert.Equal(t, "dabbbad", LongestPalindromeDP("abbadabbbadpwpw"))
	assert.Equal(t, "b", LongestPalindrome("b"))
	assert.Equal(t, "bb", LongestPalindrome("bb"))
	assert.Equal(t, "", LongestPalindrome(""))
	assert.Equal(t, "a", LongestPalindrome("ac"))
}


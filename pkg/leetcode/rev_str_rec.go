package leetcode

import (
	"fmt"
)

func PrintReverse(str string) string {
	if len(str) == 0 {
		return ""
	}
	PrintReverse(str[1:])
	fmt.Print(string(str[0]))

	return ""
}

// TIMES
// TIME
// Logic - Swap the extreme two elements recursively.
func ReverseInPlace(s []byte) {
	if len(s) == 0 || len(s) == 1 {
		return
	}

	// this will solve for sub-problems (2nd and 2nd last element) recursively
	ReverseInPlace(s[1 : len(s)-1])
	swap(&s[0], &s[len(s)-1])

}

func swap(a, b *byte) {
	t := *b
	*b = *a
	*a = t
}

package main

import "sort"

// Idea:

// calculate the sum of all chars of p and substring of s of length of p.
// if the sum is equal only then you can compare p and substring of s.
// - 1) by sorting them and compare
// OR efficiently -
// - 2) by keeping the frequency of each char in p and substring of s and compare them
// in `equals` function. (NOT implemented)
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	asciiSumP := 0
	anagramsBegIndices := make([]int, 0)
	pRune := []rune(p)
	sort.SliceStable(pRune, func(i, j int) bool {
		return pRune[i] < pRune[j]
	})

	for _, char := range p {
		asciiSumP += int(char)
	}

	asciiSumSubS := 0
	for _, char := range s[0:len(p)] {
		asciiSumSubS += int(char)
	}

	for i := 0; i < len(s)-len(p)+1; i++ {
		if asciiSumSubS == asciiSumP {
			if equals(s[i:i+len(p)], pRune) {
				anagramsBegIndices = append(anagramsBegIndices, i)
			}
			continue
		}

		if i+len(p) >= len(s) {
			break
		}

		asciiSumSubS = asciiSumSubS - int(s[i]) + int(s[i+len(p)])
	}

	return anagramsBegIndices
}

func equals(s1 string, sortedP []rune) bool {
	s1Runes := []rune(s1)

	sort.SliceStable(s1Runes, func(i, j int) bool {
		return s1Runes[i] < s1Runes[j]
	})

	return string(s1Runes) == string(sortedP)
}

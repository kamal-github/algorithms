package main

import (
	"strings"
)

func letterCombinations(digits string) []string {
	digitsLength := len(digits)
	m := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	if digitsLength == 0 {
		return []string{}
	}
	if digitsLength == 1 {
		return m[string(digits[0])]
	}

	res := m[digits[0:1]]
	for i := 1; i < digitsLength; i++ {
		res = multiply(res, m[digits[i:i+1]])
	}

	return res
}

func multiply(first, second []string) []string {
	var res []string
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			var strBuilder strings.Builder
			strBuilder.WriteString(first[i])
			strBuilder.WriteString(second[j])
			res = append(res, strBuilder.String())
		}
	}

	return res
}

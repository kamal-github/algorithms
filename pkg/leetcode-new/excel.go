package main

import (
	"math"
)

func titleToNumber(columnTitle string) int {
	if columnTitle == "" {
		return 0
	}

	m := make(map[string]int)
	j := 0
	for i := 65; i < 91; i++ {
		j++
		m[string(rune(i))] = j
	}

	if len(columnTitle) == 1 {
		return m[string(columnTitle[0])]
	}

	colNum := 0
	n := len(columnTitle)
	for i, ch := range columnTitle {
		// e.g ZB = 26^1 (the power increase as per index from back as Z is at 1 from back so 26^1) * 26 (this is for char Z) + 2 (this is for char B)
		colNum = colNum + int(int(math.Pow(26, float64(n-i-1)))*m[string(ch)])
	}

	return colNum
}

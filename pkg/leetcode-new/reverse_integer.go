package main

import "math"

// https://leetcode.com/problems/reverse-integer/description/
func reverseInteger(x int) int {
	var rev int
	digits := getDigits(x)
	lenDigits := len(digits)
	for i, d := range digits {
		pow := lenDigits - i - 1

		// 3, 2, 1
		// 3*100+2*10+1*1
		rev += d * int(math.Pow(10, float64(pow)))
	}

	return rev
}

func getDigits(x int) []int {
	var digits []int
	for x != 0 {
		r := x % 10
		digits = append(digits, r)
		x = x / 10
	}
	return digits
}

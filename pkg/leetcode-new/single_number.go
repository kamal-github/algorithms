package main

// Hint : XOR of two equal numbers gives 0. Fuck you!
func singleNumber(nums []int) int {
	xorNum := 0

	for _, n := range nums {
		xorNum ^= n
	}

	return xorNum
}

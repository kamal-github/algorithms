package main

func climbStairs(n int) int {
	var steps int
	cache := make(map[int]int)
	climbRec(n, &steps, cache)

	return steps
}

// n = 2
// 1,1,1,1,1
func climbRec(n int, ways *int, cache map[int]int) {
	if n == 0 {
		*ways++
		return
	}
	if n < 0 {
		return
	}

	climbRec(n-1, ways, cache)
	climbRec(n-2, ways, cache)
	cache[n] = *ways
}

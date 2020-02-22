package dailycoding

/*
There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time.
Given N, write a function that returns the number of unique ways you can climb the staircase.
The order of the steps matters.
For example, if N is 4, then there are 5 unique ways:

1, 1, 1, 1
2, 1, 1
1, 2, 1
1, 1, 2
2, 2

What if, instead of being able to climb 1 or 2 steps at a time, you could climb
any number from a set of positive integers X? For example, if X = {1, 3, 5},
you could climb 1, 3, or 5 steps at a time.
*/
func UniqueStaircase(n int, steps []int) int {
	//return uniqueStaircase(n, steps)
	return uniqueStaircaseDP(n, steps, make(map[int]int))
}

func uniqueStaircaseDP(n int, steps []int, dp map[int]int) int {
	if sol, ok := dp[n]; ok {
		return sol
	}

	if n == 0 {
		return 1
	}

	var count int
	for _, s := range steps {
		if n-s >= 0 {
			count = count + uniqueStaircaseDP(n-s, steps, dp)
		}
	}

	dp[n] = count
	return count
}

func uniqueStaircase(n int, steps []int) int {
	if n == 0 {
		return 1
	}

	var count int
	for _, s := range steps {
		if n-s >= 0 {
			count = count + uniqueStaircase(n-s, steps)
		}
	}

	return count
}

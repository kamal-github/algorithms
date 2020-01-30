package leetcode

// Calculate Fibonacci number for nth index
// Brute force
func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Use Memoization
func FibonacciEff(n int) int {
	return fibCalculate(make(map[int]int), n)
}

func fibCalculate(mem map[int]int, n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if m, ok := mem[n]; ok {
		return m
	}

	mem[n] = fibCalculate(mem, n-1) + fibCalculate(mem, n-2)
	return mem[n]
}

// Using Fibonacci impl starting with 1 instead of 0
// https://leetcode.com/explore/learn/card/recursion-i/255/recursion-memoization/1662/
func ClimbStairs(n int) int {
	return fibCalculateStartWith1(make(map[int]int), n)
}
func fibCalculateStartWith1(mem map[int]int, n int) int {
	if n == 2 || n == 1 {
		return n
	}
	if m, ok := mem[n]; ok {
		return m
	}

	mem[n] = fibCalculateStartWith1(mem, n-1) + fibCalculateStartWith1(mem, n-2)
	return mem[n]
}

// Brute Force approach
func ClimbStairsBF(n int) int {
	return cal(0, n)
}

func cal(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return cal(i+1, n ) + cal(i+2,n)
}

//1,2,3,5,8,13...
// With Optimal Substructure property
func ClimbStairsDP(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	if n > 1 {
		dp[2] = 2
	}

	for i:=3; i<=n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}


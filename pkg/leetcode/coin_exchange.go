package leetcode

import (
	"math"

	"github.com/kamal-github/algorithms/pkg/leetcode/common"
)

/*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.

							Sample example recursion tree

					                    f(11)
						/                 |				\
					/-1                   | -2             \ -5
				f(10)                    f(9)              f(6)
			/-1   /-2  \-5       	/-1   |-2   \-5
		f(9)     f(8)     f(5)    f(8)   f(7)   f(4)        .....
             ...
*/

func CoinExchange(coins []int, amount int) int {
	return coinChange(coins, amount, make(map[int]int))
}

// Though process =
// We try all the possible combinations to find minimum coins to make amount
// For that, for each coin we will try with all the coins and reduce the amount by that selected coin
// Optimal sub problem is for each coin we are repeating the same solution(call `coinChange()`) for each coin until we get to reach to 0
func coinChange(coins []int, amount int, dp map[int]int) int {
	if cm, ok := dp[amount]; ok {
		return cm
	}

	minCnt := math.MaxInt64
	var currMinCnt int

	if amount == 0 {
		return 0
	}

	for _, c := range coins {
		if amount-c >= 0 {
			currMinCnt = coinChange(coins, amount-c, dp)
			minCnt = common.Smaller(currMinCnt, minCnt)
		} else {
			return math.MaxInt64
		}
	}

	dp[amount] = 1 + (minCnt)

	return dp[amount]
}

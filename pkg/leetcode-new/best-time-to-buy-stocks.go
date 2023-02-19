package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	localMin := prices[0]
	localMax := prices[0]
	maxProfit := localMax - localMin

	for _, p := range prices[1:] {
		if p > localMax {
			localMax = p
		}
		if p < localMin {
			localMin = p
			localMax = p
		}

		localProfit := localMax - localMin
		if maxProfit < localProfit {
			maxProfit = localProfit
		}
	}

	return maxProfit
}

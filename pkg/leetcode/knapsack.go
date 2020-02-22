package leetcode

import "github.com/kamal-github/algorithms/pkg/leetcode/common"

/*
items = {(w:2, v:6), (w:2, v:10), (w:3, v:12)}
max weight = 5
knapsack(items, max weight) = 22
				     22

*/

func knapsack(items []Item, i int, capacity int) int {
	if capacity <= 0 || i == len(items) {
		return 0
	}

	if capacity-items[i].weight < 0 {
		return knapsack(items, i+1, capacity)

	}

	return common.Bigger(
		items[i].value+knapsack(items, i+1, capacity-items[i].weight),
		knapsack(items, i+1, capacity),
	)
}

func knapsackDP(items []Item, i int, capacity int, cache map[int]map[int]int) int {
	if i == len(items) {
		return 0
	}

	if val, ok := cache[i][capacity]; ok {
		return val
	}

	if capacity-items[i].weight < 0 {
		return knapsack(items, i+1, capacity)
	}

	maxProfit := common.Bigger(
		items[i].value+knapsack(items, i+1, capacity-items[i].weight),
		knapsack(items, i+1, capacity),
	)

	if _, ok := cache[i]; !ok {
		cache[i] = map[int]int{
			capacity: maxProfit,
		}
	} else {
		cache[i][capacity] = maxProfit
	}

	return maxProfit
}

type Item struct {
	weight, value int
}

func KnapsackWithArrayInput(w []int, values []int, maxW int) int {
	items := make([]Item, len(w))
	for i := 0; i < len(w); i++ {
		items[i].weight = w[i]
		items[i].value = values[i]
	}

	return KnapsackDP(items, maxW)
}

func KnapsackDP(items []Item, maxW int) int {
	return knapsackDP(items, 0, maxW, make(map[int]map[int]int))
	//return knapsack(items, 0, maxW)
}

//func knapsack(items []Item, maxW int, cache map[int]int) int {
//	if mw, ok := cache[maxW]; ok {
//		return mw
//	}
//
//	maxV := 0
//
//	if maxW == 0 {
//		return 0
//	}
//
//	for _, item := range items {
//		if maxW-item.weight >= 0 {
//			currMaxV := item.value + knapsack(items, maxW-item.weight, cache)
//			maxV = common.Bigger(currMaxV, maxV)
//		}
//	}
//
//	cache[maxW] = maxV
//
//	return maxV
//}

package leetcode

import "github.com/kamal-github/algorithms/pkg/leetcode/common"

/*
items = {(w:2, v:6), (w:2, v:10), (w:3, v:12)}
max weight = 5
knapsack(items, max weight) = 22

                            22
                   /6	     |       \
				20           20      19
          /6   /10   \12
		18    18        17
			...
	/ / |
    end
*/

type Item struct {
	weight, value int
}

func KnapsackDP(items []Item, maxW int) int {
	return knapsack(items, maxW, make(map[int]int))
}

func knapsack(items []Item, maxW int, cache map[int]int) int {
	if mw, ok := cache[maxW]; ok {
		return mw
	}

	maxV := 0

	if maxW == 0 {
		return 0
	}

	for _, item := range items {
		if maxW-item.weight >= 0 {
			currMaxV := item.value + knapsack(items, maxW-item.weight, cache)
			maxV = common.Bigger(currMaxV, maxV)
		}
	}

	cache[maxW] = maxV

	return maxV
}

package main

import (
	"reflect"
	"sort"
)

type Tracker struct {
	numsI, numsJ, i, j int
}

// O(n^2)
func threeSum(nums []int) [][]int {
	sols := make([][]int, 0)
	twoSumsTracker := make(map[int][]Tracker)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			twoSum := nums[i] + nums[j]
			twoSumsTracker[twoSum] = append(
				twoSumsTracker[twoSum],
				Tracker{numsI: nums[i], numsJ: nums[j], i: i, j: j},
			)
		}
	}

	for k, thirdNum := range nums {
		var trackers []Tracker
		var ok bool
		if trackers, ok = twoSumsTracker[-thirdNum]; !ok {
			continue
		}

		for _, t := range trackers {
			new3Sum := []int{t.numsI, t.numsJ, thirdNum}
			if ok := filter(t, k); ok && !contains(sols, new3Sum) {
				sols = append(sols, new3Sum)
			}
		}
	}

	return sols
}

func filter(t Tracker, k int) bool {
	return t.i != k && t.j != k
}

func contains(arrOfArr [][]int, arr []int) bool {
	sort.Ints(arr)
	for _, aa := range arrOfArr {
		sort.Ints(aa)
		if reflect.DeepEqual(aa, arr) {
			return true
		}
	}

	return false
}

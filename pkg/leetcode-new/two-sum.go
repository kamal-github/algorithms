package main

func twoSum(nums []int, target int) []int {
	valueToIdx := make(map[int][]int, 0)

	for i, v := range nums {
		valueToIdx[v] = append(valueToIdx[v], i)
	}

	for i := 0; i < len(nums); i++ {
		indexes := valueToIdx[target-nums[i]]

		if len(indexes) == 1 && i != indexes[0] {
			return []int{i, indexes[0]}
		}

		if len(indexes) > 1 {
			return []int{i, indexes[1]}
		}
	}

	return []int{}
}

package leetcode

func TwoSum(nums []int, target int) []int {
	cacheIndexes := make(map[int][]int)
	for i, n := range nums {
		cacheIndexes[n] = append(cacheIndexes[n], i)
	}
	
	for i, n := range nums {
		indexes, ok := cacheIndexes[target - n]
		if ok {
			if len(indexes) > 1 {
				return []int{i, indexes[1]}
			} else {
				if indexes[0] != i {
					return []int{i,indexes[0]}
				}
			}
			
		}
	}
	return []int{}
}

package main

func canJump(nums []int) bool {
	return checkRec(nums, 0)
}

func checkRec(nums []int, i int) bool {
	if i > len(nums)-1 {
		return false
	}
	if i == len(nums)-1 {
		return true
	}
	var val bool
	for j := 1; j <= nums[i]; j++ {
		if checkRec(nums, i+j) {
			return true
		}
	}

	return val
}

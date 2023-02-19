package main

func jump(nums []int) int {
	minJumps := len(nums)
	minJumpsRec(nums, 0, &minJumps)
	return minJumps
}

// [2,3,1,1,4]
func minJumpsRec(nums []int, i int, minJumps *int) int {
	var steps int
	if i == len(nums)-1 {
		return 1
	}

	for j := 1; j <= nums[i]; j++ {
		if i <= len(nums)-1 {
			return steps + minJumpsRec(nums, i+j, minJumps)
		}
	}
	if *minJumps > steps {
		*minJumps = steps
	}

	return steps
}

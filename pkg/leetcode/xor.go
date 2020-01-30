package leetcode

func findMaximumXOR(nums []int) int {
	largest := 0
	secLargest := 0
	
	for _, n := range nums {
		if n > secLargest && n < largest {
			secLargest = n
		}
		
		if n > largest {
			largest = n
		}
	}
	
	return largest ^ secLargest
}

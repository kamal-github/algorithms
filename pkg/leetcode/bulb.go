package leetcode

// 2,1,3,5,4,6
func NumberOfBulbShines(bulbs []int) int {
	sumBulbs := 0
	shined := 0
	largest := bulbs[0]
	for _, b := range bulbs {
		largest = larger(largest, b)
		sumBulbs += b
		if sumBulbs == sumOfInts(largest) {
			shined++
		}
	}
	return shined
}

func larger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func sumOfInts(n int) int {
	return (n*(n+1))/2
}

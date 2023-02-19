package main

// height = [1,8,6,2,5,4,8,3,7]
func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1

	for left != right {
		var runningArea int

		h := findSmallerHeight(height[left], height[right])
		b := right - left
		runningArea = h * b

		if runningArea > max {
			max = runningArea
		}

		shiftSmallerHeight(height, &left, &right)
	}

	return max
}

func shiftSmallerHeight(height []int, l, r *int) {
	if height[*l] < height[*r] {
		*l = *l + 1
		return
	}

	*r = *r - 1
}

func findSmallerHeight(hl, hr int) int {
	if hl < hr {
		return hl
	}

	return hr
}

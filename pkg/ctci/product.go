package ctci

// [1,2,3,4,5,6,7] ; n is length
// [1,2,3]

func ProductOfArrayExceptSelf(ans, inp []int, i int) int {
	n := len(inp)

	if len(inp) == 0 {
		return 1
	}

	if i == 0 {
		ans[i] = inp[i+1] * ProductOfArrayExceptSelf(ans, inp[i+1:], i+1)
	} else if i == n-1 {
		ans[i] = inp[i-1] * ProductOfArrayExceptSelf(ans, inp[0:i-1], i+1)
	} else {
		ans[i] = inp[i-1] * ProductOfArrayExceptSelf(ans, inp[0:i-1], i+1) * inp[i+1] * ProductOfArrayExceptSelf(ans, inp[i+1:], i+1)
	}

	return ans[i]
}

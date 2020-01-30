package leetcode

import "github.com/kamal-github/algorithms/pkg/leetcode/common"

// O(n)
func MaxSubArraySumDP(a []int) int {
	mem := make([]int, len(a))
	mem[0] = a[0]

	for i := 1; i < len(a); i++ {
		mem[i] = common.Bigger(mem[i-1]+a[i], a[i])
	}

	return mem[len(a)-1]
}

// Time - O(n^2)
// Space - O(n^2)
func MaxSubArraySum(a []int) int {
	largest := 0
	inpLength := len(a)

	mem := make([][]int, inpLength)
	for i := 0; i < inpLength; i++ {
		mem[i] = make([]int, inpLength)
		mem[i][i] = a[i]
		if largest < mem[i][i] {
			largest = mem[i][i]
		}
	}

	for i := 0; i < inpLength-1; i++ {
		mem[i][i+1] = a[i] + a[i+1]
		if largest < mem[i][i+1] {
			largest = mem[i][i+1]
		}
	}

	for i := 3; i <= inpLength; i++ {
		for j := 0; j <= inpLength-i; j++ {
			k := i + j - 1
			mem[j][k] = mem[j][k-1] + mem[k][k]
			if largest < mem[j][k] {
				largest = mem[j][k]
			}
		}
	}

	return largest
}

//func MaxSubArraySum(a []int) int {
//	largest := 0
//	return maxSubArraySum(a, largest )
//}

//func maxSubArraySum(a []int, largest int) int {
//	arrLength := len(a)
//	if arrLength == 0 {
//		return 0
//	}
//	if arrLength == 2 {
//		return a[0]
//	}
//	mem := make([][]bool, arrLength)
//	for i:=0; i<arrLength; i++{
//		mem[i] = make([]bool, arrLength)
//		mem[i][i] = true
//	}}
//
//	for i := {
//		sum := max(maxSubArraySum(a[0:arrLength-1], largest) + a[arrLength-1], )
//		if sum > largest {
//			largest = sum
//		}
//	}
//	return largest
//}

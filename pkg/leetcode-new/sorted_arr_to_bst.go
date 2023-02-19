package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	if mid-1 >= 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}
	if mid+1 < len(nums) {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}

	return root
}

// -------- This also works but above is shorter version of below code. --------
//
// func sortedArrayToBST(nums []int) *TreeNode {
//   if nums == nil || len(nums) == 0 {
//     return nil
//   }
//   if len(nums) == 1 {
//     return &TreeNode{Val: nums[0]}
//   }

//   return buildBST(nums, 0, len(nums) - 1)
// }

// func buildBST(nums []int, l, h int) *TreeNode {
//   if l > h {
//     return nil
//   }
//   mid := findMid(nums, l, h)
//   root := &TreeNode{Val: nums[mid]}

//   root.Left = buildBST(nums, l, mid-1)
//   root.Right = buildBST(nums, mid+1, h)

//   return root
// }

// func findMid(nums []int, l, h int) int {
//   isEven := (l + h ) % 2 == 0
//   var m int
//   if isEven {
//     m = (l + h) / 2
//   } else {
//     m = (l + h + 1) / 2
//   }

//   return m
// }

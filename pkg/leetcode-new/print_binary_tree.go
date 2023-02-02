package main

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	h := height(root)
	m := h
	n := int(math.Pow(2, float64(m)) - 1)

	res := make([][]string, m)

	for i := 0; i < m; i++ {
		row := make([]string, n)
		res[i] = append(res[i], row...)
	}

	traverseTree(root, res, 0, 0, n)

	return res
}

func traverseTree(root *TreeNode, res [][]string, level, start, end int) {
	if root == nil {
		return
	}

	mid := (start + end - 1) / 2
	res[level][mid] = strconv.Itoa(root.Val)

	traverseTree(root.Left, res, level+1, start, mid)
	traverseTree(root.Right, res, level+1, mid+1, end)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

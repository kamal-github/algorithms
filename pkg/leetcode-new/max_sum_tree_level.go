package main

type level int

func maxLevelSum(root *TreeNode) int {
	sumAtLevel := make(map[level]int)
	maxLevel := 1

	traverseTree2(root, 1, sumAtLevel)

	// iterate over sumAtLevel to find the max
	for l, sum := range sumAtLevel {
		if sum > sumAtLevel[level(maxLevel)] {
			maxLevel = int(l)
		}
	}

	return maxLevel
}

func traverseTree2(root *TreeNode, l level, sumAtLevel map[level]int) {
	if root == nil {
		return
	}
	if sum, ok := sumAtLevel[l]; ok {
		sumAtLevel[l] = sum + root.Val
	} else {
		sumAtLevel[l] = root.Val
	}
	traverseTree2(root.Left, l+1, sumAtLevel)
	traverseTree2(root.Right, l+1, sumAtLevel)
}

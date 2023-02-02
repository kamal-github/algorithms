package main

func rightSideView(root *TreeNode) []int {
	rightNodePresenceIndicator := make(map[level]bool)
	nodeValues := make([]int, 0)
	traverseTree3(root, 1, rightNodePresenceIndicator, &nodeValues)
	return nodeValues
}

func traverseTree3(root *TreeNode, l level, rightNodePresenceIndicator map[level]bool, nodeValues *[]int) {
	if root == nil {
		return
	}
	if _, ok := rightNodePresenceIndicator[l]; !ok {
		rightNodePresenceIndicator[l] = true
		*nodeValues = append(*nodeValues, root.Val)
	}

	traverseTree3(root.Right, l+1, rightNodePresenceIndicator, nodeValues)
	traverseTree3(root.Left, l+1, rightNodePresenceIndicator, nodeValues)
}

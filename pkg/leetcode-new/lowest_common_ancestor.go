package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// Idea: Base logic
// if p and q are on the opposite side then root is the LCA
// otherwise if they both are on the left side then keep looking left
// until they are on the opposite side.
// OR if they both are on the right side then keep looking right
// until they are on the opposite side.

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
package trees

type Node struct {
	V   int
	Left  *Node
	Right *Node
}

func Height(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + Max(Height(root.Left), Height(root.Right))
}

func Size(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + Size(root.Left) + Size(root.Right)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


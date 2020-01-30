package bst

import (
	"fmt"
)

//Common Node struct for AVL and BST
type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	V      int
	H      int //Height is only used for AVL
}

func newNode(l *Node, r *Node, p *Node, v int) *Node {
	return &Node{Left: l, Right: r, Parent: p, V: v, H: 1}
}

func Insert(root *Node, v int) *Node {
	if root == nil {
		root = newNode(nil, nil, nil, v)
		return root
	}

	if v < root.V {
		t := Insert(root.Left, v)
		root.Left = t
		t.Parent = root
	} else {
		t := Insert(root.Right, v)
		root.Right = t
		t.Parent = root
	}

	return root
}

// Return true/false if a value found
func Exists(root *Node, v int) bool {
	if root == nil {
		return false
	}

	if v == root.V {
		return true
	}

	if v < root.V {
		return Exists(root.Left, v)
	} else {
		return Exists(root.Right, v)
	}

}

// Return Node if found, else if current node has no left/right node [depends upon value of node]
// then return the current node
func Find(root *Node, v int) *Node {

	if v == root.V {
		return root
	}

	if v < root.V {
		//Optimization step (not needed but avoid one more function call)
		if root.Left != nil {
			return Find(root.Left, v)
		}
		return root
	} else {
		//Optimization step (not needed but avoid one more function call)
		if root.Right != nil {
			return Find(root.Right, v)
		}
		return root
	}

}

//return immediate bigger number aka inorder successor (smallest in the right subtree)
func Next(root *Node, v int) *Node {
	n := Find(root, v)

	if n.Right != nil {
		return Find(n.Right, v)
	} else {
		return rightAncestor(n, v)
	}

}

func rightAncestor(n *Node, v int) *Node {
	if n.Parent != nil {
		return nil
	}

	if n.Parent.V > n.V {
		return n.Parent
	} else {
		return rightAncestor(n.Parent, v)
	}
}

func Delete(root *Node, v int) *Node {
	if root == nil {
		return root
	}

	if v < root.V {
		root.Left = Delete(root.Left, v)
	} else if v > root.V {
		root.Right = Delete(root.Right, v)
	} else {

		if root.Left == nil {
			temp := root.Right
			root = nil // free the root
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil // free the root
			return temp
		}

		succNode := Next(root, v)
		root.V = succNode.V
		root.Right = Delete(root.Right, succNode.V)
	}

	return root
}

func Inorder(root *Node) {
	if root == nil {
		return
	}
	Inorder(root.Left)
	fmt.Println(root.V)
	Inorder(root.Right)
}

func Preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.V)
	Preorder(root.Left)
	Preorder(root.Right)
}

//Inorder of BST always give sorted seq
func GetSortedArr(root *Node, arr *[]int) {

	if root == nil {
		return
	}

	GetSortedArr(root.Left, arr)
	*arr = append(*arr, root.V)
	GetSortedArr(root.Right, arr)

}

func Balance(arr []int, s int, e int) *Node {
	if s > e {
		return nil
	}

	m := ((s + e) / 2)

	root := newNode(nil, nil, nil, arr[m])
	root.Left = Balance(arr, s, m-1)
	root.Right = Balance(arr, m+1, e)

	return root
}

package trees

import (
	"fmt"

	"github.com/algorithms/pkg/trees"
)

//AVL tree node has height as extra attribute
func newNode(l *trees.Node, r *trees.Node, v int) *trees.Node {
	return &trees.Node{Left: l, Right:r, V:v, H:1}
}


func Insert(root *trees.Node, v int) *trees.Node {

	// INFO - We can not reuse bst insert method - as it always return the top element
	//step 1 - perform normal BST insertion
	if root == nil {
		root = newNode(nil, nil, v)
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

	// Step 2 - update the height
	root.H = 1 +  trees.Max(getHeight(root.Left), getHeight(root.Right))
	fmt.Println("Height of Root", root.V, root.H)

	//Step 3 - get balance factor

	balFact := getBalanceFactor(root)

	// Step 4 - Rotation
	if balFact > 1 && v < root.Left.V  {
		//Case 1 LeftLeft - left tree is heavier than right (balFactor >= 2)
		root = rotateRight(root)
	} else if balFact < -1 && v > root.Right.V  {
		//Case 2 - RightRight- right tree is heavier than right (balFactor <= -2)
		root = rotateLeft(root)
	} else if balFact > 1 && v > root.Left.V {
		//Case 3 LeftRight- left tree is heavier than left (balFactor >= 2)
		root.Left = rotateLeft(root.Left)
		root = rotateRight(root)
	} else if balFact < -1 && v < root.Right.V {
		//Case 4 - RightLeft - right tree is heavier than left (balFactor <= -2)
		root.Right = rotateRight(root.Right)
		root = rotateLeft(root)
	}


	fmt.Println("Height of Root post rotation", root.V, root.H)
	return root
}

func Delete(root *trees.Node, v int) *trees.Node {

	//step 1 - perform normal BST deletion
	//root = bst.Delete(root, v)

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

		succNode := minValueNode(root.Right, v)
		root.V = succNode.V
		root.Right = Delete(root.Right, succNode.V)

	}

	// Step 2 - update the height
	root.H = 1 +  trees.Max(getHeight(root.Left), getHeight(root.Right))

	//Step 3 - get balance factor

	balFact := getBalanceFactor(root)

	// Step 4 - Rotation
	if balFact > 1 && getBalanceFactor(root.Left) >= 0  {
		//Case 1 LeftLeft - left tree is heavier than right (balFactor >= 2)
		root = rotateRight(root)
	} else if balFact < -1 && getBalanceFactor(root.Right) <= 0 {
		//Case 2 - RightRight- right tree is heavier than right (balFactor <= -2)
		root = rotateLeft(root)
	} else if balFact > 1 && getBalanceFactor(root.Left) < 0 {
		//Case 3 LeftRight- left tree is heavier than left (balFactor >= 2)
		root.Left = rotateLeft(root.Left)
		root = rotateRight(root)
	} else if balFact < -1 && getBalanceFactor(root.Right) < 0 {
		//Case 4 - RightLeft - right tree is heavier than left (balFactor <= -2)
		root.Right = rotateRight(root.Right)
		root = rotateLeft(root)
	}

	return root
}


//return immediate bigger number aka inorder successor (smallest in the right subtree)
func minValueNode(root *trees.Node, v int) *trees.Node {
	//currNode := root
	//for  currNode.Left != nil {
	//	currNode = currNode.Left
	//}
	//
	//return currNode
	if root.Left == nil {
		return root
	}

	return minValueNode(root.Left, v)
}

func rotateRight(root *trees.Node) *trees.Node {
	x := root.Left
	T2 := x.Right

	root.Left = T2
	x.Right = root

	//update height
	root.H =  1 + trees.Max(getHeight(root.Left), getHeight(root.Right))
	x.H = 1 + trees.Max(getHeight(x.Left), getHeight(x.Right))

	return x
}

func rotateLeft(root *trees.Node) *trees.Node {
	x := root.Right
	T2 := x.Left

	x.Left = root
	root.Right = T2

	//update height
	root.H =  1 + trees.Max(getHeight(root.Left), getHeight(root.Right))
	x.H = 1 + trees.Max(getHeight(x.Left), getHeight(x.Right))

	return x

}

func getHeight(n *trees.Node) int {
	if n == nil {
		return 0
	}
	return n.H

}

func getBalanceFactor(n *trees.Node) int {
	if n == nil {
		return 0
	}

	return getHeight(n.Left) - getHeight(n.Right)
}
func Preorder(root *trees.Node) {
	if root == nil {
		return
	}
	fmt.Println(root.V, getBalanceFactor(root))
	Preorder(root.Left)
	Preorder(root.Right)
}


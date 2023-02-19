package ctci

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type Node struct {
	Val  int
	Next *Node
}

//func (ll *LinkedList) IsPalindrome() bool {
//	trav := trav2 := ll.Head
//	palindrome(trav, trav2)
//}
//
//func palindrome() int {
//
//}

/*
Intersection: Given two (singly) linked lists, determine if the two lists intersect. Return the intersecting
node. Note that the intersection is defined based on reference, not value. That is, if the kth
node of the first linked list is the exact same node (by reference) as the jth node of the second
linked list, then they are intersecting.
*/
func Intersection(l1, l2 *LinkedList) *Node {
	h1 := l1.Head
	h2 := l2.Head
	var l1Len, l2Len int

	for h1 != nil {
		l1Len++
		h1 = h1.Next
	}

	for h2 != nil {
		l2Len++
		h2 = h2.Next
	}

	var trav *Node

	if l1Len > l2Len {
		trav = l1.kthToTheFirst(l1Len - l2Len)
		return intersectingNode(trav, h2)
	} else if l2Len > l1Len {
		trav = l2.kthToTheFirst(l2Len - l1Len)
		return intersectingNode(trav, h1)
	} else {
		return intersectingNode(h1, h2)
	}
}

func (ll *LinkedList) kthToTheFirst(k int) *Node {
	i := 0
	t := ll.Head

	for i != k && t != nil {
		i++
		t = t.Next
	}

	return t
}

func intersectingNode(n1, n2 *Node) *Node {
	for n1 != nil && n2 != nil {
		if n1 == n2 {
			return n1
		}

		n1 = n1.Next
		n2 = n2.Next
	}

	return nil
}

/*
Loop Detection: Given a circular linked list, implement an algorithm that returns the node at the
beginning of the loop.
DEFINITION
Circular linked list: A (corrupt) linked list in which a node's Next pointer points to an earlier node, so
as to make a loop in the linked list.
EXAMPLE
Input: A -> B -> C - > D -> E -> C [the same C as earlier]
Output: C
*/
func (ll *LinkedList) DetectLoop() *Node {
	//slow := ll.Head
	//fast := ll.Head
	return nil
}

/*
Partition: Write code to partition a linked list around a value x, such that all nodes less than x come
before all nodes greater than or equal to x. If x is contained within the list, the values of x only need
to be after the elements less than x (see below). The partition element x can appear anywhere in the
"right partition"; it does not need to appear between the left and right partitions.
*/
//3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1
func (ll *LinkedList) Partition(x int) *LinkedList {
	smallerLL := NewLinkedList()
	LargerLL := NewLinkedList()

	if ll.Head == nil {
		return smallerLL
	}

	trav := ll.Head
	for trav != nil {
		if trav.Val < x {
			smallerLL.AddToTail(trav.Val)
		} else {
			LargerLL.AddToTail(trav.Val)
		}

		trav = trav.Next
	}

	trav = smallerLL.Head

	for trav.Next != nil {
		trav = trav.Next
	}

	trav.Next = LargerLL.Head

	return smallerLL
}

/*
Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but
the first and last node, not necessarily the exact middle) of a singly linked list, given only access to
that node.
EXAMPLE
lnput:the node c from the linked lista->b->c->d->e->f
Result: nothing is returned, but the new linked list looks like a->b->d->e- >f
*/
func (ll LinkedList) RemoveMid() {
	if ll.Head == nil {
		return
	}

	trav := ll.Head.Next

	if trav == nil {
		// it means there is no element except dummy node, so dups
		return
	}

	prev := trav
	prevToPrev := prev
	prev = prev.Next
	trav = trav.Next.Next

	for trav != nil && trav.Next != nil {
		prevToPrev = prevToPrev.Next
		prev = prev.Next
		trav = trav.Next.Next
	}

	prevToPrev.Next = prev.Next
}

// Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
// 1 2 3 4 5 6 7
func (ll LinkedList) KthToTheLast(k int) int {
	if ll.Head == nil {
		return -1
	}

	trav := ll.Head.Next

	if trav == nil {
		// it means there is no element except dummy node, so dups
		return -1
	}

	prev := trav
	for i := 0; i < k; i++ {
		trav = trav.Next
	}

	for trav != nil {
		prev = prev.Next
		trav = trav.Next
	}

	return prev.Val
}

// Remove Duplicates - Write code to remove duplicates from an unsorted linked list.
// 3 3 4 4 5 5
// 3 4 5 6 7
func (ll LinkedList) RemoveDuplicates() {
	m := make(map[int]bool)

	if ll.Head == nil {
		return
	}

	trav := ll.Head.Next

	if trav == nil {
		// it means there is no element except dummy node, so dups
		return
	}

	prev := trav
	m[trav.Val] = true
	trav = trav.Next

	for trav != nil {
		if _, ok := m[trav.Val]; !ok {
			m[trav.Val] = true
			trav = trav.Next
			prev = prev.Next
		} else {
			// remove the duplicate node
			prev.Next = trav.Next
			trav = trav.Next
		}
	}
}

// Add to the tail of the linked list, if its the first element,
// create a new linkelist with a new node
func (ll *LinkedList) AddToTail(data int) {
	newNode := &Node{
		Val:  data,
		Next: nil,
	}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	trav := ll.Head
	for trav.Next != nil {
		trav = trav.Next
	}

	trav.Next = newNode
}

func (ll *LinkedList) Print() {
	if ll.Head == nil {
		return
	}

	fmt.Println()
	trav := ll.Head
	for trav != nil {
		fmt.Printf("%d->", trav.Val)
		trav = trav.Next
	}
	fmt.Println()
}

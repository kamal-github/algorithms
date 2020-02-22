package ctci

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type Node struct {
	data int
	next *Node
}

//func (ll *LinkedList) IsPalindrome() bool {
//	trav := trav2 := ll.head
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
	h1 := l1.head
	h2 := l2.head
	var l1Len, l2Len int

	for h1 != nil {
		l1Len++
		h1 = h1.next
	}

	for h2 != nil {
		l2Len++
		h2 = h2.next
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
	t := ll.head

	for i != k && t != nil {
		i++
		t = t.next
	}

	return t
}

func intersectingNode(n1, n2 *Node) *Node {
	for n1 != nil && n2 != nil {
		if n1 == n2 {
			return n1
		}

		n1 = n1.next
		n2 = n2.next
	}

	return nil
}

/*
Loop Detection: Given a circular linked list, implement an algorithm that returns the node at the
beginning of the loop.
DEFINITION
Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node, so
as to make a loop in the linked list.
EXAMPLE
Input: A -> B -> C - > D -> E -> C [the same C as earlier]
Output: C
*/
func (ll *LinkedList) DetectLoop() *Node {
	//slow := ll.head
	//fast := ll.head
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

	if ll.head == nil {
		return smallerLL
	}

	trav := ll.head
	for trav != nil {
		if trav.data < x {
			smallerLL.AddToTail(trav.data)
		} else {
			LargerLL.AddToTail(trav.data)
		}

		trav = trav.next
	}

	trav = smallerLL.head

	for trav.next != nil {
		trav = trav.next
	}

	trav.next = LargerLL.head

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
	if ll.head == nil {
		return
	}

	trav := ll.head.next

	if trav == nil {
		// it means there is no element except dummy node, so dups
		return
	}

	prev := trav
	prevToPrev := prev
	prev = prev.next
	trav = trav.next.next

	for trav != nil && trav.next != nil {
		prevToPrev = prevToPrev.next
		prev = prev.next
		trav = trav.next.next
	}

	prevToPrev.next = prev.next
}

// Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
// 1 2 3 4 5 6 7
func (ll LinkedList) KthToTheLast(k int) int {
	if ll.head == nil {
		return -1
	}

	trav := ll.head.next

	if trav == nil {
		// it means there is no element except dummy node, so dups
		return -1
	}

	prev := trav
	for i := 0; i < k; i++ {
		trav = trav.next
	}

	for trav != nil {
		prev = prev.next
		trav = trav.next
	}

	return prev.data
}

// Remove Duplicates - Write code to remove duplicates from an unsorted linked list.
//3 3 4 4 5 5
// 3 4 5 6 7
func (ll LinkedList) RemoveDuplicates() {
	m := make(map[int]bool)

	if ll.head == nil {
		return
	}

	trav := ll.head.next

	if trav == nil {
		// it means there is no element except dummy node, so dups
		return
	}

	prev := trav
	m[trav.data] = true
	trav = trav.next

	for trav != nil {
		if _, ok := m[trav.data]; !ok {
			m[trav.data] = true
			trav = trav.next
			prev = prev.next
		} else {
			// remove the duplicate node
			prev.next = trav.next
			trav = trav.next
		}
	}
}

// Add to the tail of the linked list, if its the first element,
// create a new linkelist with a new node
func (ll *LinkedList) AddToTail(data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	trav := ll.head
	for trav.next != nil {
		trav = trav.next
	}

	trav.next = newNode
}

func (ll *LinkedList) Print() {
	if ll.head == nil {
		return
	}

	fmt.Println()
	trav := ll.head
	for trav != nil {
		fmt.Printf("%d->", trav.data)
		trav = trav.next
	}
	fmt.Println()
}

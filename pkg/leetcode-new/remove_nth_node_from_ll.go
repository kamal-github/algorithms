package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Iterative one pass.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	firstNode := head
	secondNode := head
	i := 0
	for firstNode != nil {
		if i > n {
			secondNode = secondNode.Next
		}
		firstNode = firstNode.Next
		i += 1
	}

	if n == i {
		head = head.Next
	} else {
		secondNode.Next = secondNode.Next.Next
	}
	return head
}

// Recursive
// TODO: fails for 1 -> 2 and n = 1
func removeNthFromEndRec(head *ListNode, n int) *ListNode {
	var deleted bool
	head, _ = remove(head, head, n, &deleted)
	return head
}

func remove(head, tail *ListNode, n int, deleted *bool) (*ListNode, int) {
	if tail == nil {
		return head, 0
	}

	head, level := remove(head, tail.Next, n, deleted)
	level++
	if level == n+1 && !(*deleted) {
		head = head.Next
		*deleted = true
	} else if tail == head && !(*deleted) {
		temp := tail.Next
		tail.Next = tail.Next.Next
		temp.Next = nil
		*deleted = true
	}

	return head, level
}

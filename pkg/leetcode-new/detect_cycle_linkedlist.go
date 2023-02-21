package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Run O(n) and space O(n)
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	seen := make(map[*ListNode]bool)
	seen[head] = true

	for head.Next != nil {
		if _, ok := seen[head.Next]; ok {
			return true
		}

		seen[head.Next] = true
		head = head.Next
	}

	return false
}

// Alternative efficient Solution.
// Fast and Slow Algorithm aka Hare and tortoise algo
// Run O(n) and space O(1)
func hasCycleConstSpace(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

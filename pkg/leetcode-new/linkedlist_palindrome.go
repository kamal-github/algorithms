package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isLLPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	stack := make([]*ListNode, 0)
	fast, slow := head, head
	for fast != nil {
		stack = append(stack, slow)
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	// slow is at mid and fast is at end or one before
	for slow != nil {
		if fast.Next == nil {
			if len(stack) == 0 {
				break
			}
			poppedNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if poppedNode.Val != slow.Val {
				return false
			}
			slow = slow.Next
		} else {
			if len(stack) == 0 {
				break
			}
			poppedNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			slow = slow.Next
			if poppedNode.Val != slow.Val {
				return false
			}
		}
	}

	return true
}

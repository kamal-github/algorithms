package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// TODO Replace this with actual LinkedList implementation
type ListNode struct {
      Val int
      Next *ListNode
}

// Even Case1 : 1->2->3->4   => 2->1->3->4
// Odd Case2 : 1->2->3->4->5 => 2->1->4->3->5
func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	newHead := &ListNode{
		Val: head.Next.Val,
		Next: &ListNode{Val: head.Val},
	}

	newHead.Next.Next = SwapPairs(head.Next.Next)

	return newHead
}


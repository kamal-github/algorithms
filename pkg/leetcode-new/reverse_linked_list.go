package main

// recursive
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	revHead, _ := reverse(head)
	return revHead
}

func reverse(head *ListNode) (revHead, revTail *ListNode) {
	if head.Next == nil {
		revHead = &ListNode{Val: head.Val}
		return revHead, revHead
	}
	revHead, revTail = reverse(head.Next)
	revTail.Next = &ListNode{Val: head.Val}
	return revHead, revTail.Next
}

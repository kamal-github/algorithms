package leetcode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rem := 0
	
	// Dummy node
	head := new(ListNode)
	trav := head
	
	for l1 != nil && l2 != nil {
		trav.Next = new(ListNode)
		trav = trav.Next
		
		sum := l1.Val + l2.Val + rem
		rem = sum / 10
		
		trav.Val = sum % 10
		
		l1 = l1.Next
		l2 = l2.Next
	}
	
	for l1 != nil && trav != nil {
		sum := l1.Val + rem
		rem = sum / 10
		
		trav.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		
		l1 = l1.Next
		trav = trav.Next
	}
	
	for l2 != nil && trav != nil {
		sum := l2.Val + rem
		rem = sum / 10
		
		trav.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		
		l2 = l2.Next
		trav = trav.Next
	}
	
	if rem > 0 && trav != nil {
		trav.Next = &ListNode{
			Val:  rem,
			Next: nil,
		}
	}
	
	// return the node after dummy node, which makes a correct LL
	return head.Next
}

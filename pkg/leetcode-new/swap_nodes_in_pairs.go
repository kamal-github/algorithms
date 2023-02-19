package main

func swapPairs(ll *LinkedList) *LinkedList {
	head := ll.Head
	if head == nil || head.Next == nil {
		return &LinkedList{Head: head}
	}

	secondNode := ListNode{}
	firstNode := ListNode{Next: &secondNode}
	resultList := LinkedList{Head: &firstNode}
	cur := &firstNode

	for head != nil && head.Next != nil {
		cur.Next.Val, cur.Val = head.Val, head.Next.Val

		head = head.Next.Next

		if head != nil && head.Next != nil {
			second := ListNode{}
			first := ListNode{Next: &second}
			cur.Next.Next = &first
			cur = &first
		}
	}

	if head != nil && head.Next == nil {
		lastNode := ListNode{Val: head.Val}
		cur.Next.Next = &lastNode
	}

	return &resultList
}

// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
func swapNodes(ll *LinkedList, k int) *LinkedList {
	if k == 0 {
		return nil
	}
	if ll == nil && ll.Head != nil {
		return ll
	}

	resLinkedList := new(LinkedList)
	numNodes := 0

	curNode := ll.Head
	for curNode != nil {
		resLinkedList.InsertEnd(curNode.Val)
		numNodes++
		curNode = curNode.Next
	}

	kFromBack := numNodes - k + 1 // (n - k + 1) as it is 1-indexed, otherwise (n - k)
	curNode = resLinkedList.Head
	curPos := 1 // as per question, LL are indexed 1.

	var kthNode, kthFromBackNode *ListNode

	for curPos != kFromBack {
		if curPos == k {
			kthNode = curNode
		}

		curNode = curNode.Next
		curPos++
	}
	if curPos == kFromBack {
		kthFromBackNode = curNode
	}

	kthNode.Val, kthFromBackNode.Val = kthFromBackNode.Val, kthNode.Val

	return resLinkedList
}

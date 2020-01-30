package leetcode

import "math"

// My solution using Circular Linked list
func MyMaxRotateFunction(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	cl := makeCircularList(A)
	head := cl
	trav := head
	max := math.MinInt64
	
	for i:=0; i<len(A); i++ {
		localMax := 0
		j := 0
		
		localMax += j * trav.Val
		j++
		trav = trav.Next
		
		for trav != head {
			localMax += j * trav.Val
			j++
			trav = trav.Next
		}
		
		if max < localMax {
			max = localMax
		}
		//fmt.Println(localMax, max)
		trav = trav.Next
		head = trav
	}
	
	return max
}

// Leetcode Discuss page solution (not mine!)
func MaxRotateFunction(A []int) int {
	summary := 0
	base := 0
	for i, v := range A {
		summary += v
		base += i*v
	}
	
	max := base
	for i := 0; i < len(A)-1; i++ {
		// rotate once:
		// step 1: plus all elements once
		// step 2: minus current last-elem len(A) times
		base += summary
		base -= len(A)*A[len(A)-1-i]
		if base > max {
			max = base
		}
	}
	return max
}

func makeCircularList(A []int) *ListNode {
	dummy := new(ListNode)
	head := dummy
	for _, e := range A {
		head.Next = &ListNode{
			Val:  e,
			Next: nil,
		}
		head = head.Next
	}
	head.Next = dummy.Next
	head = head.Next
	
	return head
}

package leetcode

// MergeTwoSortedInPlace - https://leetcode.com/problems/merge-sorted-array/
func MergeTwoSortedInPlace(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >=0 && j>=0 {
		if nums1[i] > nums2[j] {
			nums1[k]= nums1[i]
			i--
		} else if nums1[i] < nums2[j] {
			nums1[k]= nums2[j]
			j--
		} else {
			nums1[k]= nums1[i]
			i--
			k--
			nums1[k]= nums2[j]
			j--
		}
		k--
	}
	
	// Remaining numbers
	for j >=0 {
		nums1[k]= nums2[j]
		j--
		k--
	}
}

// Recursive Solution --------------
//Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
//Examples:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4
// I1 = 1-6-7 , I2 = 3-4-5-6-8
// O = 1-3-4-5-6-6-7-8
func MergeTwoListsRec(l1, l2 *ListNode ) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	l3 := new(ListNode)
	if l1.Val < l2.Val {
		l3.Val = l1.Val
		l3.Next = MergeTwoListsRec(l1.Next, l2)
	} else if l2.Val < l1.Val{
		l3.Val = l2.Val
		l3.Next = MergeTwoListsRec(l1, l2.Next)
	} else {
		l3.Val = l1.Val
		l3.Next = &ListNode{Val: l2.Val}
		l3.Next.Next = MergeTwoListsRec(l1.Next, l2.Next)
	}

	return l3
}

// Iterative Solution ----------
//Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
//Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)

	l1Head := l1
	l2Head := l2
	l3Head := l3

	for l1Head != nil && l2Head != nil {
		if l1Head.Val < l2Head.Val {
			// set value
			l3Head.Val = l1Head.Val
			// update l1
			l1Head = l1Head.Next
			// update l3
			l3Head.Next = &ListNode{}
			l3Head = l3Head.Next
		} else if l2Head.Val < l1Head.Val {
			// set value
			l3Head.Val = l2Head.Val
			// update l2
			l2Head = l2Head.Next
			// update l3
			l3Head.Next = &ListNode{}
			l3Head = l3Head.Next
		} else {
			// set L1 to L3
			l3Head.Val = l1Head.Val
			// update l3
			l3Head.Next = &ListNode{}
			l3Head = l3Head.Next

			//set L2 to L3
			l3Head.Val = l2Head.Val
			// update l3
			l3Head.Next = &ListNode{}
			l3Head = l3Head.Next

			// update l2 and l1
			l1Head = l1Head.Next
			l2Head = l2Head.Next
		}

	}

	// If two given lists are not of equal size
	if l2Head != nil {
		for l2Head != nil {
			//set L2 to L3
			l3Head.Val = l2Head.Val
			// update l3
			l3Head.Next = &ListNode{}
			l3Head = l3Head.Next
			l2Head = l2Head.Next
		}

	} else {
		for l1Head != nil {
			//set L1 to L3
			l3Head.Val = l1Head.Val
			// update l3
			l3Head.Next = &ListNode{}
			l3Head = l3Head.Next
			l2Head = l1Head.Next
		}

	}
	return l3
}


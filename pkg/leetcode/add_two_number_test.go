package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/kamal-github/algorithms/pkg/leetcode"
)

func TestAddTwoNumbersWhenNoRemiander(t *testing.T) {
	l1 := leetcode.ListNode{
		Val: 2,
		Next: &leetcode.ListNode{
			Val: 9,
			Next: &leetcode.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := leetcode.ListNode{
		Val: 2,
		Next: &leetcode.ListNode{
			Val: 1,
			Next: &leetcode.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l3 := leetcode.ListNode{
		Val: 4,
		Next: &leetcode.ListNode{
			Val: 0,
			Next: &leetcode.ListNode{
				Val:  7,
				Next: nil,
			},
		},
	}

	actual := leetcode.AddTwoNumbers(&l1, &l2)
	if !compare(t, &l3, actual) {
		t.Fail()
	}
}

func TestAddTwoNumbersWhenAdditionalRemainder(t *testing.T) {
	l1 := leetcode.ListNode{
		Val: 2,
		Next: &leetcode.ListNode{
			Val: 9,
			Next: &leetcode.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := leetcode.ListNode{
		Val: 2,
		Next: &leetcode.ListNode{
			Val: 1,
			Next: &leetcode.ListNode{
				Val:  7,
				Next: nil,
			},
		},
	}
	l3 := leetcode.ListNode{
		Val: 4,
		Next: &leetcode.ListNode{
			Val: 0,
			Next: &leetcode.ListNode{
				Val: 1,
				Next: &leetcode.ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	actual := leetcode.AddTwoNumbers(&l1, &l2)
	if !compare(t, &l3, actual) {
		t.Fail()
	}

}

func TestAddTwoNumbersWhenInputListsSizeDifferent(t *testing.T) {
	l1 := leetcode.ListNode{
		Val: 2,
		Next: &leetcode.ListNode{
			Val: 9,
			Next: &leetcode.ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := leetcode.ListNode{
		Val: 2,
		Next: &leetcode.ListNode{
			Val: 1,
			Next: &leetcode.ListNode{
				Val: 7,
				Next: &leetcode.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	l3 := leetcode.ListNode{
		Val: 4,
		Next: &leetcode.ListNode{
			Val: 0,
			Next: &leetcode.ListNode{
				Val: 1,
				Next: &leetcode.ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	actual := leetcode.AddTwoNumbers(&l1, &l2)
	if !compare(t, &l3, actual) {
		t.Fail()
	}

}

func compare(t *testing.T, l1 *leetcode.ListNode, l2 *leetcode.ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if (l1 != nil && l2 == nil) || (l2 != nil && l1 == nil) {
		t.Log("lists length mismatch")
		return false
	}

	if l1.Val == l2.Val {
		fmt.Println(l1.Val, l2.Val)
		return compare(t, l1.Next, l2.Next)
	} else {
		t.Log("lists values mismatch", l1.Val, l2.Val)
		return false
	}
}

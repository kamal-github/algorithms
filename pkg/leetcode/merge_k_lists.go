package leetcode

import (
	"container/heap"
	"math"

	"github.com/kamal-github/algorithms/pkg/leetcode/common"
)

func MergeKLists(lists []*ListNode) *ListNode {
	out := new(ListNode)
	mergeHeap(lists, out)
	return out.Next
}

func mergeHeap(lists []*ListNode, out *ListNode) {

	h := &common.IntHeap{}
	heap.Init(h)

	for _, list := range lists {
		for list != nil {
			heap.Push(h, list.Val)
			list = list.Next
		}
	}

	for h.Len() > 0 {
		e := heap.Pop(h)
		out.Next = &ListNode{
			Val:  e.(int),
			Next: nil,
		}
		out = out.Next
	}
}

func merge(lists []*ListNode, out *ListNode) {
	var nilLists int
	//nilAccounted := make(map[int]bool)
	k := len(lists)

	for nilLists < k {
		for i, l := range lists {
			if l == nil {
				//if _, yes := nilAccounted[i]; !yes {
				nilLists += 1
				lists = append(lists[:i], lists[i+1:]...)
				//nilAccounted[i] = true
				//}
				continue
			}

			if l.Val == findLeast(lists) {
				out.Next = &ListNode{
					Val:  l.Val,
					Next: nil,
				}
				out = out.Next
				lists[i] = l.Next
			}
		}
	}

}

func findLeast(lists []*ListNode) int {
	// Increase to `MaxInt64` if numbers are expected to be Big
	least := math.MaxInt8
	for _, l := range lists {
		if l != nil && l.Val <= least {
			least = l.Val
		}
	}
	return least
}

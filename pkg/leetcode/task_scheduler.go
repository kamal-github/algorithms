package leetcode

import (
	"container/heap"

	"github.com/kamal-github/algorithms/pkg/leetcode/common"
)

/*
Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation: A -> B -> idle -> A -> B -> idle -> A -> B.
*/

func LeastInterval(tasks []byte, n int) int {
	countCache := make(map[byte]int, 0)

	for _, c := range tasks {
		countCache[c]++
	}

	// Adding into min heap
	h := &common.IntHeap{}
	heap.Init(h)

	for _, v := range countCache {
		heap.Push(h, -v)
	}
	maxVal := (-heap.Pop(h).(int)) - 1

	idleTime := maxVal * n

	for h.Len() > 0 {
		idleTime -= common.Smaller(maxVal, -heap.Pop(h).(int))
	}

	if idleTime < 0 {
		return len(tasks)
	}
	return idleTime + len(tasks)
}

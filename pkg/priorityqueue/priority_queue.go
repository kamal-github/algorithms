package priorityqueue

import (
	"errors"
	"fmt"
)


//type Item struct {
//	value    int
//	index    int  // i
//	priority int  //`key / k`
//}

type Pq []int

func (pq *Pq) exch(m, n int) {
	t := (*pq)[m]
	(*pq)[m] = (*pq)[n]
	(*pq)[n] = t
}

// Min heap (just change the comparison opr to make it max heap)
func (pQueue *Pq) swim(k int) {
	pq := *pQueue
	for (k/2) != 0  {
		if pq[k/2] > pq[k] {
			pq.exch(k/2, k)
		}
		k = k/2
	}
	*pQueue = pq

}

// Min heap (just change the comparison opr to make it max heap)
func (pq *Pq) sink(k int) {
	fmt.Println("Satrting PQ", pq)

	n := len(*pq)

	fmt.Println("PQ length", len(*pq))

	var c int

	for n > 2*k   { //7 >= 6
		j := 2*k
		c = j    // c = 6
		if n > j + 1  && (*pq)[j] < (*pq)[j+1] {
			c = j + 1
		}
		fmt.Println("c", c)
		if ((*pq)[k] > (*pq)[c]) {
			pq.exch(k, c)
		}
		fmt.Println("PQ", pq)
		k = c
	}
}

func (pq *Pq) insert (item int) {
	*pq = append(*pq, item)
	pq.swim(len(*pq) - 1)
}

func (pq *Pq) delete () (int, error){
	//pq := *pQueue
	if (len(*pq) == 1) {
		return -1, errors.New("insufficient elements")
	}

	k := (*pq)[1]
	last := len(*pq) - 1
	pq.exch(last, 1)

	fmt.Println("PQ after xchange-- ", pq, len(*pq))

	(*pq) = (*pq)[0:last]

	fmt.Println("PQ after removing last element -- ", pq, len(*pq))
	pq.sink(1)

	return k, nil
}

// Heap sort uses priority queue
// Time - O(log n)
// Space - O(n) for storing additional priority queue
func (pq *Pq) heapSort() []int {
	sortedArr := make([]int, len(*pq))
	i := len(*pq) - 1
	for len(*pq) > 1 {
		k, err := pq.delete()
		if err != nil {
			fmt.Println("error while deleting", err)
		}
		// For descending order out of max-heap - uncomment below line
		//sortedArr = append(sortedArr, k)

		// For ascending order out of max-heap
		sortedArr[i]= k
		i--
	}
	return sortedArr
}

//Pq is []int - use Pq as it needs sink method to invoke
func createHeap(arr *Pq) {
	arrLen := len(*arr)

	for i := (arrLen / 2); i >= 1; i-- {
		arr.sink(i)
	}

	fmt.Println("Heap", arr)
}

// Uses NO additional space and creating Binary heap in the same input array.
// Space - O(1)
// Time - O(log n)
func improvedHeapSort() Pq {
	arr := &Pq{0,20,10, 30, 50, 5, 1, 3,2, 90}
	createHeap(arr)
	fmt.Println("Heap", arr)
	size := len(*arr) - 1
	for size > 1 {
		arr.exch(size, 1)
		size--
		tempArr := (*arr)[0:size]
		tempArr.sink(1)
	}
	return *arr
}

func main() {
	pq := make(Pq, 1)//new(Pq)
	//old := *pq
	pq[0] = 0
	pqPtr := &pq
	pqPtr.insert(100)
	pqPtr.insert(90)
	pqPtr.insert(300)
	pqPtr.insert(40)
	pqPtr.insert(30)
	pqPtr.insert(20)
	pqPtr.insert(50)

	fmt.Println("Final Priority Queue", *pqPtr)
	//sortedArr := pqPtr.heapSort()
	//fmt.Println(sortedArr)
	//sortedArr := improvedHeapSort()
	//fmt.Println("sortedArr", sortedArr)


}

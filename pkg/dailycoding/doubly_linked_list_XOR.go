package dailycoding

import (
	"fmt"
	"unsafe"
)

// XOR linked list
type DoublyLinkedList struct {
	head *DNode
	tail *DNode
}

type DNode struct {
	data int
	both *DNode
}

// Add element to the end keeping reference as XOR of previous and next node
//for current node in `both` field to optimize the memory
func (dl *DoublyLinkedList) Add(item int) {
	trav := dl.head
	n := &DNode{
		data: item,
		both: nil,
	}

	// If list is empty
	if trav == nil {
		dl.head = n
		dl.tail = n
		n.both = (*DNode)(unsafe.Pointer(nil))
		return
	}

	// When list is having at least one element
	prev := (*DNode)(unsafe.Pointer(nil))
	for uintptr(unsafe.Pointer(trav.both))^uintptr(unsafe.Pointer(prev)) != uintptr(unsafe.Pointer(nil)) {
		temp := trav
		trav = (*DNode)(unsafe.Pointer(uintptr(unsafe.Pointer(trav.both)) ^ uintptr(unsafe.Pointer(prev))))
		prev = temp
	}

	// Updating trav both with XOR of prev and new element
	trav.both = (*DNode)(unsafe.Pointer(uintptr(unsafe.Pointer(prev)) ^ uintptr(unsafe.Pointer(n))))

	// setting the recently added node with XOR of its previous as (trav ^ nil = trav)
	n.both = trav

	// maintaining tail - always a last node when added
	dl.tail = n
}

// Get at a given index
func (dl *DoublyLinkedList) Get(index int) int {
	prev := (*DNode)(unsafe.Pointer(nil))
	trav := dl.head
	i := 0

	for i != index {
		// tracking the previous
		temp := trav

		// moving to next as we are moving forward
		trav = (*DNode)(unsafe.Pointer(uintptr(unsafe.Pointer(trav.both)) ^ uintptr(unsafe.Pointer(prev))))

		// setting previous
		prev = temp
		i++
	}

	return trav.data
}

// Print forward
func (dl *DoublyLinkedList) Print() {
	prev := (*DNode)(unsafe.Pointer(nil))
	trav := dl.head

	fmt.Println()
	for trav != dl.tail {
		fmt.Printf("%d%s", trav.data, "->")
		// tracking the previous
		temp := trav

		// moving to next as we are moving forward
		trav = (*DNode)(unsafe.Pointer(uintptr(unsafe.Pointer(trav.both)) ^ uintptr(unsafe.Pointer(prev))))

		// setting previous
		prev = temp
	}

	fmt.Printf("%d", trav.data)
}

// Print Reverse
func (dl *DoublyLinkedList) PrintReverse() {
	prev := (*DNode)(unsafe.Pointer(nil))
	trav := dl.tail

	fmt.Println()
	for trav != dl.head {
		fmt.Printf("%d%s", trav.data, "->")

		temp := trav

		// Getting previous node as we are moving reverse
		trav = (*DNode)(unsafe.Pointer(uintptr(unsafe.Pointer(trav.both)) ^ uintptr(unsafe.Pointer(prev))))
		prev = temp
	}

	fmt.Printf("%d", trav.data)
	fmt.Println()
}

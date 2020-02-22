package dailycoding_test

import (
	"fmt"
	"testing"

	"github.com/kamal-github/algorithms/pkg/dailycoding"
)

func TestDoublyLinkedList_Add(t *testing.T) {
	dl := new(dailycoding.DoublyLinkedList)

	dl.Add(1)
	dl.Add(2)
	dl.Add(3)
	dl.Add(5)
	dl.Add(10)

	dl.Print()

	dl.PrintReverse()

	fmt.Printf("The element at index - %d is %d", 3, dl.Get(3))
}

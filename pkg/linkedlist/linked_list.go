package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  interface{}
	Next *Node
}

type LinkedList struct {
	first *Node
}

func (ll *LinkedList) InsertBeg(v interface{}) {
	newNode := &Node{Val: v, Next: ll.first}
	ll.first = newNode
}

func (ll *LinkedList) Remove(v interface{}) error {

	if ll.first == nil {
		return errors.New("LinkedList is empty")
	}

	if ll.first != nil {
		if ll.first.Val == v {
			ll.first = ll.first.Next
			return nil
		}
	}

	node := ll.first
	prevNode := node
	for node != nil {
		if v.(int) == node.Val.(int) {
			fmt.Println("matched")
			prevNode.Next = node.Next
			return nil
		}
		prevNode = node
		node = node.Next
	}

	return errors.New("could not remove as candidate not found")

}

func (ll *LinkedList) Find(v interface{}) (bool, error) {
	node := ll.first
	for node != nil {
		if v == node.Val {
			return true, nil
		}
		node = node.Next
	}
	return false, errors.New("could not find")
}

func (ll *LinkedList) loopOver() {
	node := ll.first
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	ll := new(LinkedList)
	ll.InsertBeg(1)
	ll.InsertBeg(2)
	ll.InsertBeg(3)
	ll.InsertBeg(4)
	ll.loopOver()

	found, err := ll.Find(4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(found)

	errR := ll.Remove(1)
	if errR != nil {
		fmt.Println(err)
	}
	fmt.Println("After removal")
	ll.loopOver()
}

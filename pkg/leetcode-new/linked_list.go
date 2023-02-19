package main

import (
	"errors"
	"fmt"
)

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (ll *LinkedList) InsertBeg(v interface{}) {
	newNode := &ListNode{Val: v, Next: ll.Head}
	ll.Head = newNode
}

func (ll *LinkedList) InsertEnd(v interface{}) {
	newNode := ListNode{Val: v}
	node := ll.Head
	if node == nil {
		ll.Head = &newNode
		return
	}
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &newNode
}

func (ll *LinkedList) Remove(v interface{}) error {

	if ll.Head == nil {
		return errors.New("LinkedList is empty")
	}

	if ll.Head != nil {
		if ll.Head.Val == v {
			ll.Head = ll.Head.Next
			return nil
		}
	}

	node := ll.Head
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
	node := ll.Head
	for node != nil {
		if v == node.Val {
			return true, nil
		}
		node = node.Next
	}
	return false, errors.New("could not find")
}

func (ll *LinkedList) loopOver() {
	node := ll.Head
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

//
//func main() {
//	ll := new(LinkedList)
//	ll.InsertBeg(1)
//	ll.InsertBeg(2)
//	ll.InsertBeg(3)
//	ll.InsertBeg(4)
//	ll.loopOver()
//
//	found, err := ll.Find(4)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(found)
//
//	errR := ll.Remove(1)
//	if errR != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("After removal")
//	ll.loopOver()
//}

package ctci

import (
	"reflect"
	"testing"
)

func TestLinkedList_RemoveDuplicates(t *testing.T) {
	type fields struct {
		head *Node
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "with empty Head",
			fields: fields{head: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{}
			//3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1
			ll.AddToTail(3)
			ll.AddToTail(5)
			ll.AddToTail(8)
			ll.AddToTail(5)
			ll.AddToTail(10)
			ll.AddToTail(2)
			ll.AddToTail(1)

			ll.Print()

			//ll.RemoveDuplicates()

			//ll.Print()

			//fmt.Println(ll.KthToTheLast(3))
			//fmt.Println(ll.KthToTheLast(6))

			//ll.RemoveMid()

			pList := ll.Partition(8)
			pList.Print()
		})
	}
}

func TestIntersection(t *testing.T) {
	type args struct {
		l1 *LinkedList
		l2 *LinkedList
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

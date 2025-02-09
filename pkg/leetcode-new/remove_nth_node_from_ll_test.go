package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head func() *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want func() *ListNode
	}{
		{
			name: "1",
			args: args{
				n: 2,
				head: func() *ListNode {
					ll := new(LinkedList)
					ll.InsertEnd(10)
					ll.InsertEnd(20)
					ll.InsertEnd(30)
					ll.InsertEnd(40)
					ll.InsertEnd(50)
					return ll.Head
				},
			},
			want: func() *ListNode {
				ll := new(LinkedList)
				ll.InsertEnd(10)
				ll.InsertEnd(20)
				ll.InsertEnd(30)
				ll.InsertEnd(50)
				return ll.Head
			},
		},
		{
			name: "2",
			args: args{
				n: 1,
				head: func() *ListNode {
					ll := new(LinkedList)
					ll.InsertEnd(10)
					ll.InsertEnd(20)
					return ll.Head
				},
			},
			want: func() *ListNode {
				ll := new(LinkedList)
				ll.InsertEnd(10)
				return ll.Head
			},
		},
		{
			name: "3",
			args: args{
				n: 1,
				head: func() *ListNode {
					ll := new(LinkedList)
					ll.InsertEnd(1)
					return ll.Head
				},
			},
			want: func() *ListNode {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want(), removeNthFromEndRec(tt.args.head(), tt.args.n), "removeNthFromEnd(%v, %v)", tt.args.head, tt.args.n)
		})
	}
}

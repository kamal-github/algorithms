package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head func() *ListNode
	}
	tests := []struct {
		name string
		args args
		want func() *ListNode
	}{
		{
			name: "",
			args: args{head: func() *ListNode {
				ll := new(LinkedList)
				ll.InsertEnd(1)
				ll.InsertEnd(2)
				ll.InsertEnd(3)
				return ll.Head
			}},
			want: func() *ListNode {
				ll := new(LinkedList)
				ll.InsertEnd(3)
				ll.InsertEnd(2)
				ll.InsertEnd(1)
				return ll.Head
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want(), reverseList(tt.args.head()), "reverseList(%v)", tt.args.head)
		})
	}
}

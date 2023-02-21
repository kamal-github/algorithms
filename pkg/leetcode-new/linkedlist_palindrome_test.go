package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isLLPalindrome(t *testing.T) {
	type args struct {
		head func() *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "palindrome with even number list",
			args: args{head: func() *ListNode {
				ll := new(LinkedList)
				ll.InsertEnd(1)
				ll.InsertEnd(2)
				ll.InsertEnd(5)
				ll.InsertEnd(5)
				ll.InsertEnd(2)
				ll.InsertEnd(1)
				return ll.Head
			}},
			want: true,
		},
		{
			name: "palindrome with even number list",
			args: args{head: func() *ListNode {
				ll := new(LinkedList)
				ll.InsertEnd(1)
				ll.InsertEnd(1)
				return ll.Head
			}},
			want: true,
		},

		{
			name: "not a palindrome",
			args: args{head: func() *ListNode {
				ll := new(LinkedList)
				ll.InsertEnd(1)
				ll.InsertEnd(2)
				ll.InsertEnd(5)
				ll.InsertEnd(5)
				ll.InsertEnd(2)
				return ll.Head
			}},
			want: false,
		},
		{
			name: "palindrome with odd number list",
			args: args{head: func() *ListNode {
				ll := new(LinkedList)
				ll.InsertEnd(1)
				ll.InsertEnd(2)
				ll.InsertEnd(1)
				return ll.Head
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isLLPalindrome(tt.args.head()), "isLLPalindrome(%v)", tt.args.head)
		})
	}
}

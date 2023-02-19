package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head func() *LinkedList
	}
	tests := []struct {
		name string
		args args
		want func() *LinkedList
	}{
		{
			args: args{head: func() *LinkedList {
				ll := new(LinkedList)
				ll.InsertEnd(1)
				ll.InsertEnd(2)
				ll.InsertEnd(3)
				ll.InsertEnd(4)

				return ll
			}},
			want: func() *LinkedList {
				ll := new(LinkedList)
				ll.InsertEnd(2)
				ll.InsertEnd(1)
				ll.InsertEnd(4)
				ll.InsertEnd(3)

				return ll
			},
		},
		{
			name: "2",
			args: args{head: func() *LinkedList {
				ll := new(LinkedList)
				ll.InsertEnd(1)
				ll.InsertEnd(2)
				ll.InsertEnd(3)

				return ll
			}},
			want: func() *LinkedList {
				ll := new(LinkedList)
				ll.InsertEnd(2)
				ll.InsertEnd(1)
				ll.InsertEnd(3)

				return ll
			},
		},
		{
			args: args{head: func() *LinkedList {
				ll := new(LinkedList)
				ll.InsertEnd(1)
				ll.InsertEnd(2)
				ll.InsertEnd(3)
				ll.InsertEnd(4)
				ll.InsertEnd(5)
				ll.InsertEnd(6)

				return ll
			}},
			want: func() *LinkedList {
				ll := new(LinkedList)
				ll.InsertEnd(2)
				ll.InsertEnd(1)
				ll.InsertEnd(4)
				ll.InsertEnd(3)
				ll.InsertEnd(6)
				ll.InsertEnd(5)

				return ll
			},
		},
		{
			name: "2",
			args: args{head: func() *LinkedList {
				ll := new(LinkedList)
				ll.InsertEnd(1)
				ll.InsertEnd(2)
				ll.InsertEnd(3)
				ll.InsertEnd(4)
				ll.InsertEnd(5)

				return ll
			}},
			want: func() *LinkedList {
				ll := new(LinkedList)
				ll.InsertEnd(2)
				ll.InsertEnd(1)
				ll.InsertEnd(4)
				ll.InsertEnd(3)
				ll.InsertEnd(5)

				return ll
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want(), swapPairs(tt.args.head()), "swapPairs(%v)", tt.args.head)
		})
	}
}

func Test_swapNodes(t *testing.T) {
	type args struct {
		ll func() *LinkedList
		k  int
	}
	tests := []struct {
		name string
		args args
		want func() *LinkedList
	}{
		{
			name: "At start k = 1",
			args: args{
				ll: func() *LinkedList {
					ll := new(LinkedList)
					ll.InsertEnd(1)
					ll.InsertEnd(2)
					ll.InsertEnd(3)
					ll.InsertEnd(4)
					ll.InsertEnd(5)

					return ll
				},
				k: 1,
			},
			want: func() *LinkedList {
				ll := new(LinkedList)
				ll.InsertEnd(5)
				ll.InsertEnd(2)
				ll.InsertEnd(3)
				ll.InsertEnd(4)
				ll.InsertEnd(1)

				return ll
			},
		},
		{
			name: "At mid, k = 3: result LL is same",
			args: args{
				ll: func() *LinkedList {
					ll := new(LinkedList)
					ll.InsertEnd(1)
					ll.InsertEnd(2)
					ll.InsertEnd(3)
					ll.InsertEnd(4)
					ll.InsertEnd(5)

					return ll
				},
				k: 1,
			},
			want: func() *LinkedList {
				ll := new(LinkedList)
				ll.InsertEnd(1)
				ll.InsertEnd(2)
				ll.InsertEnd(3)
				ll.InsertEnd(4)
				ll.InsertEnd(5)

				return ll
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want(), swapNodes(tt.args.ll(), tt.args.k), "swapNodes(%v, %v)", tt.args.ll, tt.args.k)
		})
	}
}

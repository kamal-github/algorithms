package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Returns merged sorted list",
			args: args{
				lists: []*ListNode{
					{
						Val:  1,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
					{
						Val:  1,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "When all of second list contains bigger items",
			args: args{
				lists: []*ListNode{
					{
						Val:  1,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
					{
						Val:  4,
						Next: &ListNode{
							Val:  5,
							Next: &ListNode{
								Val:  6,
								Next: &ListNode{
									Val:  7,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{
								Val:  5,
								Next: &ListNode{
									Val:  6,
									Next: &ListNode{
										Val:  7,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Returns merged sorted list of different sizes",
			args: args{
				lists: []*ListNode{
					{
						Val:  1,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
					{
						Val:  1,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
					{
						Val:  2,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{
								Val:  4,
								Next: &ListNode{
									Val:  4,
									Next: &ListNode{
										Val:  5,
										Next: &ListNode{
											Val:  6,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

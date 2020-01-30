package leetcode

import "testing"

func TestKnapsack(t *testing.T) {
	type args struct {
		items []Item
		maxW  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				items: []Item{{2, 6}, {2, 10}, {3, 12}},
				maxW:  5,
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KnapsackDP(tt.args.items, tt.args.maxW); got != tt.want {
				t.Errorf("knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}

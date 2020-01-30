package leetcode

import (
	"math"
	"testing"
)

func TestCoinExchange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with positive integer, return valid count",
			args:args{
				coins:  []int{1,2,5},
				amount: 11,
			},
			want:3,
		},
		{
			name: "with positive integer and valid count, return -1",
			args:args{
				coins:  []int{2},
				amount: 1,
			},
			want: math.MaxInt64,
		},
		{
			name: "with other numbers",
			args:args{
				coins:  []int{10,6,1},
				amount: 12,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinExchange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("CoinExchange() = %v, want %v", got, tt.want)
			}
		})
	}
}

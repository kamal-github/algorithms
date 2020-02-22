package leetcode

import (
	"testing"
)

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

func TestTransformInput(t *testing.T) {
	type args struct {
		w      []int
		values []int
		maxW   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				w:      []int{2, 3, 1, 4},
				values: []int{4, 5, 3, 7},
				maxW:   5,
			},
			want: 10,
		},

		//{
		//	args: args{
		//		w: []int{7, 0, 30, 22, 80, 94, 11, 81, 70, 64, 59, 18, 0, 36, 3, 8, 15, 42, 9, 0,
		//			42, 47, 52, 32, 26, 48, 55, 6, 29, 84, 2, 4, 18, 56, 7, 29, 93, 44, 71,
		//			3, 86, 66, 31, 65, 0, 79, 20, 65, 52, 13},
		//		values: []int{360, 83, 59, 130, 431, 67, 230, 52, 93, 125, 670, 892, 600, 38, 48, 147,
		//			78, 256, 63, 17, 120, 164, 432, 35, 92, 110, 22, 42, 50, 323, 514, 28,
		//			87, 73, 78, 15, 26, 78, 210, 36, 85, 189, 274, 43, 33, 10, 19, 389, 276,
		//			312},
		//		maxW: 850,
		//	},
		//	want: 7534,
		//},
		//{
		//	args: args{
		//		w: []int{382745,
		//			799601,
		//			909247,
		//			729069,
		//			467902,
		//			44328,
		//			34610,
		//			698150,
		//			823460,
		//			903959,
		//			853665,
		//			551830,
		//			610856,
		//			670702,
		//			488960,
		//			951111,
		//			323046,
		//			446298,
		//			931161,
		//			31385,
		//			496951,
		//			264724,
		//			224916,
		//			169684},
		//		values: []int{
		//			825594,
		//			1677009,
		//			1676628,
		//			1523970,
		//			943972,
		//			97426,
		//			69666,
		//			1296457,
		//			1679693,
		//			1902996,
		//			1844992,
		//			1049289,
		//			1252836,
		//			1319836,
		//			953277,
		//			2067538,
		//			675367,
		//			853655,
		//			1826027,
		//			65731,
		//			901489,
		//			577243,
		//			466257,
		//			369261,
		//		},
		//		maxW: 6404180,
		//	},
		//	want: 2,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KnapsackWithArrayInput(tt.args.w, tt.args.values, tt.args.maxW); got != tt.want {
				t.Errorf("KnapsackWithArrayInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

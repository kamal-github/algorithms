package leetcode

import "testing"

func TestMaxRotateFunction(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"With positive numbers",
			args:args{A:[]int{4,3,2,6}},
			want:26,
		},
		{
			name:"With negative numbers",
			args:args{A:[]int{-2,-2}},
			want:-2,
		},
		{
			name:"With empty array",
			args:args{A:[]int{}},
			want:0,
		},
		{
			name:"With array of size 1",
			args:args{A:[]int{}},
			want:0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxRotateFunction(tt.args.A); got != tt.want {
				t.Errorf("MaxRotateFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

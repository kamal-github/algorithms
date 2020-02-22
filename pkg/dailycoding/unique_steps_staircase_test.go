package dailycoding

import "testing"

func TestUniqueStaircase(t *testing.T) {
	type args struct {
		n     int
		steps []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:     4,
				steps: []int{1, 2, 3},
			},
			want: 7,
		},
		{
			args: args{
				n:     40,
				steps: []int{1, 2, 3},
			},
			want: 23837527729,
		},
		{
			args: args{
				n:     100,
				steps: []int{1, 2, 3},
			},
			want: 7367864567128947527,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueStaircase(tt.args.n, tt.args.steps); got != tt.want {
				t.Errorf("UniqueStaircase() = %v, want %v", got, tt.want)
			}
		})
	}
}

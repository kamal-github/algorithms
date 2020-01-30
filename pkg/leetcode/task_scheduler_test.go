package leetcode

import "testing"

func TestLeastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"With some input",
			args    {
				tasks: []byte{'A','A','A','A','A','A', 'B', 'C', 'D','E','F','G'},
				n:     2,
			},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("LeastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

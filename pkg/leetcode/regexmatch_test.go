package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args:args{
				s: "abc",
				p: "abc",
			},
			want: true,
		},
		{
			name: "",
			args:args{
				s: "abc",
				p: "ab",
			},
		},
		{
			name: "",
			args:args{
				s: "abc",
				p: "abcd",
			},
		},
		{
			name: "",
			args:args{
				s: "aa",
				p: "a*",
			},
			want:true,
		},
		{
			name: "for all dots, returns true",
			args:args{
				s: "aabuu",
				p: "a.b..",
			},
			want:true,
		},
		{
			name: "for all start, returns true",
			args:args{
				s: "aab",
				p: "c*a*b",
			},
			want:true,
		},
		{
			name: "for all start, returns true",
			args:args{
				s: "aa",
				p: "a",
			},
			want:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("IsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

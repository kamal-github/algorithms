package ctci

import "testing"

func TestOneAway(t *testing.T) {
	type args struct {
		s1 []byte
		s2 []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s1: []byte("pale"),
				s2: []byte("ple"),
			},
			want: true,
		},
		{
			args: args{
				s1: []byte("ple"),
				s2: []byte("pale"),
			},
			want: true,
		},
		{
			args: args{
				s1: []byte("pale"),
				s2: []byte("pal"),
			},
			want: true,
		},
		{
			args: args{
				s1: []byte("ple"),
				s2: []byte("pales"),
			},
		},
		{
			args: args{
				s1: []byte("pale"),
				s2: []byte("bake"),
			},
		},
		{
			args: args{
				s1: []byte("pale"),
				s2: []byte("bale"),
			},
			want: true,
		},
		{
			args: args{
				s1: []byte("ppale"),
				s2: []byte("tppale"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OneEditAway(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("OneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}

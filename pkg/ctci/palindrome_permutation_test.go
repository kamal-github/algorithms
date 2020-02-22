package ctci

import "testing"

func TestPalindromePermutation(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{s: []byte("tact coa")},
			want: true,
		},
		{
			args: args{s: []byte("akhil")},
		},
		{
			args: args{s: []byte("sam masst")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromePermutation(tt.args.s); got != tt.want {
				t.Errorf("PalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

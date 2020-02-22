package dailycoding

import "testing"

func TestWaysOfEncoding(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{str: "111"},
			want: 3,
		},
		{
			args: args{str: "1234"},
			want: 3,
		},
		{
			args: args{str: "12345"},
			want: 5,
		},
		{
			args: args{str: "000"},
			want: 0,
		},
		{
			args: args{str: "001"},
			want: 0,
		},
		{
			args: args{str: "10"},
			want: 1,
		},
		{
			args: args{str: "301"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumDecodings(tt.args.str); got != tt.want {
				t.Errorf("NumDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}

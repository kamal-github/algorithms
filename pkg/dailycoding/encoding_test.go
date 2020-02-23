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
			name: "111",
			args: args{str: "111"},
			want: 3,
		},
		{
			name: "1234",
			args: args{str: "1234"},
			want: 3,
		},
		{
			name: "12345",
			args: args{str: "12345"},
			want: 3,
		},
		{
			args: args{str: "000"},
			want: 0,
		},
		{
			name: "001",
			args: args{str: "001"},
			want: 0,
		},
		{
			name: "10",
			args: args{str: "10"},
			want: 1,
		},
		{
			name: "301",
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

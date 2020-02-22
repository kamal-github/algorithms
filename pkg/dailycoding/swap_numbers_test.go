package dailycoding

import "testing"

func TestSwap(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			args: args{
				a: 3,
				b: 4,
			},
			want:  4,
			want1: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Swap(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("Swap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Swap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwap2(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			args: args{
				a: 0,
				b: 4,
			},
			want:  4,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Swap2(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("Swap2() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Swap2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

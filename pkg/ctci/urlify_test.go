package ctci

import (
	"reflect"
	"testing"
)

func TestURLify(t *testing.T) {
	type args struct {
		s []byte
		n int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				s: []byte("Mr John Smith    "),
				n: 13,
			},
			want: []byte("Mr%20John%20Smith"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLify(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLify() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

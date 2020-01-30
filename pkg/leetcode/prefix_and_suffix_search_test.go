package leetcode

import (
	"testing"
)

func TestTrieNode_Insert(t *testing.T) {
	type fields struct {
		m         map[int32]*TrieNode
		weights   []int
		endOfPath bool
	}
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "",
			args:   args{
				s: "apple",
				i: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tn := NewTrieNode()
			
			tn.Insert(tt.args.s, tt.args.i)
			tn.Insert("applica", 1)
			tn.Search("apple")
			t.Log(tn)
		})
	}
}

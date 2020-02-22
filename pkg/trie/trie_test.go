package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t1 *testing.T) {
	type fields struct {
		children    map[byte]*Trie
		endOfBranch bool
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				children: make(map[byte]*Trie),
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				children:    tt.fields.children,
				endOfBranch: tt.fields.endOfBranch,
			}

			t.Insert("duck")
			t.Insert("fuck")
			t.Insert("ducking")
			t.Insert("doom")

			s := t.GetStringsStartsWith("fu")
			assert.Equal(t1, []string{"fuck"}, s)

			s1 := t.GetStringsStartsWith("d")
			assert.ElementsMatch(t1, []string{"doom", "duck", "ducking"}, s1)

			s2 := t.GetStringsStartsWith("xxx")
			assert.Equal(t1, []string{}, s2)

			assert.True(t1, t.StartsWith("du"))
			assert.False(t1, t.StartsWith("dc"))
			assert.True(t1, t.StartsWith("duck"))
			assert.True(t1, t.StartsWith("do"))
			assert.True(t1, t.StartsWith("ducki"))

			assert.True(t1, t.Search("duck"))
			assert.True(t1, t.Search("fuck"))
			assert.False(t1, t.Search("fu"))

		})
	}
}

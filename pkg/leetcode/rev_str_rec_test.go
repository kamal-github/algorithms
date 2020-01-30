package leetcode_test

import (
	"testing"

	"github.com/kamal-github/algorithms/pkg/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestReverseInPlace(t *testing.T) {
	type testCase struct {
		name string
		in   []byte
		out  []byte
	}

	testCases := []testCase{
		{
			name: "reverse a odd length string",
			in:   []byte("abc"),
			out:  []byte("cba"),
		},
		{
			name: "reverse a even length string",
			in:   []byte("abcd"),
			out:  []byte("dcba"),
		},
		{
			name: "reverse a empty string",
			in:   []byte(""),
			out:  []byte(""),
		},
		{
			name: "reverse a one byte string",
			in:   []byte("K"),
			out:  []byte("K"),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.NotPanics(t, func() {
				leetcode.ReverseInPlace(tc.in)
				assert.Equal(t, tc.out, tc.in)
			})
		})
	}
}

package leetcode_test

import (
	"testing"

	"github.com/kamal-github/algorithms/pkg/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	out := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	arr := leetcode.Generate(5)
	assert.Equal(t, out, arr)
}

func TestGenerateEff(t *testing.T) {
	out := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1}, // k = 3
		{1, 4, 6, 4, 1},
		{1, 5, 10, 10, 5, 1},
		{1, 6, 15, 20, 15, 6, 1},
		{1, 7, 21, 35, 35, 21, 7, 1},
		{1, 8, 28, 56, 70, 56, 28, 8, 1},
		{1, 9, 36, 84, 126, 126, 84, 36, 9, 1}}
	arr := leetcode.GenerateEff(10)
	assert.Equal(t, out, arr)
}

func TestGetRow(t *testing.T) {
	out := []int{1,
		33,
		528,
		5456,
		40920,
		237336,
		1107568,
		4272048,
		13884156,
		38567100,
		92561040,
		193536720,
		354817320,
		573166440,
		818809200,
		1037158320,
		1166803110,
		1166803110,
		1037158320,
		818809200,
		573166440,
		354817320,
		193536720,
		92561040,
		38567100,
		13884156,
		4272048,
		1107568,
		237336,
		40920,
		5456,
		528,
		33,
		1,
	}
	assert.Equal(t, out, leetcode.GetRow(33))
}

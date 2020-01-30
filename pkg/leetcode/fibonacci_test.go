package leetcode_test

import (
	"testing"

	"github.com/kamal-github/algorithms/pkg/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	n := leetcode.Fibonacci(3)
	assert.Equal(t, 2, n)
}
func TestFibonacciEff(t *testing.T) {
	n := leetcode.FibonacciEff(30)
	assert.Equal(t, 832040, n)
}

func TestClimbStairs(t *testing.T) {
	c := leetcode.ClimbStairs(30)
	assert.Equal(t, 1346269, c)

	d := leetcode.ClimbStairsBF(5)
	assert.Equal(t, 8, d)

	e := leetcode.ClimbStairsDP(30)
	assert.Equal(t, 1346269, e)
	f := leetcode.ClimbStairsDP(1)
	assert.Equal(t, 1, f)
}

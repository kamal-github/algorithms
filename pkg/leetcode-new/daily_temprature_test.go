package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	act := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	exp := []int{1, 1, 4, 2, 1, 1, 0, 0}
	assert.Equal(t, exp, act)
}

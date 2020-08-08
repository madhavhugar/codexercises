package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthMissingPositiveNumber(t *testing.T) {
	nums := []int{2, 3, 4, 7, 11}
	k := 5
	got := findKthPositive(nums, k)
	wanted := 9
	assert.Equal(t, got, wanted)

	nums = []int{1, 2, 3, 4}
	k = 2
	got = findKthPositive(nums, k)
	wanted = 6
	assert.Equal(t, got, wanted)
}

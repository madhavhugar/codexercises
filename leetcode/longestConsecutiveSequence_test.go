package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutiveSequence_nlogn(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
		expected := 9

		got := longestconsecutivesequenceNlogn(nums)
		assert.Equal(t, expected, got)
	})

	t.Run("example 2", func(t *testing.T) {
		nums := []int{100, 4, 200, 1, 3, 2}
		expected := 4

		got := longestconsecutivesequenceNlogn(nums)
		assert.Equal(t, expected, got)
	})

	t.Run("example 3", func(t *testing.T) {
		nums := []int{}
		expected := 0

		got := longestconsecutivesequenceNlogn(nums)
		assert.Equal(t, expected, got)
	})

	t.Run("example 4", func(t *testing.T) {
		nums := []int{1, 2, 0, 1}
		expected := 3

		got := longestconsecutivesequenceNlogn(nums)
		assert.Equal(t, expected, got)
	})
}

func TestLongestConsecutiveSequence_n(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
		expected := 9

		got := longestConsecutiveSequenceN(nums)
		assert.Equal(t, expected, got)
	})

	t.Run("example 2", func(t *testing.T) {
		nums := []int{100, 4, 200, 1, 3, 2}
		expected := 4

		got := longestConsecutiveSequenceN(nums)
		assert.Equal(t, expected, got)
	})

	t.Run("example 3", func(t *testing.T) {
		nums := []int{}
		expected := 0

		got := longestConsecutiveSequenceN(nums)
		assert.Equal(t, expected, got)
	})
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		num1 := []int{1,3}
		num2 := []int{2}

		got := findMedianSortedArrays(num1, num2)
		expected := 2.0

		assert.Equal(t, expected, got)
	})

	t.Run("example 2", func(t *testing.T) {
		num1 := []int{1,2}
		num2 := []int{3,4}

		got := findMedianSortedArrays(num1, num2)
		expected := 2.5

		assert.Equal(t, expected, got)
	})

	t.Run("example 3", func(t *testing.T) {
		num1 := []int{}
		num2 := []int{1}

		got := findMedianSortedArrays(num1, num2)
		expected := 1.0

		assert.Equal(t, expected, got)
	})

	t.Run("example 4", func(t *testing.T) {
		num1 := []int{2}
		num2 := []int{}

		got := findMedianSortedArrays(num1, num2)
		expected := 2.0

		assert.Equal(t, expected, got)
	})
}
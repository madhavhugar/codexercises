package main

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func dotest(t *testing.T, items []interface{}, k int, wanted []interface{}) {
	fmt.Println("input:", items, k)
	var dead []interface{}
	got := josephusPermutation(items, k, dead, k-1)
	assert.Equal(t, wanted, got)
}

func TestJosephusPermutation(t *testing.T) {
	items := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 1
	dotest(t, items, k, result)

	items = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result = []interface{}{2, 4, 6, 8, 10, 3, 7, 1, 9, 5}
	k = 2
	dotest(t, items, k, result)

	items = []interface{}{1, 2, 3, 4, 5, 6, 7}
	result = []interface{}{3, 6, 2, 7, 5, 1, 4}
	k = 3
	dotest(t, items, k, result)

	// items = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// result = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// k = 40
	// dotest(t, items, k, result)
}

func TestRemoveByIndex(t *testing.T) {
	items := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wanted := []interface{}{1, 2, 3, 4, 6, 7, 8, 9, 10}
	index := 4
	got, removed := removeByIndex(items, index)
	assert.Equal(t, wanted, got)
	assert.Equal(t, 5, removed)
}

func TestRemoveIndices(t *testing.T) {
	items := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wanted := []interface{}{1, 2, 3, 4, 6, 7, 8, 10}
	indices := []int{4, 8}
	got, removed := removeIndices(items, indices)
	assert.Equal(t, wanted, got)
	assert.Equal(t, []interface{}{5, 9}, removed)

	items = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wanted = []interface{}{1, 2, 4, 6, 7, 8, 10}
	indices = []int{2, 4, 8}
	got, removed = removeIndices(items, indices)
	assert.Equal(t, wanted, got)
	assert.Equal(t, []interface{}{3, 5, 9}, removed)
}

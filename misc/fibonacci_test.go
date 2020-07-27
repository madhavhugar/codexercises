package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	got := fibonacci(5)
	assert.Equal(t, 5, got)
}

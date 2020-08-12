package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	inputs := []int{123321, 12321, 12332}
	expected := []bool{true, true, false}

	for index := range inputs {
		assert.Equal(t, expected[index], isPalindrome(inputs[index]))
	}
}

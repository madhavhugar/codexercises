package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCsvParser(t *testing.T) {
	filename := "regular.csv"
	expected := map[string]string{
		"key":   "value",
		"foo":   "bar",
		"hello": "world",
	}

	got, err := csvParser(filename)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleFileReader(t *testing.T) {
	expected := "Hello World"
	got, err := simpleFileReader("./dat")

	require.NoError(t, err)
	require.Equal(t, expected, got)
}

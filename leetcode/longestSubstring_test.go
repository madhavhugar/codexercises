package main

import "testing"

func TestLongestSubstring(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		input := "abcabcbb"
		got := lengthOfLongestSubstring(input)
		expected := 3
		if got != expected {
			t.Errorf("got != expected; %d != %d", got, expected)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		input := "bbbb"
		got := lengthOfLongestSubstring(input)
		expected := 1
		if got != expected {
			t.Errorf("got != expected; %d != %d", got, expected)
		}
	})

	t.Run("example 3", func(t *testing.T) {
		input := ""
		got := lengthOfLongestSubstring(input)
		expected := 0
		if got != expected {
			t.Errorf("got != expected; %d != %d", got, expected)
		}
	})

	t.Run("example 4", func(t *testing.T) {
		input := "pwwkew"
		got := lengthOfLongestSubstring(input)
		expected := 3
		if got != expected {
			t.Errorf("got != expected; %d != %d", got, expected)
		}
	})
}
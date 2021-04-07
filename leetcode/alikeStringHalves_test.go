package main

import "testing"

func TestAlikeStringHalves(t *testing.T) {
	input := "book"
	want := true

	got := alikeStringHalves(input)
	if got != want {
		t.Errorf("got != want")
	}
}

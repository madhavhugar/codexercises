package codility

import "testing"

func TestMaxProfit(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		prices := []int{7, 1, 5, 3, 6, 4}
		expected := 5

		got := maxProfit(prices)
		if got != expected {
			t.Errorf("got != expected; %d != %d", got, expected)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		prices := []int{7, 6, 4, 3, 1}
		expected := 0

		got := maxProfit(prices)
		if got != expected {
			t.Errorf("got != expected; %d != %d", got, expected)
		}
	})

	t.Run("example 3", func(t *testing.T) {
		prices := []int{1, 2}
		expected := 1

		got := maxProfit(prices)
		if got != expected {
			t.Errorf("got != expected; %d != %d", got, expected)
		}
	})
}

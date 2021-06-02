package codility

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombineOppositeSorted(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		asc := []int{4, 7, 9, 10}
		desc := []int{12, 9, 7, 4, 2, 1}
		sorted := combineOppositeSorted(asc, desc)

		expected := []int{1, 2, 4, 4, 7, 7, 9, 9, 10, 12}
		assert.Equal(t, expected, sorted)
	})
}

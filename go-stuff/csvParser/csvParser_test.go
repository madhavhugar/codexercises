package csvParser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCsvParser(t *testing.T) {
	csv := NewCsvReader("regular.csv")

	f, err := csv.FieldNames()
	assert.NoError(t, err)
	assert.Equal(t, f, []string{"key", "value"})
}

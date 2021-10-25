package csvParser

import (
	"bufio"
	"os"
	"strings"
)

type CsvReader interface {
	ReadAll() [][]string
	ReadColumn(columnName string) []string
	FieldNames() ([]string, error)
}

type csv struct {
	filename string
}

func NewCsvReader(filename string) CsvReader {
	return &csv{
		filename: filename,
	}
}

func (c *csv) ReadAll() [][]string {
	return nil
}

func (c *csv) ReadColumn(columnName string) []string {
	return nil
}

func (c *csv) FieldNames() ([]string, error) {
	f, err := os.Open(c.filename)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)
	s.Scan()
	line := s.Text()
	cols := strings.Split(line, ",")
	colNames := make([]string, 0, len(cols))

	for _, e := range cols {
		colNames = append(colNames, strings.TrimSpace(e))
	}

	return colNames, nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func csvParser(filename string) (map[string]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	parsedMap := make(map[string]string)
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
		v := strings.Split(line, ",")
		key := strings.TrimSpace(v[0])
		value := strings.TrimSpace(v[1])
		parsedMap[key] = value
	}

	return parsedMap, nil
}
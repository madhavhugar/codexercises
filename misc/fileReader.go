package main

import (
	"fmt"
	"io/ioutil"
)

func simpleFileReader(name string) (string, error) {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}
	fmt.Println(string(f))
	return string(f), nil
}

package main

import "fmt"

func isPalindrome(x int) bool {
	xStr := fmt.Sprintf("%d", x)
	isItThough := true
	length := len(xStr)

	for i := 0; i < length; i++ {
		j := length - i - 1
		if xStr[i] != xStr[j] {
			isItThough = false
		}
	}
	return isItThough
}

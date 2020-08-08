package main

// https://leetcode.com/contest/biweekly-contest-32/problems/kth-missing-positive-number/

func findKthPositive(arr []int, k int) int {
	originalK := k
	first := arr[0]
	if first != 1 {
		k = k - first + 1
	}
	if k-1 < 0 {
		return originalK
	}
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		currK := k
		if diff > 1 {
			k = k - diff + 1
		}
		if k <= 0 {
			return arr[i] + currK
		}
	}
	return arr[len(arr)-1] + k
}

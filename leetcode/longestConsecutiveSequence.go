package main

import "sort"

func longestconsecutivesequenceNlogn(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	sort.Ints(nums)

	maxCon, currCon := 0, 0
	for i := 0; i < n-1; i++ {
		if nums[i+1] == nums[i] {
			continue
		} else if (nums[i+1] - nums[i]) == 1 {
			currCon++
			if maxCon < currCon {
				maxCon = currCon
			}
		} else {
			currCon = 0
		}
	}

	return maxCon + 1
}

func longestConsecutiveSequenceN(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	seen := make(map[int]bool)
	for _, n := range nums {
		seen[n] = true
	}

	currCon, maxCon := 0, 0
	for _, n := range nums {
		currCon = 0
		if seen[n-1] {
			continue
		}

		l := n
		for seen[l+1] {
			currCon++
			l++
		}

		if maxCon < currCon {
			maxCon = currCon
		}
	}

	return maxCon + 1
}

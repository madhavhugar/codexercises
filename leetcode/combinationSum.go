package main

import "sort"

var count int

func combinationSum(nums []int, target int) int {
	count = 0
	sort.Ints(nums)
	helper(nums, target)
	return count
}

func helper(nums []int, target int) {
	if target <= 0 {
		count += 1
		return
	}

	for i := 0; i < len(nums); i += 1 {
		if nums[i] <= target {
			helper(nums, target - nums[i])
		} else {
			return
		}
	}
}
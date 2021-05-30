package main

func findMedianSortedArrays(num1 []int, num2 []int) float64 {
	m := len(num1)
	n := len(num2)

	nums := make([]int, m+n)
	for i, j, k := 0, 0, 0; k < m+n; k += 1 {
		if i >= m {
			nums[k] = num2[j]
			j += 1
		} else if j >= n {
			nums[k] = num1[i]
			i += 1
		} else if num1[i] <= num2[j] {
			nums[k] = num1[i]
			i += 1
		} else {
			nums[k] = num2[j]
			j += 1
		}
	}

	midIndex := (m+n-1)/2
	if r := (m+n) % 2; r != 0 {
		return float64(nums[midIndex])
	}
	return float64(nums[midIndex] + nums[int(midIndex)+1])/2
}
package main

func lengthOfLongestSubstring(s string) int {
	max := 0
	currMax := 0
	seen := make(map[uint8]int)

	for i := 0; i < len(s); i += 1 {
		if seenAtIndex, ok := seen[s[i]]; ok {
			if currMax > max {
				max = currMax
			}
			currMax = 0
			seen = make(map[uint8]int)
			i = seenAtIndex
		} else {
			seen[s[i]] = i
			currMax += 1
		}
	}

	if currMax > max {
		return currMax
	}

	return max
}
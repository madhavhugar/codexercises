package codility

func combineOppositeSorted(asc []int, desc []int) []int {
	m := len(asc)
	n := len(desc)
	totalLen := m + n
	sorted := make([]int, totalLen)

	for i, j, k := 0, n-1, 0; k < totalLen; k += 1 {
		if i >= m {
			sorted[k] = desc[j]
			j--
		} else if j < 0 {
			sorted[k] = asc[i]
			i++
		} else if asc[i] <= desc[j] {
			sorted[k] = asc[i]
			i++
		} else {
			sorted[k] = desc[j]
			j--
		}
	}

	return sorted
}

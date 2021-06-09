package codility

// DH slight variation of the task
// find max profit with two transactions

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	profits := make([]int, len(prices))
	for i := 0; i < len(prices)-1; i++ {
		buyPrice := prices[i]
		profit := 0
		for j := i; j < len(prices); j++ {
			possibleSellPrice := prices[j] - buyPrice
			if possibleSellPrice > profit {
				profit = possibleSellPrice
			}
		}
		profits[i] = profit
	}

	maximumProfit := 0
	for _, n := range profits {
		if n > maximumProfit {
			maximumProfit = n
		}
	}

	return maximumProfit
}

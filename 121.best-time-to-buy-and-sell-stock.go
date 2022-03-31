package leetcode

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min, maxProfit := prices[0], 0
	for _, v := range prices {
		if v < min {
			min = v
		} else if (v - min) > maxProfit {
			maxProfit = v - min
		}
	}
	return maxProfit
}

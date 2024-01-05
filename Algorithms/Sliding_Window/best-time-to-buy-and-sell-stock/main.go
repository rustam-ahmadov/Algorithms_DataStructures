package main

// Constraints:

// 	1 <= prices.length <= 105
// 	0 <= prices[i] <= 104

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	profit, max := 0, 0
	f := prices[0]
	for i := 1; i < len(prices); i++ {
		profit = prices[i] - f

		if f > prices[i] {
			f = prices[i]
		}
		if profit > max {
			max = profit
		}
	}
	return max
}

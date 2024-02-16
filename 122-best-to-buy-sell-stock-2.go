package main

func MaxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var maxProfit = 0

	for i := 0; i < len(prices)-1; i++ {
		if profit := prices[i+1] - prices[i]; profit > 0 {
			maxProfit += profit
		}
	}

	return maxProfit
}

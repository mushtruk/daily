package main

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var (
		minPrice  = prices[0]
		maxProfit = 0
	)

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

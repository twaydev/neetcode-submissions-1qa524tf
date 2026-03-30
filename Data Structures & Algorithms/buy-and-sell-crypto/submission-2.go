
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buyPrice := 0
	maxProfit := 0
	for i, price := range prices {
		// set min buyPrice
		if i == 0 || buyPrice > price {
			buyPrice = price
		}

		if price > buyPrice {
			profit := price - buyPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
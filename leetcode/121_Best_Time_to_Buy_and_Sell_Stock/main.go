package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	buy, sell := prices[0]+1, prices[0]+1
	amount := 0
	local_amount := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			sell = prices[i]
			for j := i + 1; j < len(prices); j++ {
				if sell < prices[j] {
					sell = prices[j]
				}
				local_amount = sell - buy
			}
		}
		if local_amount > amount {
			amount = local_amount
		}
	}

	return amount
}

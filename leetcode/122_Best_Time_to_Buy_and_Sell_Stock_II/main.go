package main

func maxProfit(prices []int) int {
	maxSum := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxSum += (prices[i] - prices[i-1])
		}
	}

	return maxSum
}

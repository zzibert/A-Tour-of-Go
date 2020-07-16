package main

func maxProfit(prices []int) int {
	valley := prices[0]
	peak := prices[0]
	maxSum := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < valley {
			maxSum += (peak - valley)
			valley = prices[i]
			peak = prices[i]
		} else if prices[i] > peak {
			peak = prices[i]
		} else if prices[i] < peak {
			maxSum += (peak - valley)
			peak = prices[i]
			valley = prices[i]
		}
	}

	if peak > valley {
		maxSum += (peak - valley)
	}

	return maxSum
}

package main

func coinChange(coins []int, amount int) int {
	coins = sort(coins)
	numberOfCoins := 0

	for i := 0; i < len(coins); i++ {
		quotient, remainder := amount/coins[i], amount%coins[i]
		numberOfCoins += quotient
		amount = remainder
	}

	if amount != 0 {
		return -1
	}

	return numberOfCoins
}

func sort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		min := array[i]
		index := i
		for j := i + 1; j < len(array); j++ {
			if array[j] > min {
				min = array[j]
				index = j
			}
		}
		temp := array[i]
		array[i] = min
		array[index] = temp
	}

	return array
}

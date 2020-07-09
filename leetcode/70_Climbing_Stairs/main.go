package main

func climbStairs(n int) int {
	distincWays := make(map[int]int)
	distincWays[0] = 0
	distincWays[1] = 1
	distincWays[2] = 2

	for i := 3; i < n+1; i++ {
		distincWays[i] = distincWays[i-1] + distincWays[i-2]
	}

	return distincWays[n]
}

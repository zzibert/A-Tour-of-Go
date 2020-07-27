package main

func numJewelsInStones(J string, S string) int {
	jewels := make(map[rune]bool)

	var counter int

	for _, jewel := range J {
		jewels[jewel] = true
	}

	for _, stone := range S {
		if jewels[stone] {
			counter++
		}
	}

	return counter
}

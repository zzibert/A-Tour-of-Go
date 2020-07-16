package main

func firstUniqChar(s string) int {

	counts := make(map[rune]int)
	for _, letter := range s {
		counts[letter]++
	}

	for i, letter := range s {
		if counts[letter] == 1 {
			return i
		}
	}

	return -1
}

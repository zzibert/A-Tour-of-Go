package main

func isAnagram(s string, t string) bool {
	letters := make(map[rune]int)

	for _, letter := range s {
		letters[letter]++
	}

	for _, letter := range t {
		letters[letter]--
	}

	for _, count := range letters {
		if count != 0 {
			return false
		}
	}

	return true
}

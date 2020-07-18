package main

func shortestDistance(words []string, word1 string, word2 string) int {
	var index1, index2 int
	distance := len(words)

	for i, word := range words {
		if word == word1 {
			index1 = i
			index2 = findWord(words[i+1:], word2)
			if index2 == 0 {
				break
			}
			if index2 < distance {
				distance = index2
			}
		}
		if word == word2 {
			index2 = i
			index1 = findWord(words[i+1:], word1)
			if index1 == 0 {
				break
			}
			if index1 < distance {
				distance = index1
			}
		}
	}
	return distance
}

func findWord(words []string, word string) int {
	for i, val := range words {
		if val == word {
			return i + 1
		}
	}
	return 0
}

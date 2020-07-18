package main

func shortestDistance(words []string, word1 string, word2 string) int {
	distance := len(words)

	index1, index2 := -1, -1

	for i, word := range words {
		if word == word1 {
			index1 = i
			if index2 != -1 && index1-index2 < distance {
				distance = index1 - index2
			}
		} else if word == word2 {
			index2 = i
			if index1 != -1 && index2-index1 < distance {
				distance = index2 - index1
			}
		}
	}

	return distance
}

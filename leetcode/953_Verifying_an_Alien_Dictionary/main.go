package main

func isAlienSorted(words []string, order string) bool {
	alphabet := make(map[byte]int)

	for i := 0; i < len(order); i++ {
		alphabet[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		if !compareWords(words[i], words[i+1], alphabet) {
			return false
		}
	}
	return true

}

func compareWords(word1 string, word2 string, alphabet map[byte]int) bool {
	for i := 0; i < len(word1); i++ {
		if i == len(word2) {
			return false
		}

		if alphabet[word1[i]] < alphabet[word2[i]] {
			return true
		}
		if alphabet[word1[i]] > alphabet[word2[i]] {
			return false
		}
	}
	return true
}

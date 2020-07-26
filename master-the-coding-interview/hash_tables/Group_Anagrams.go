package main

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, word := range strs {
		anagram := sortString(word)
		anagrams[anagram] = append(anagrams[anagram], word)
	}

	array := make([][]string, 0)

	for _, words := range anagrams {
		array = append(array, words)
	}

	return array
}

func sortString(str string) string {
	bytes := []byte(str)

	for i := 0; i < len(bytes); i++ {
		letter := bytes[i]
		index := i
		for j := i + 1; j < len(bytes); j++ {
			if bytes[j] < letter {
				letter = bytes[j]
				index = j
			}
		}
		temp := bytes[i]
		bytes[i] = letter
		bytes[index] = temp
	}

	return string(bytes)
}

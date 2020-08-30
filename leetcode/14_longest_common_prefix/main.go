package main

func longestCommonPrefix(strs []string) string {
	longestCommonPrefix := ""

	shortestString := findShortestString(strs)

	for i := 0; i < shortestString; i++ {
		if compareChar(i, strs) {
			longestCommonPrefix += string(strs[0][i])
		} else {
			break
		}
	}
	return longestCommonPrefix
}

func compareChar(index int, strs []string) bool {
	for i := 0; i < len(strs)-1; i++ {
		if strs[i][index] != strs[i+1][index] {
			return false
		}
	}
	return true
}

func findShortestString(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	length := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < length {
			length = len(strs[i])
		}
	}

	return length
}

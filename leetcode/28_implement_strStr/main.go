package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	length := len(haystack)

	for i := 0; i < length-(len(needle)-1); i++ {
		if compareString(haystack[i:(i+len(needle))], needle) {
			return i
		}
	}

	return -1
}

func compareString(str1, str2 string) bool {

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}

	return true
}

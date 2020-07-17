package main

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	size := len(str)

	for i := 0; i < (size / 2); i++ {
		if str[i] != str[(size-1)-i] {
			return false
		}
	}

	return true
}

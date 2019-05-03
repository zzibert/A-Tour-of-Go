package camelCase

import "unicode"

func camelcase(s string) int32 {
	var counter int32 = 1
	for _, char := range s {
		if char == unicode.ToUpper(char) {
			counter++
		}
	}
	return counter
}

package main

import (
	"strconv"
	"strings"
)

func groupStrings(strings []string) [][]string {
	sequences := make(map[string][]string)

	for _, word := range strings {
		sequence := toSequence(word)
		sequences[sequence] = append(sequences[sequence], word)
	}

	array := make([][]string, 0)

	for _, words := range sequences {
		array = append(array, words)
	}

	return array

}

func toSequence(str string) string {
	bytes := []byte(str)
	array := make([]string, 0)

	for i := 0; i < len(bytes)-1; i++ {
		var diff int
		if bytes[i] < bytes[i+1] {
			diff = int(bytes[i] - (bytes[i+1] - 26))
		} else {
			diff = int(bytes[i] - bytes[i+1])
		}

		array = append(array, strconv.Itoa(diff))
	}

	integers := strings.Join(array, "")

	return integers
}

package main

import (
	"fmt"
	"strings"
)

func caesarCipher(s string, k int32) {
	var alphabet = "abcdefghijklmnopqrstuvwxyz"
	newAlphabet := make(map[rune]rune)
	for _, ch := range alphabet {
		newCh := (ch + k) % 122
		newAlphabet[ch] = newCh
	}
	newString := make([]string, 0)
	for _, ch := range s {
		if 96 <= ch && ch <= 122 {
			ch = newAlphabet[ch]
		}
		runeStr := string(ch)
		newString = append(newString, runeStr)
	}

	result := strings.Join(newString, "")
	fmt.Println(result)

}

func main() {
	caesarCipher("middle-Outz", 2)
}

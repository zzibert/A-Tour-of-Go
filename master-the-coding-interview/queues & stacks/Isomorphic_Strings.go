package main

func isIsomorphic(s string, t string) bool {
	characterOrderS := make(map[rune][]int)
	characterOrderT := make(map[rune][]int)

	for i, letter := range s {
		characterOrderS[letter] = append(characterOrderS[letter], i)
	}

	for i, letter := range t {
		characterOrderT[letter] = append(characterOrderT[letter], i)
	}

	for i, letter := range s {
		otherLetter := rune(t[i])
		if !sameSlice(characterOrderS[letter], characterOrderT[otherLetter]) {
			return false
		}
	}

	return true
}

func sameSlice(array1 []int, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}

	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			return false
		}
	}

	return true
}

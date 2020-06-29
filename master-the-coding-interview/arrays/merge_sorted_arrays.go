package main

import "fmt"

func main() {
	fmt.Println(mergeSortedArrays([]int{0, 3, 4, 31}, []int{4, 6, 30}))
}

// naive version
func mergeSortedArrays(array1, array2 []int) []int {
	var mergedArray []int
	var i, j int

	for {
		if array1[i] < array2[j] {
			mergedArray = append(mergedArray, array1[i])
			i++
		} else {
			mergedArray = append(mergedArray, array2[j])
			j++
		}
		if i == len(array1) {
			return append(mergedArray, array2[j:]...)
		}
		if j == len(array2) {
			return append(mergedArray, array1[i:]...)
		}
	}
}

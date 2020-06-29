package main

import "fmt"

func main() {
	fmt.Println(mergeSortedArrays([]int{0, 3, 4, 31}, []int{4, 6, 30}))
	fmt.Println(mergeSortedArrays([]int{0, 3, 4, 31}, []int{}))
	fmt.Println(mergeSortedArrays([]int{}, []int{4, 6, 30}))
}

// naive version, I would say time complexity is O(n)
func mergeSortedArrays(array1, array2 []int) []int {
	var mergedArray []int
	var i, j int
	length1, length2 := len(array1), len(array2)

	// Checking for empty arrays
	if length1 == 0 {
		return array2
	} else if length2 == 0 {
		return array1
	}

	// Comparing elements from each array, appending the smaller one, incrementing index.
	for {
		if array1[i] < array2[j] {
			mergedArray = append(mergedArray, array1[i])
			i++
		} else {
			mergedArray = append(mergedArray, array2[j])
			j++
		}

		// Checking if we got to the end of array.
		if i == length1 {
			return append(mergedArray, array2[j:]...)
		}
		if j == length2 {
			return append(mergedArray, array1[i:]...)
		}
	}
}

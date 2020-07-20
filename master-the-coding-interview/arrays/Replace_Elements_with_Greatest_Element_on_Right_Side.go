package main

func replaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = greatestElement(arr, i)
	}

	return arr
}

func greatestElement(arr []int, index int) int {
	length := len(arr)

	if index == length-1 {
		return -1
	}

	max := arr[index+1]
	for index = index + 1; index < length; index++ {
		if arr[index] > max {
			max = arr[index]
		}
	}

	return max
}

package main

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			moveElements(arr, i+1)
			i++
			break
		}
	}
}

func moveElements(arr []int, index int) {
	temp := 0
	temp2 := arr[index]

	for index < len(arr) {
		arr[index] = temp
		temp = temp2
		if index < len(arr)-1 {
			index++
			temp2 = arr[index]
		} else {
			break
		}
	}
}

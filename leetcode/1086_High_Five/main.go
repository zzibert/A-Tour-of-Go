package main

func highFive(items [][]int) [][]int {
	students := make(map[int][]int, 0)

	for _, item := range items {
		id, score := item[0], item[1]
		students[id] = append(students[id], score)
	}
	newArray := make([][]int, 0)

	for id, scores := range students {
		newArray = append(newArray, []int{id, Average(scores)})
	}

	return SortStudents(newArray)

}

func SortStudents(array [][]int) [][]int {
	for i := 0; i < len(array); i++ {
		index := i
		localMin := array[i][0]
		for j := i + 1; j < len(array); j++ {
			if array[j][0] < localMin {
				localMin = array[j][0]
				index = j
			}
		}
		temp := array[i]
		array[i] = array[index]
		array[index] = temp
	}
	return array
}

func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		localMax := array[i]
		index := i
		for j := i + 1; j < len(array); j++ {
			if array[j] > localMax {
				localMax = array[j]
				index = j
			}
		}
		temp := array[i]
		array[i] = array[index]
		array[index] = temp
	}
	return array
}

func Average(array []int) int {
	array = SelectionSort(array)
	array = array[:5]

	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	return sum / len(array)
}

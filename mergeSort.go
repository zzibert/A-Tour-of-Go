package main

import "fmt"

// Insert the ints one by one in stdio

func sort(array []int, c chan []int) {
	finished := 0
	var counter int32
	for finished != 1 {
		counter = 0
		for i := 0; i < (len(array) - 1); i++ {
			if array[i] > array[i+1] {
				counter = 1
				temp := array[i]
				array[i] = array[i+1]
				array[i+1] = temp
			}
		}
		if counter == 0 {
			finished = 1
		}
	}
	fmt.Println(array)
	c <- array
}

func mergeSort(c chan []int) {
	channel := make(chan []int, 1)
	sliceArray := make([][]int, 0)
	for i := 0; i < 4; i++ {
		sliceArray = append(sliceArray, <-c)
	}
	array := make([]int, 0)
	for _, slice := range sliceArray {
		for _, number := range slice {
			array = append(array, number)
		}
	}
	fmt.Println()
	sort(array, channel)
}

func main() {
	c := make(chan []int, 4)
	var number int
	array := make([]int, 0)
	for {
		_, err := fmt.Scanf("%d\n", &number)
		if err != nil {
			break
		}
		array = append(array, number)
	}
	size := len(array) / 4
	number = 0
	for i := 0; i < len(array); i = i + size {
		if number == 3 {
			go sort(array[i:len(array)], c)
			break
		} else {
			go sort(array[i:(i+size)], c)
			number++
		}
	}
	mergeSort(c)
}

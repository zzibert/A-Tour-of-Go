package main

import "fmt"

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
	array := []int{82, 33, 34, 10, 83, 9, 42, 17, 65, 16, 3, 1, 32, 12, 49, 64}
	for i := 0; i < len(array); i = i + (len(array) / 4) {
		go sort(array[i:(i+(len(array)/4))], c)
	}
	mergeSort(c)
}

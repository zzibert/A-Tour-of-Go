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
	c <- array
}

func main() {
	c := make(chan []int, 4)
	array := []int{20, 5, 76, 15, 95, 34, 11, 48, 80, 3, 30, 12, 77, 24, 72, 66}
	for i := 0; i < len(array); i = i + 4 {
		sort(array[i:(i+4)], c)
	}
	a := <-c
	b := <-c
	d := <-c
	e := <-c
	fmt.Println(a, b, d, e)
}

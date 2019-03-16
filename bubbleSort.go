package main

import "fmt"

func BubbleSort(slice []int) {

	done := 0

	for done != 1 {
		counter := 0
		for index := 0; index < (len(slice) - 1); index++ {
			if slice[index] > slice[index+1] {
				Swap(slice, index)
				counter++
			}
		}
		if counter == 0 {
			done = 1
		}
	}
}

func Swap(slice []int, index int) {
	temp := slice[index+1]
	slice[index+1] = slice[index]
	slice[index] = temp
}

func main() {
	slice := make([]int, 0)
	number := 0

	for i := 0; i < 10; i++ {
		_, err := fmt.Scanf("%d\n", &number)
		if err != nil {
			fmt.Println(err)
		}
		slice = append(slice, number)
	}

	fmt.Println("original slice")

	fmt.Println(slice)

	fmt.Println("Slice after bubble sort")

	BubbleSort(slice)

	fmt.Println(slice)
}

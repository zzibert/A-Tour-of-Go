package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := make([]int, 3)
	var number int
	for true {
		fmt.Println("Enter an integer!")
		_, ok := fmt.Scanf("%d\n", &number)
		slice = append(slice, number)
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
		fmt.Println(slice)
		if ok != nil {
			fmt.Println(ok)
			break
		}
	}
}

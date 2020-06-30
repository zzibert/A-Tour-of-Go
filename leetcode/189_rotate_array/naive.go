package main

import (
	"fmt"
)

func main() {
	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	fmt.Println(rotate([]int{-1, -100, 3, 99}, 2))
}

// naive solution, trying O(n)
func rotate(nums []int, k int) []int {
  
  newArray := make([]int, len(nums))
  for 
}

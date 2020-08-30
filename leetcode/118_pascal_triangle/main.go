package main

import "fmt"

func main() {

	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	array := make([][]int, 0)

	if numRows == 0 {
		return array
	}

	triangle := []int{1}

	array = append(array, triangle)

	for i := 1; i < numRows; i++ {
		triangle = generateTriangle(triangle)
		array = append(array, triangle)
	}

	return array

}

func generateTriangle(triangle []int) []int {

	newTriangle := make([]int, 0)
	for i := 0; i < len(triangle)+1; i++ {
		if i == 0 || i == len(triangle) {
			newTriangle = append(newTriangle, 1)
		} else {
			newTriangle = append(newTriangle, triangle[i-1]+triangle[i])
		}
	}
	return newTriangle
}

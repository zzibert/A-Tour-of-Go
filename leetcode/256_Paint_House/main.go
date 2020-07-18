package main

import "fmt"

func main() {
	costs := [][]int{[]int{6, 4, 13}, []int{10, 9, 15}, []int{14, 15, 11}, []int{17, 15, 9}, []int{7, 10, 13}, []int{18, 9, 4}, []int{5, 20, 12}}

	fmt.Println(minCost(costs))
}

type Level struct {
	n     int
	index int
}

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	values := make(map[Level]int)
	house := costs[0]
	first := createTree(house[0], 0, costs[1:], &values)
	second := createTree(house[1], 1, costs[1:], &values)
	third := createTree(house[2], 2, costs[1:], &values)

	if first <= second && first <= third {
		return first
	} else if second <= first && second <= third {
		return second
	}
	return third

}

func createTree(root int, index int, costs [][]int, values *map[Level]int) int {
	length := len(costs)

	if length == 0 {
		return root
	}

	if value, ok := (*values)[Level{n: length, index: index}]; ok {
		return value
	}

	house := costs[0]
	leftIndex, rightIndex := 0, 0

	if index == 0 {
		house = house[1:]
		leftIndex = 1
		rightIndex = 2
	} else if index == 1 {
		house = []int{house[0], house[2]}
		leftIndex = 0
		rightIndex = 2
	} else {
		house = house[:2]
		leftIndex = 0
		rightIndex = 1
	}
	left := createTree(house[0], leftIndex, costs[1:], values)
	right := createTree(house[1], rightIndex, costs[1:], values)

	if left < right {
		(*values)[Level{n: length, index: index}] = left + root
		return left + root
	}
	(*values)[Level{n: length, index: index}] = root + right
	return right + root
}

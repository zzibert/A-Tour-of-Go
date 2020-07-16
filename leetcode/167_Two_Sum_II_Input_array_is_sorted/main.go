package main

func twoSum(numbers []int, target int) []int {
	indexes := make(map[int]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		complement := target - numbers[i]

		if indexes[complement] != 0 {
			return []int{indexes[complement], i + 1}
		}
		indexes[numbers[i]] = i + 1
	}

	return []int{}
}

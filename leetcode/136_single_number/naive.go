package main

func main() {

}

// naive approach
func singleNumber(nums []int) int {
	numbers := make(map[int]int)

	for _, number := range nums {
		if _, ok := numbers[number]; !ok {
			numbers[number] = 1
		} else {
			numbers[number] += 1
		}
	}

	for number, count := range numbers {
		if count == 1 {
			return number
		}
	}
	return 0
}

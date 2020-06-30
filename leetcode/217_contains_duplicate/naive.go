package main

func main() {

}

// naive solution, O(n)
func containsDuplicate(nums []int) bool {
	seenNumbers := make(map[int]bool, 0)
	for _, num := range nums {
		if seenNumbers[num] {
			return true
		} else {
			seenNumbers[num] = true
		}
	}
	return false
}

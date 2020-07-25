package main

func containsDuplicate(nums []int) bool {
	alreadyAppeared := make(map[int]bool, 0)

	for _, number := range nums {
		if alreadyAppeared[number] {
			return true
		}
		alreadyAppeared[number] = true
	}

	return false
}

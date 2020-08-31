package main

func removeElement(nums []int, val int) int {
	length := 0
	first := 0
	last := len(nums) - 1

	if len(nums) == 0 {
		return 0
	}

	for first != last {
		if nums[first] != val {
			first++
			length++
		} else {
			for last > first && nums[last] == val {
				last--
			}
			if last == first {
				return length
			}
			nums[first] = nums[last]
			last--
		}
	}

	if nums[first] != val {
		length++
	}

	return length
}

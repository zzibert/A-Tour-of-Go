package main

func rob(nums []int) (amount int) {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	scoring := make(map[int]int, 0)

	scoring[0] = nums[0]

	if nums[0] > nums[1] {
		amount = nums[0]
	} else {
		amount = nums[1]
	}

	scoring[1] = amount

	for i := 2; i < len(nums); i++ {
		value := nums[i] + scoring[i-2]
		if value > amount {
			scoring[i] = value
			amount = value
		} else {
			scoring[i] = amount
		}
	}

	return

}

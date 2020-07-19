package main

func maxSubArray(nums []int) int {
	length := len(nums)

	globalSum := nums[0]

	// the size of the array
	for size := 1; size <= length; size++ {
		// first element in the array
		for i := 0; i < length; i++ {
			// last element in the array
			localSum := 0
			last := i + size
			if last > length {
				break
			}
			for j := i; j < last; j++ {
				localSum += nums[j]
			}
			if localSum > globalSum {
				globalSum = localSum
			}
		}
	}

	return globalSum
}

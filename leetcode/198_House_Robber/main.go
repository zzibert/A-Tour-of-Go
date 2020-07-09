package main

func rob(nums []int) int {
	houses := make(map[int]bool)
	var money int

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return maxNum(nums[0], nums[1])
	} else if len(nums) == 3 {
		return maxNum(nums[1], nums[0]+nums[2])
  }
  
  for i:=0; i < len(nums); i++ {
    if !houses[i] {
      if nums 
    }
  }
}

func maxNum(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

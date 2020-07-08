package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	length := len(nums)

	if length == 0 {
		return nil
	}

	if length == 1 {
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}

	index, max := findMax(nums)

	node := TreeNode{Val: max, Left: constructMaximumBinaryTree(nums[:index]), Right: constructMaximumBinaryTree(nums[(index + 1):])}

	return &node
}

func findMax(nums []int) (index int, max int) {
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
			index = i
		}
	}
	return
}

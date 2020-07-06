package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}

	middle := len(nums) / 2.0
	node := &TreeNode{Val: nums[middle]}

	node.Left = sortedArrayToBST(nums[:middle])
	node.Right = sortedArrayToBST(nums[(middle + 1):])

	return node
}

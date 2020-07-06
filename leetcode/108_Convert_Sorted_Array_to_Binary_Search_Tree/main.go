package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	if n == 1 {
		return &TreeNode{Val: nums[0]}
	}

	middle := n / 2.0
	node := &TreeNode{Val: nums[middle]}

	node.Left = sortedArrayToBST(nums[:middle])
	node.Right = sortedArrayToBST(nums[(middle + 1):])

	return node
}

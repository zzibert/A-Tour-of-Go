package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := make([]int, 0)

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	return append(append([]int{root.Val}, left...), right...)
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	array := append(inorderTraversal(root.Left), root.Val)

	return append(array, inorderTraversal(root.Right)...)
}

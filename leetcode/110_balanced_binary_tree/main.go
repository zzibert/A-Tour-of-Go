package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

}

func findMax(node *TreeNode) int {

	if node == nil {
		return 0
	}

	left, right := findMax(node.Left)+1, findMax(node.Right)+1

	if left > right {
		return left
	}

	return right
}

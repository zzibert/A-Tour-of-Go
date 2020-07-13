package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	center := root.Val

	left := maxPathSum(root.Left)
	right := maxPathSum(root.Right)

	if (left + center) < 0 {
		return right
	}

	if (right + center) < 0 {
		return left
	}

	if left < 0 {
		left = 0
	}

	if right < 0 {
		right = 0
	}

	if (center + left + right) <= 0 {
		return 0
	}

	return (center + left + right)
}

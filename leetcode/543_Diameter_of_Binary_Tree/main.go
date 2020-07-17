package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	diameterHelper(root, &maxDiameter)

	return maxDiameter
}

func diameterHelper(node *TreeNode, maxDiameter *int) int {
	if node == nil {
		return 0
	}

	left := diameterHelper(node.Left, maxDiameter)
	right := diameterHelper(node.Right, maxDiameter)

	if left+right+1 > *maxDiameter {
		*maxDiameter = left + right
	}

	if left > right {
		return left + 1
	}

	return right + 1
}

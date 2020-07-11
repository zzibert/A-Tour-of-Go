package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	maxDepth := maxDepth(root)

	return compareDepth(root, 0, maxDepth)

}

func compareDepth(node *TreeNode, depth, maxDepth int) bool {
	if node == nil {
		if maxDepth-depth <= 2 {
			return true
		}

		return false
	}

	return compareDepth(node.Left, depth+1, maxDepth) && compareDepth(node.Right, depth+1, maxDepth)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 1
	}

	left := 1 + maxDepth(root.Left)
	right := 1 + maxDepth(root.Right)

	if left > right {
		return left
	}

	return right
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {

	depth := maxDepth(root, 0)

	return findMaxLeaves(root, 0, depth)

}

func findMaxLeaves(node *TreeNode, depth, maxDepth int) int {
	if node == nil {
		return 0
	}

	if depth == maxDepth {
		return node.Val
	}

	return findMaxLeaves(node.Left, depth+1, maxDepth) + findMaxLeaves(node.Right, depth+1, maxDepth)
}

func maxDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	left := maxDepth(node.Left, depth+1)
	right := maxDepth(node.Right, depth+1)

	if left > right {
		return left
	}

	return right
}

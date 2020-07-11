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

	if !hasSon(root) {
		return true
	}

	leftGrandSons, rightGrandSons := hasGrandSon(root.Left), hasGrandSon(root.Right)

	if !leftGrandSons && !rightGrandSons {
		return true
	}

	if !leftGrandSons || !rightGrandSons {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)

}

func hasGrandSon(node *TreeNode) bool {
	if node == nil {
		return false
	}

	if !hasSon(node.Left) && !hasSon(node.Right) {
		return false
	}

	return true
}

func hasSon(node *TreeNode) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return false
	}

	return true
}

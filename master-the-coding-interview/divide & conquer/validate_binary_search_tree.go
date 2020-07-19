package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && Max(root.Left) >= root.Val {
		return false
	}

	if root.Right != nil && Min(root.Right) <= root.Val {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)

}

func Min(root *TreeNode) int {
	if root.Left == nil {
		return root.Val
	}
	return Min(root.Left)
}

func Max(root *TreeNode) int {
	if root.Right == nil {
		return root.Val
	}
	return Max(root.Right)
}

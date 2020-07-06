package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	if L > root.Val {
		root = trimBST(root.Right, L, R)
	} else if R < root.Val {
		root = trimBST(root.Left, L, R)
	} else {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
	}

	return root
}

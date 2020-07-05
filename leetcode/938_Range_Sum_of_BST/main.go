package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	if root.Val >= L && root.Val <= R {
		return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
	} else if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	} else {
		return rangeSumBST(root.Left, L, R)
	}
}

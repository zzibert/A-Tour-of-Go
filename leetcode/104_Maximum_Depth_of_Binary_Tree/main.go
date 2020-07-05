package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxRight := maxDepth(root.Right)
	maxLeft := maxDepth(root.Left)
	if maxRight < maxLeft {
		return 1 + maxLeft
	} else {
		return 1 + maxRight
	}

}

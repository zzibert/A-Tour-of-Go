package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		temp := root.Right
		root.Right = invertTree(root.Left)
		root.Left = invertTree(temp)
	}
	return root
}

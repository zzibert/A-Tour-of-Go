package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	sum := 0

	if root == nil {
		return sum
	}

	if root.Val%2 == 0 {
		grandSons := getGrandSons(root)
		for _, grandson := range grandSons {
			if grandson != nil {
				sum += grandson.Val
			}
		}
	}

	return sumEvenGrandparent(root.Left) + sumEvenGrandparent(root.Right) + sum
}

func getGrandSons(node *TreeNode) []*TreeNode {
	if node == nil {
		return []*TreeNode{}
	}
	return append(getSons(node.Left), getSons(node.Right)...)
}

func getSons(node *TreeNode) []*TreeNode {
	if node == nil {
		return []*TreeNode{}
	}
	return []*TreeNode{node.Left, node.Right}
}

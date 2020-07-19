package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	inOrder := inOrderTraversal(root)

	var node *TreeNode

	for _, val := range inOrder {
		node = insert(node, val)
	}

	return node

}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	root.Right = insert(root.Right, val)

	return root

}

func inOrderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	array := append([]int{node.Val}, inOrderTraversal(node.Right)...)
	return append(inOrderTraversal(node.Left), array...)
}

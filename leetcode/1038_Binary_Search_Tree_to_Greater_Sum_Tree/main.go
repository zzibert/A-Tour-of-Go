package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {

	preorderArray := PreorderTraversal(root)
	TraverseAndReplaceVal(root, preorderArray)

	return root
}

func TraverseAndReplaceVal(node *TreeNode, preorder []int) {
	if node == nil {
		return
	}

	sum := 0
	i := 0

	for preorder[i] < node.Val && i < len(preorder) {
		i++
	}
	for ; i < len(preorder); i++ {
		sum += preorder[i]
	}

	node.Val = sum
	TraverseAndReplaceVal(node.Left, preorder)
	TraverseAndReplaceVal(node.Right, preorder)
}

func PreorderTraversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	array := append(PreorderTraversal(node.Left), node.Val)

	return append(array, PreorderTraversal(node.Right)...)
}

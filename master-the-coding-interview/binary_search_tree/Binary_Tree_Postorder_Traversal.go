package main

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	array := append(postorderTraversal(root.Left), postorderTraversal(root.Right)...)

	return append(array, root.Val)
}

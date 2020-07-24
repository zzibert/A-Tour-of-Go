package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	rootElement := preorder[0]

	root := &TreeNode{Val: rootElement}

	rootInorder := findRootInOrder(inorder, rootElement)

	if rootInorder == -1 {
		return nil
	}

	if includes(inorder[:rootInorder], preorder[1]) {
		root.Left = buildTree(preorder[1:], inorder[:rootInorder])
		var index int
		for index = 2; index < len(preorder) && includes(inorder[:rootInorder], preorder[index]); index++ {
		}
		root.Right = buildTree(preorder[index:], inorder[(rootInorder+1):])
	} else {
		root.Left = nil
		root.Right = buildTree(preorder[1:], inorder[(rootInorder+1):])
	}

	return root
}

func findRootInOrder(inorder []int, root int) int {
	for index, val := range inorder {
		if val == root {
			return index
		}
	}

	return -1
}

func includes(array []int, val int) bool {
	for _, number := range array {
		if number == val {
			return true
		}
	}

	return false
}

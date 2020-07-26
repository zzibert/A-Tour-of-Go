package main

import "strconv"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	traversals := make(map[string]int)

	rootNodes := make(map[string]*TreeNode)

	array := make([]*TreeNode, 0)

	findDuplicateSubtreesHelper(root, &traversals, &rootNodes)

	for traverseString, count := range traversals {
		if traverseString != "." && count > 1 {
			array = append(array, rootNodes[traverseString])
		}
	}

	return array

}

func findDuplicateSubtreesHelper(root *TreeNode, traversals *map[string]int, rootNodes *map[string]*TreeNode) {

	if root == nil {
		return
	}

	traverseString := inorderTraversal(root)

	(*traversals)[traverseString] = (*traversals)[traverseString] + 1
	(*rootNodes)[traverseString] = root

	findDuplicateSubtreesHelper(root.Left, traversals, rootNodes)
	findDuplicateSubtreesHelper(root.Right, traversals, rootNodes)

}

func inorderTraversal(root *TreeNode) string {
	if root == nil {
		return ""
	}

	if root.Left == nil && root.Right == nil {
		return "L" + strconv.Itoa(root.Val) + "R"
	}
	if root.Left == nil {
		return "L" + strconv.Itoa(root.Val) + inorderTraversal(root.Right)
	}
	if root.Right == nil {
		return inorderTraversal(root.Left) + strconv.Itoa(root.Val) + "R"
	}

	return inorderTraversal(root.Left) + strconv.Itoa(root.Val) + inorderTraversal(root.Right)
}

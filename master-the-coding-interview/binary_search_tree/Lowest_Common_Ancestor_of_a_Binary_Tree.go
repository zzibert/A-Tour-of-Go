package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p == q {
		return p
	}

	pPath := BstPath(root, p, []*TreeNode{})

	qPath := BstPath(root, q, []*TreeNode{})

	if len(pPath) > len(qPath) {
		diff := len(pPath) - len(qPath)

		pPath = pPath[:(len(pPath) - diff)]
	}

	if len(pPath) < len(qPath) {
		diff := len(qPath) - len(pPath)

		qPath = qPath[:(len(qPath) - diff)]
	}

	for length := len(qPath); length > 0; length-- {
		if pPath[length-1] == qPath[length-1] {
			return qPath[length-1]
		}
		pPath = pPath[:length-1]
		qPath = qPath[:length-1]
	}

	return nil
}

func BstPath(root, node *TreeNode, array []*TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	array = append(array, root)

	if root == node {
		return array
	}

	if BstPath(root.Left, node, array) == nil {
		return BstPath(root.Right, node, array)
	} else if BstPath(root.Left, node, array) == nil {
		return BstPath(root.Left, node, array)
	}

	return nil
}

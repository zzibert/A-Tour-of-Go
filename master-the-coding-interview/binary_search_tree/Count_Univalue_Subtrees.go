package main

func countUnivalSubtrees(root *TreeNode) int {
	count, _ := countUnivalSubtreesHelper(root)

	return count
}

func countUnivalSubtreesHelper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	if root.Left == nil && root.Right == nil {
		return 1, true
	}

	if root.Left == nil {
		count, unival := countUnivalSubtreesHelper(root.Right)
		if unival && root.Val == root.Right.Val {
			return count + 1, true
		} else {
			return count, false
		}
	}

	if root.Right == nil {
		count, unival := countUnivalSubtreesHelper(root.Left)
		if unival && root.Val == root.Left.Val {
			return count + 1, true
		} else {
			return count, false
		}
	}

	leftCount, leftUnival := countUnivalSubtreesHelper(root.Left)
	rightCount, rightUnival := countUnivalSubtreesHelper(root.Right)

	if leftUnival && rightUnival && root.Val == root.Right.Val && root.Val == root.Left.Val {
		return 1 + rightCount + leftCount, true
	}

	return rightCount + leftCount, false
}

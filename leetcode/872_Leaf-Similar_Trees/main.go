package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := leafs(root1)
	leafs2 := leafs(root2)

	return compare(leafs1, leafs2)
}

func leafs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	return append(leafs(root.Left), leafs(root.Right)...)
}

func compare(array1, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}

	for i := 0; i < len(array2); i++ {
		if array2[i] != array1[i] {
			return false
		}
	}

	return true
}

package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	array := make([]int, 0)

	if root == nil {
		return array
	}

	for _, child := range root.Children {
		array = append(array, preorder(child)...)
	}

	return append([]int{root.Val}, array...)
}

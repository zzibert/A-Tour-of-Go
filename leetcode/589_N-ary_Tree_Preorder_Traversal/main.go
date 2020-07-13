package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {

	if root == nil {
		return []int{}
	}

	array := []int{root.Val}

	for i := 0; i < len(root.Children); i++ {
		array = append(array, preorder(root.Children[i])...)
	}

	return array
}

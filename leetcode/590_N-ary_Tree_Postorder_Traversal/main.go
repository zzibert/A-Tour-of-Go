package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	array := make([]int, 0)

	if root == nil {
		return array
	}

	for i := 0; i < len(root.Children); i++ {
		array = append(array, postorder(root.Children[i])...)
	}

	return append(array, root.Val)

}

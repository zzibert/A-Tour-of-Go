package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depths := make([]int, 0)

	for _, node := range root.Children {
		depth := maxDepth(node)
		depths = append(depths, depth)
	}

	return max(depths) + 1
}

func max(array []int) int {
	max := 0
	for _, val := range array {
		if val > max {
			max = val
		}
	}
	return max
}

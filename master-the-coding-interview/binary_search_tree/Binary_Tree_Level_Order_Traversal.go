package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	masterArray := make([][]int, 0)

	queue := make([]*TreeNode, 0)

	queue = append(queue, root)

	counter := 0

	for {
		localQueue := queue[counter:]
		level := make([]int, 0)
		for len(localQueue) != 0 {
			node := localQueue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			level = append(level, node.Val)
			localQueue = localQueue[1:]
			counter++
		}
		if isEmpty(level) {
			return masterArray
		}
		masterArray = append(masterArray, level)
	}
}

func isEmpty(array []int) bool {
	return len(array) == 0
}

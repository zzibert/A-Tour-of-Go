package main

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftQueue := make([]*Node, 0)
	rightQueue := make([]*Node, 0)

	if root.Left != nil {
		leftQueue = append(leftQueue, root.Left)
	}
	if root.Right != nil {
		rightQueue = append(rightQueue, root.Right)
	}

	for len(leftQueue) != 0 || len(rightQueue) != 0 {
		leftCounter := len(leftQueue)
		rightCounter := len(rightQueue)
		// left nodes
		for i := 0; i < leftCounter; i++ {
			left := leftQueue[i]

			if left.Left != nil {
				leftQueue = append(leftQueue, left.Left)

			}
			if left.Right != nil {
				leftQueue = append(leftQueue, left.Right)

			}
			if i < leftCounter-1 {
				left.Next = leftQueue[i+1]
			}

		}
		// right nodes
		for i := 0; i < rightCounter; i++ {
			right := rightQueue[i]

			if right.Left != nil {
				rightQueue = append(rightQueue, right.Left)

			}
			if right.Right != nil {
				rightQueue = append(rightQueue, right.Right)
			}

			if i < rightCounter-1 {
				right.Next = rightQueue[i+1]
			}

		}
		if len(rightQueue) != 0 && len(leftQueue) != 0 {
			leftQueue[leftCounter-1].Next = rightQueue[0]
		}
		leftQueue = leftQueue[leftCounter:]
		rightQueue = rightQueue[rightCounter:]
	}

	return root
}

package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftQueue := make([]*Node, 0)
	rightQueue := make([]*Node, 0)

	leftQueue = append(leftQueue, root.Left)
	rightQueue = append(rightQueue, root.Right)

	for leftQueue[0] != nil {
		counter := len(leftQueue)
		for i := 0; i < counter; i++ {
			left := leftQueue[i]
			right := rightQueue[i]

			leftQueue = append(leftQueue, left.Left)
			leftQueue = append(leftQueue, left.Right)
			rightQueue = append(rightQueue, right.Left)
			rightQueue = append(rightQueue, right.Right)

			if i < counter-1 {
				left.Next = leftQueue[i+1]
				right.Next = rightQueue[i+1]
			}

		}
		leftQueue[counter-1].Next = rightQueue[0]
		leftQueue = leftQueue[counter:]
		rightQueue = rightQueue[counter:]
	}

	return root
}

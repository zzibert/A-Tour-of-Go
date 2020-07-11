package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {}

func maxLevelSum(root *TreeNode) int {
	maxAmount := root.Val
	queue := make([]*TreeNode, 0)
	counter := 1
	maxCounter := 1

	if root == nil {
		return 0
	}

	queue = append(queue, root)

	for {
		level := make([]int, 0)
		localAmount := 0
		for i := 0; i < counter; i++ {
			node := queue[0]

			if node == nil {
				level = append(level, 0)
				queue = append(queue, nil)
				queue = append(queue, nil)
			} else {
				//Left Node
				queue = append(queue, node.Left)
				//Right Node
				queue = append(queue, node.Right)

				level = append(level, node.Val)
			}

			queue = queue[1:]
		}

		if allNil(level) {
			return getPower(maxCounter)
		}

		for i := 0; i < len(level); i++ {
			localAmount += level[i]
		}

		if localAmount > maxAmount {
			maxAmount = localAmount
			maxCounter = counter
		}

		counter = counter * 2
	}
}

func allNil(array []int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] != 0 {
			return false
		}
	}
	return true
}

func getPower(counter int) int {
	power := 0
	for counter != 1 {
		counter = counter / 2
		power++
	}
	return power + 1
}

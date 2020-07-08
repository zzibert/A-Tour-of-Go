package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	counter := 1
	level := make([]interface{}, 0)

	if root == nil {
		return true
	}

	queue = append(queue, root)

	for {
		level = make([]interface{}, 0)

		for i := 0; i < counter; i++ {
			current := queue[0]

			// If current is nil
			if current == nil {
				queue = append(queue, nil)
				queue = append(queue, nil)

				level = append(level, nil)
				level = append(level, nil)
			} else {
				// left node
				queue = append(queue, current.Left)
				if current.Left == nil {
					level = append(level, nil)
				} else {
					level = append(level, current.Left.Val)
				}

				// right node

				queue = append(queue, current.Right)
				if current.Right == nil {
					level = append(level, nil)
				} else {
					level = append(level, current.Right.Val)
				}
			}
			queue = queue[1:]
		}

		if !isPalindrome(level) {
			return false
		}

		if isNilArray(level) {
			return true
		}

		counter = counter * 2
	}
}

func isNilArray(array []interface{}) bool {
	length := len(array)

	for i := 0; i < length; i++ {
		if array[i] != nil {
			return false
		}
	}
	return true
}

func isPalindrome(array []interface{}) bool {
	length := len(array)

	for i := 0; i < length/2; i++ {
		if array[i] != array[(length-1)-i] {
			return false
		}
	}

	return true
}

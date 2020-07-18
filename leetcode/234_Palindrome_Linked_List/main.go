package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	values := make([]int, 0)

	if head == nil {
		return true
	}

	// Traversing linked list, adding values to an array.
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	return palindrome(values)

}

// Checking if the array itself is a palindrome.
func palindrome(array []int) bool {
	for i := 0; i < (len(array) / 2); i++ {
		if array[i] != array[len(array)-1-i] {
			return false
		}
	}
	return true
}

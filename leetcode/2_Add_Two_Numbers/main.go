package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	counter := 0.0
	number := 0.0

	for l1 != nil {
		number += float64(l1.Val) * math.Pow(10.0, counter)
		l1 = l1.Next
		counter++
	}

	for l2 != nil {
		number += float64(l2.Val) * math.Pow(10.0, counter)
		l2 = l2.Next
		counter++
	}

	return number
}

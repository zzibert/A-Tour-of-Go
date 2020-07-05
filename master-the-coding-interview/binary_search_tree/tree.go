package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func main() {
	bst := &BST{}
	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(20)
	bst.Insert(1)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(170)

	Traverse(bst.root)
}

func (bst *BST) Insert(value int) {
	node := &Node{val: value, left: nil, right: nil}
	if bst.root == nil {
		bst.root = node
	} else {
		parentNode := bst.root
		current := parentNode
		for current != nil {
			parentNode = current
			if node.val < current.val {
				current = current.left
			} else {
				current = current.right
			}
		}
		if parentNode.val < node.val {
			parentNode.right = node
		} else {
			parentNode.left = node
		}
	}
}

func Traverse(node *Node) {
	if node != nil {
		fmt.Println(node.val)
		Traverse(node.left)
		Traverse(node.right)
	}
}

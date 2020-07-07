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

	fmt.Println(bst.Lookup(9))
	fmt.Println(bst.Lookup(8))
	fmt.Println(bst.Lookup(170))
	fmt.Println(bst.Lookup(1923))
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

func (bst *BST) Lookup(value int) *Node {

	current := bst.root
	for current != nil {
		if value == current.val {
			return current
		} else if value < current.val {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil
}

func (bst *BST) Remove(value int) *Node {
	if bst.root == nil {
		return nil
	}

	var parentNode *Node
	current := bst.root

	for current != nil && current.val != value {
		parentNode = current
		if current.val < value {
			current = current.left
		} else {
			current = current.left
		}
	}

	if current == nil {
		return bst.root
  }
  
  if parentNode.left

}

func Traverse(node *Node) {
	if node != nil {
		fmt.Println(node.val)
		Traverse(node.left)
		Traverse(node.right)
	}
}

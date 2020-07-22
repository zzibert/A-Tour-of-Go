package main

type Node struct {
	Prev  *Node
	Next  *Node
	Child *Node
}

//DEPTH FIRST SEARCH
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	current := root
	for current != nil {
		if current.Child != nil {
			temp := current.Next
			last := flattenHelper(current, current.Child)
			current.Child = nil
			last.Next = temp
			if temp != nil {
				temp.Prev = last
			}
		}
		current = current.Next
	}

	return root
}

func flattenHelper(root *Node, child *Node) *Node {
	if root == nil || child == nil {
		return root
	}
	root.Next = child
	child.Prev = root
	if child.Child == nil {
		flatten(child)
	}

	return flattenHelper(child, child.Next)
}

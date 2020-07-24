package main

// func buildTree(inorder []int, postorder []int) *TreeNode {

// 	current := &TreeNode{}
// 	root := current

// 	for len(inorder) > 0 && len(postorder) > 0 {
// 		in := inorder[0]
// 		post := postorder[0]

// 		if in == out {
// 			if current.Left == nil {
// 				current.Left = &TreeNode{Val: in}
// 				inorder = inorder[1:]
// 				postorder = postorder[1:]
// 				current.Val = inorder[0]
// 				inorder = inorder[1:]
// 			} else {
// 				current.Right = &TreeNode{Val: in}
// 				inorder = inorder[1:]
// 				postorder = postorder[1:]
// 			}
// 		}
// 		if in >= post {
// 			current = current.Right
// 			current.Left = &TreeNode{Val: post}
// 			current.Right = &TreeNode{Val: in}
// 			inorder = inorder[2:]
// 			postorder = postorder[2:]
// 		}
// 	}

// 	return root
// }

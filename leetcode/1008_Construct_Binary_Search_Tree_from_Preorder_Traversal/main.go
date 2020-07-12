package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]

	var index int

	for ; index < len(preorder); index++ {
		if preorder[index] > val {
			break
		}
	}

	return &TreeNode{Val: val, Left: bstFromPreorder(preorder[1:index]), Right: bstFromPreorder(preorder[index:])}

}

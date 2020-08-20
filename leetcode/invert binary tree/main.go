package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	mirror(root)
	return root
}

func mirror(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	mirror(root.Left)
	mirror(root.Right)
}

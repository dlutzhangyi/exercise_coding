package main

// https://leetcode.com/problems/symmetric-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isTwoTreeSame(root.Left, root.Right)
}

func isTwoTreeSame(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isTwoTreeSame(left.Left, right.Right) && isTwoTreeSame(left.Right, right.Left)
}

func main() {

}

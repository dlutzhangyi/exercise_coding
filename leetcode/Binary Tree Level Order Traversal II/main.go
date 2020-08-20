package main

import "fmt"

//https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	result := [][]int{}
	level := []int{}
	if root == nil {
		return result
	}
	queue = append(queue, root)
	queue = append(queue, nil)

	for len(queue) != 0 {
		front := queue[0]
		if len(queue) == 1 {
			queue = queue[0:0]
		} else {
			queue = queue[1:]
		}

		if front != nil {
			level = append(level, front.Val)
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		} else {
			temp := make([]int, len(level))
			copy(temp, level)
			result = append(result, temp)
			if len(queue) != 0 {
				queue = append(queue, nil)
				level = level[0:0]
			}
		}
	}
	return result
}

func createNode() *TreeNode {
	root := &TreeNode{
		Val: 3,
	}
	left := &TreeNode{
		Val: 4,
	}
	right := &TreeNode{
		Val: 5,
	}
	root.Right = right
	root.Left = left
	return root
}
func main() {
	root := createNode()
	fmt.Println(levelOrderBottom(root))
}

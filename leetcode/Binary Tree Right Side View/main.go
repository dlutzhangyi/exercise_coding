package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	queue = append(queue, nil)

	rightValue := root.Val
	for len(queue) != 0 {
		front := queue[0]
		if len(queue) == 1 {
			queue = queue[0:0]
		} else {
			queue = queue[1:]
		}

		if front != nil {
			rightValue = front.Val
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		} else {
			result = append(result, rightValue)
			if len(queue) != 0 {
				queue = append(queue, nil)
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
	fmt.Println(rightSideView(root))
}

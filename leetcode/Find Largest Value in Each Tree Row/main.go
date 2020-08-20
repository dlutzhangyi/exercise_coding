package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	INT_MAX = int(^uint32((0)) >> 1)
	INT_MIN = ^INT_MAX
)

func largestValues(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	queue = append(queue, nil)

	max := INT_MIN
	for len(queue) != 0 {
		front := queue[0]
		if len(queue) == 1 {
			queue = queue[0:0]
		} else {
			queue = queue[1:]
		}
		if front != nil {
			if front.Val > max {
				max = front.Val
			}
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		} else {
			result = append(result, max)
			max = INT_MIN
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

// https://leetcode.com/problems/find-largest-value-in-each-tree-row/
func main() {
	root := createNode()
	fmt.Println(largestValues(root))
}

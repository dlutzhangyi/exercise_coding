package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func construct(preorder []int, preorderStart int, preorderEnd int, inorder []int, inorderStart int, inorderEnd int) *TreeNode {
	if preorderStart > preorderEnd {
		return nil
	}
	rootValue := preorder[preorderStart]
	rootNode := &TreeNode{
		Val: rootValue,
	}
	inorderRootIndex := 0
	for i := inorderStart; i <= inorderEnd; i++ {
		if inorder[i] == rootValue {
			inorderRootIndex = i
			break
		}
	}

	leftLength := inorderRootIndex - inorderStart
	rightLength := inorderEnd - inorderRootIndex
	if leftLength > 0 {
		rootNode.Left = construct(preorder, preorderStart+1, preorderStart+leftLength, inorder, inorderStart, inorderRootIndex-1)
	}
	if rightLength > 0 {
		rootNode.Right = construct(preorder, preorderStart+leftLength+1, preorderEnd, inorder, inorderRootIndex+1, inorderEnd)
	}
	return rootNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(preorder) != len(inorder) {
		return nil
	}
	return construct(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func preorderTraverse(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d \n", node.Val)
	preorderTraverse(node.Left)
	preorderTraverse(node.Right)
}

func inorderTraverse(node *TreeNode) {
	if node != nil {
		return
	}
	inorderTraverse(node.Left)
	fmt.Printf("%d \n", node.Val)
	inorderTraverse(node.Right)
}
func main() {
	//preorder := []int{3, 9, 20, 15, 7}
	//inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inorder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	treeNode := buildTree(preorder, inorder)
	preorderTraverse(treeNode)
	fmt.Println()
	inorderTraverse(treeNode)
}
